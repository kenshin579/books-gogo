package concurrency

import "fmt"

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
