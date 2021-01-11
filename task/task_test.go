package task

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func ExampleTask_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
		nil,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"DONE","deadline":1439739780}
}

func ExampleTask_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":"DONE","Deadline":1439739780}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Buy Milk
	// DONE
	// 2015-08-16 15:43:00 +0000 UTC
}

//구조체가 아닌 자료형(배열, 맵)도 직렬화/역질력화 가능하다
func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]interface{}{"Name": "John", "Age": 16}) //아무형의 값을 담기 위해서 interface{} 자료형을 씀
	fmt.Println(string(b))
	// Output:
	// {"Age":16,"Name":"John"}
}

func ExampleTask_String() {
	fmt.Println(Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
		Priority: 0,
		SubTasks: []Task{{"Wash", DONE, nil, 0, nil}, {"Dry", DONE, nil, 0, nil}},
	})
	// Output: [v] Laundry <nil>
}

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}))
	// Output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Put <nil>
	//     [ ] Detergent <nil>
	//   [ ] Dry <nil>
	//   [ ] Fold <nil>
}

//구조체 내장을 이용하면 원래 구조체를 고치지 않고, 원하는 필드들만 제외하거나 추가하여 직렬화할 수 있음
func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields               //내장시킴
		InvisibleField string `json:"invisibleField,omitempty"`
		Additional     string `json:"additional,omitempty"`
	}{Fields: f, Additional: "c"})
	fmt.Println(string(b))

	//Output:{"visibleField":"a","additional":"c"}
}
