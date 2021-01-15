package must

import (
	"fmt"
	"strconv"
)

func Example_error1() {
	id, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("id", id)

	//Output:id 123
}

func Must(i int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return i
}

func Example_error_must() {
	id := Must(strconv.ParseInt("123", 10, 64))
	fmt.Println("id", id)

	//Output:id 123
}
