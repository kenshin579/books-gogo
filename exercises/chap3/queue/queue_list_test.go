package queue

import (
	"container/list"
	"fmt"
)

func Example() {
	customQueue := &customQueue{
		queue: list.New(),
	}
	customQueue.Enqueue("A")
	customQueue.Enqueue("B")
	fmt.Printf("Size: %d\n", customQueue.Size())
	for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Dequeue: %s\n", frontVal)
		customQueue.Dequeue()
	}
	fmt.Printf("Size: %d\n", customQueue.Size())

	//Output:
	//Size: 2
	//Front: A
	//Dequeue: A
	//Front: B
	//Dequeue: B
	//Size: 0
}

func Example_slice_test() {
	var queue []string

	queue = append(queue, "Hello") // Enqueue
	queue = append(queue, "world!")

	fmt.Println(queue[0]) // First element
	queue[0] = ""         //Erase element (write zero value) - 이렇게 하지 않으면 memory leak가 발생할 수 있음
	queue = queue[1:]     // Dequeue
	fmt.Println(queue[0]) // First element

	//Output:
	//Hello
	//world!
}
