package pattern

import "fmt"

func Example_Sqrt() {
	fmt.Printf("%.5f\n", Sqrt(2))
	// Output:
	// 1.41421
}

func Example_SlowerSqrt() {
	fmt.Printf("%.5f\n", SlowerSqrt(2))
	//Output:
	//1.41421
}
