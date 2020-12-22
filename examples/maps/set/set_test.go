package set

import "fmt"

func ExampleHasDupeRune1() {
	fmt.Println(hasDupeRune1("숨바꼭질"))
	fmt.Println(hasDupeRune1("다시합시다"))
	// Output:
	// false
	// true
}

func ExampleHasDupeRune2() {
	fmt.Println(hasDupeRune2("숨바꼭질"))
	fmt.Println(hasDupeRune2("다시합시다"))
	// Output:
	// false
	// true
}
