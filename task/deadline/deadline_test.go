package deadline

import (
	"fmt"
	"time"
)

type Deadline time.Time

//Deadline이라는 자료형에 OverDue라는 메서드를 정의함
func (d Deadline) OverDue() bool {
	return time.Time(d).Before(time.Now())
}

func ExampleDeadline_OverDue() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))

	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())

	//Output:
	//true
	//false
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

//데드라인에이 없는 경우
func (d *Deadline) OverDueWithNoDeadline() bool {
	return d != nil && time.Time(*d).Before(time.Now())
}
func (t Task) OverDue() bool {
	return t.Deadline.OverDueWithNoDeadline()
}

func ExampleDeadline_OverDue_With_NoDeadline() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	t1 := Task{"4h ago", TODO, &d1}
	t2 := Task{"4h later", TODO, &d2}
	t3 := Task{"no duer", TODO, nil}

	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())

	//Output:
	//true
	//false
	//false
}
