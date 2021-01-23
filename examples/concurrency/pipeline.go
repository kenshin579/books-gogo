package concurrency

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

//in으로 받은 값 nums에 대해 num + 1 시켜 채널로 반환한다
//즉, 받기 전용 채널을 받아서 다른 받기 전용 채널을 돌려주는 함수임
func SimplePlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

//done 채널이 닫히면 output 채널도 담힘
func SimplePlusOneDone(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
				log.Printf("out <- sending %d - %d\n", num+1, num)
			case <-done: //close(done)을 하면 이곳을 타게 됨
				log.Println("done channel is closing", num)
				return
			}

		}
	}()
	return out
}

type SimpleIntPipe func(<-chan int) <-chan int

func SimpleChain(intPipes ...SimpleIntPipe) SimpleIntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, pipeFunc := range intPipes {
			c = pipeFunc(c) //첫번째 pipeFunc은 채널을 반환한고 다시 그 채널로 두번째 pipeFunc을 실행하는 구조이다
		}
		return c
	}
}

func FanOut(in <-chan int) {
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range in {
				time.Sleep(1)
				fmt.Printf("worker-%d %d\n", i, n)
			}
		}(i)
	}
}

// PlusOne returns a channel of num + 1 for nums received from in.
func PlusOne(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num + 1:
			case <-ctx.Done():
				log.Println("done", num)
				return
			}
		}
	}()
	return out
}

type IntPipe func(context.Context, <-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(ctx, c)
		}
		return c
	}
}

var PlusTwo = Chain(PlusOne, PlusOne)

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() { //todo : 왜 고루틴으로 실행하나? - 특별한 이유가 있나?
		wg.Wait()
		close(out)
	}()
	return out
}

func SimpleDistribute(p SimpleIntPipe, n int) SimpleIntPipe {
	return func(in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(in)
		}
		return FanIn(cs...)
	}
}

func Distribute(p IntPipe, n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		cs := make([]<-chan int, n)
		for i := 0; i < n; i++ {
			cs[i] = p(ctx, in)
		}
		return FanIn(cs...)
	}
}

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	openCnt := 3
	closeChan := func(c *<-chan int) bool {
		*c = nil //nil로 하면 채널이 block이 된다
		openCnt--
		return openCnt == 0
	}
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in1:
				if ok {
					out <- n
				} else if closeChan(&in1) {
					return
				}
			case n, ok := <-in2:
				if ok {
					out <- n
				} else if closeChan(&in2) {
					return
				}
			case n, ok := <-in3:
				if ok {
					out <- n
				} else if closeChan(&in3) {
					return
				}
			}
		}
	}()
	return out

}
