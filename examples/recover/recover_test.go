package recover

import "fmt"

func f() (i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
		i = -1
	}()
	g()
	return 100
}

func g() {
	panic("I panic!")
}

func Example() {
	fmt.Println("f() = ", f())
	//Output:
	//Recovered in f I panic!
	//f() =  -1
}
