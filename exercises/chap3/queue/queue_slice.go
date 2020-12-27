package queue

//https://www.calhoun.io/lets-learn-algorithms-queues-in-go/
type Queue struct {
	list []int
}

func (q *Queue) Enqueue(i int) {
	q.list = append(q.list, i)
}

func (q *Queue) Dequeue() int {
	var value = q.list[0]
	q.list = q.list[1:]
	return value
}

func (q *Queue) Size() int {
	return len(q.list)
}
