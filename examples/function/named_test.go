package function

import "fmt"

func ExampleBinOpToBinSub() {
	sub := BinOpToBinSub(func(a, b int) int {
		return a + b
	})
	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}

func Example_BinOp_임명_함수라_형변환이_자동으로_일어난다() {
	OpThreeAndFour(func(a, b int) int { //임명함수 (함수 리터럴)
		return 1 + 3
	})
	// Output:
	//4
}

func Example_BinOp_자료형_검사를_한다() {
	binOp := BinOp(func(a, b int) int {
		return 1 + 3
	})

	OpThreeAndFour(binOp)
	// Output:
	//4
}

func Example_BinSub_자료형_검사를_한다2() {
	binSub := BinSub(func(a, b int) int {
		return a + b
	})
	//OpThreeAndFour(binSub) //Cannot use 'binSub' (type BinSub) as type BinOp
	fmt.Println(binSub)

	// Output:
}
