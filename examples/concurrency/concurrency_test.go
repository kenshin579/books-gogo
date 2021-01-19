package concurrency

import "fmt"

func Example_goroutine() {
	go func() {
		fmt.Println("In goroutine")
	}()
	fmt.Println("In main routine")
	// Non-deterministic!
}

func Example_create_channels() {
	c1 := make(chan int)
	var c2 chan int = c1   //동일한 채널을 가리킨다
	var c3 <-chan int = c1 //c3 : 받기 전용
	var c4 chan<- int = c1 //c4 : 보내기 전용

	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	//Output:

}

//일대일 단방향 채널 소통
//이 코드는 서로 몇개의 데이터를 보내는지 알아야 함
func Example_simpleChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 1
	// 2
	// 3
}

//서로 데이터의 개수를 모르는 경우
func Example_simpleChannelForLoop() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c) //다 보내고 나서 채널을 닫음
	}()
	for num := range c {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func Example_simpleChannelReturnChannel() {
	ch := func() <-chan int { //채널을 반환하는 함수 - 데이터를 받아가기만하는 채널임
		c := make(chan int)
		go func() {
			defer close(c) //이 함수 끝나면 채널을 닫음
			c <- 1
			c <- 2
			c <- 3
		}()
		return c
	}()
	for num := range ch {
		fmt.Println(num)
	}
	// Output:
	// 1
	// 2
	// 3
}

func BabyNames(first, second string) <-chan string { //받기만하는 채널 반환
	c := make(chan string)
	go func() {
		defer close(c)
		for _, f := range first {
			for _, s := range second {
				c <- string(f) + string(s)
			}
		}
	}()
	return c
}

func Example_babyNames() {
	for n := range BabyNames("성정명재경", "준호우훈진") {
		fmt.Print(n, ", ")
	}
	// Output:
	// 성준, 성호, 성우, 성훈, 성진, 정준, 정호, 정우, 정훈, 정진, 명준, 명호, 명우, 명훈, 명진, 재준, 재호, 재우, 재훈, 재진, 경준, 경호, 경우, 경훈, 경진,
}

func Example_closedChannel() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// Output:
	// 0
	// 0
	// 0
}
