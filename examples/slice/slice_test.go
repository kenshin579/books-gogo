package slice

import "fmt"

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	fmt.Println(nums[:len(nums)-1])
	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
	// [1 2 3 4]
}

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)     // 이어붙이기 (가변 인자)
	f4 := append(f1[:2], f2...) // 토마토를 제외하고 이어붙이기

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
	// [사과 바나나 토마토]
	// [포도 딸기]
	// [사과 바나나 토마토 포도 딸기]
	// [사과 바나나 포도 딸기]
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
	// Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// [1 2 3]
	// len: 3
	// cap: 5
	//
	// [3 4 5]
	// len: 3
	// cap: 3
	//
	// [1 2 3 4]
	// len: 4
	// cap: 5
	//
	// [1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}

func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	//1. loop
	//for i:= range src {
	//	dest[i] = src[i]
	//}

	//2.copy()
	//if n:= copy(dest,src); n!= len(src){
	//	fmt.Println("복사가 덜 됐습니다.")
	//}

	//3.append
	dest = append([]int(nil), src...)

	//dest := src (이건 그냥 배열 포인터, 길이, 용량이 복사되는 거임)

	fmt.Println(dest)
	// Output:
	// [30 20 50 10 40]
}

//todo : 이 부분은 다시 보도록 함
func Example_slice_insert() {
	x := []int{7, 8, 9}
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a)

	a = append(a, x...)
	fmt.Println(a)

	i := 3
	copy(a[i+len(x):], a[i:])
	copy(a[i:], x)
	fmt.Print(a)

	// Output:
	//[1 2 3 4 5]
	//[1 2 3 4 5 7 8 9]
	//[1 2 3 7 8 9 4 5]
}

//todo : 이 부분은 다시 보도록 함
func Example_slice_delete() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a)
	i, k := 3, 1
	copy(a[i:], a[i+k:])
	for i := 0; i < k; i++ {
		a[len(a)-1-i] = 0
	}
	a = a[:len(a)-k]
	fmt.Println(a)

	// Output:
	//[1 2 3 4 5]
	//[1 2 3 5]
}
