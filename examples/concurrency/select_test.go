package concurrency

import (
	"fmt"
	"time"
)

func Example_select() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)

	select {
	case n := <-c1:
		fmt.Println(n, "is from c1")
	case n := <-c2:
		fmt.Println(n, "is from c2")
	case c3 <- f():
		fmt.Println("1 is from c2")
	default:
		fmt.Println("No channel is ready")
	}

	//Output:
	//No channel is ready
}

func f() int {
	return 1
}

//채널값이 있으면 받고, 없으면 그냥 스킵하는 흐름
func Example_select_채널을_기다리지_않고_받기() {
	c := make(chan int)

	select {
	case n := <-c:
		fmt.Println(n)
	default:
		fmt.Println("Data is not ready. Skipping...")
	}

	//Output:
	//Data is not ready. Skipping...
}

func Example_select_시간_제한() {
	recv, send := make(chan int), make(chan int)

	select {
	case n := <-recv:
		fmt.Println(n)
	case send <- 1:
		fmt.Println("sent 1")
	case <-time.After(5 * time.Second):
		fmt.Println("No send and receive communications for 5 seconds")
		return
	}

	//Output:
}

//todo : for안에 타이머이 있으면 반복될 떄마다 새로 생성되므로, 한 번의 채널 통신마다 5초마다 제한 시간이 걸림 (1/2)
func Example_select_시간_제한2() {
	recv, send := make(chan int), make(chan int)
	timeout := time.After(5 * time.Second)

	for {
		select {
		case n := <-recv:
			fmt.Println(n)
		case send <- 1:
			fmt.Println("sent 1")
		case <-timeout:
			fmt.Println("No send and receive communications for 5 seconds")
			return
		}
	}

	//Output:
}
