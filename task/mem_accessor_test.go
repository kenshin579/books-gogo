package task

import (
	"log"
	"testing"
	"time"
)

func Test_Post(t *testing.T) {
	inMemoryAccessor := NewInMemoryAccessor()
	t1 := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
		nil,
	}

	id, err := inMemoryAccessor.Post(t1)
	if err != nil {
		log.Fatal("error occurred while creating task", err)
	}
	//assert.Equal(t, interface{}(id).(string), "1")
	if id != "1" {
		t.Errorf("got %s, want %s", id, "1")
	}

	task, err := inMemoryAccessor.Get(id)
	if err != nil {
		log.Fatal("error occurred while getting task", err)
	}
	if task.Title != t1.Title {
		t.Errorf("got %s, want %s", id, t1.Title)
	}
}
