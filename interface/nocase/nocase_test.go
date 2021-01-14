package nocase

import (
	"container/heap"
	"fmt"
	"sort"
)

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	sort.Sort(apple) //todo : Interface인데 slice를 넘겨줌 - 인터페이스 스터디가 더 필요함
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MacBook]
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}

func ExampleCaseInsensitive_heapString() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}
