package queue

import "fmt"

func Example_queue() {
	var q *Queue = new(Queue)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Size())
	//Output:
	//2
	//3
	//1
}
