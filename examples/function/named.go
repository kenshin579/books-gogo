package function

import "fmt"

type BinOp func(int, int) int
type BinSub func(int, int) int

//pure function이 아님
func BinOpToBinSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}

//Pure function
func OpThreeAndFour(f BinOp) {
	fmt.Println(f(3, 4))
}
