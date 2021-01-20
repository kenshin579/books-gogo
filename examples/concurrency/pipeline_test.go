package concurrency

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"context"
)

func Example_SimplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range SimplePlusOne(SimplePlusOne(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

func Example_SimplePlusTwo() {
	SimplePlusTwo := SimpleChain(SimplePlusOne, SimplePlusOne)
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range SimplePlusTwo(c) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

func Example_FanOut() {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	FanOut(c)
	//Output:
}

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	ctx := context.Background()
	for num := range PlusOne(ctx, PlusOne(ctx, c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

func Example_contextSwitching_fanOut() {
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	// Non-deterministic!
}

func ExampleFanIn3() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)
	c := FanIn3(c1, c2, c3) //여러 채널에서 나온 자료는 모두 c로 나오게 된다
	for n := range c {
		fmt.Print(n, ",")
	}

	//Output:
}

//todo : fanOut -> fanIn 작성해보기
func Example_Distribute() {
	//Output:
}

func ExamplePlusOne_consumeAll() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx := context.Background()
	nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	// fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	log.Println("NumGoroutine: ", runtime.NumGoroutine())
	for _ = range nums {
		// Consume all nums
	}
	time.Sleep(100 * time.Millisecond)
	// fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	log.Println("NumGoroutine: ", runtime.NumGoroutine())
	// Output:
	// 8
	// 18
}

func ExamplePlusOne_withCancel() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	nums := PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, PlusOne(ctx, c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			cancel()
			break
		}
	}
	// Output:
	// 8
	// 18
}
