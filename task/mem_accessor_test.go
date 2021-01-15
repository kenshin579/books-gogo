package task

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Post_Get(t *testing.T) {
	inMemoryAccessor := NewInMemoryAccessor()
	t1 := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
		nil,
	}

	id, _ := inMemoryAccessor.Post(t1)
	//assert.Equal(t, interface{}(id).(string), "1")
	if id != "1" {
		t.Errorf("got %s, want %s", id, "1")
	}

	task, _ := inMemoryAccessor.Get(id)
	if task.Title != t1.Title {
		t.Errorf("got %s, want %s", id, t1.Title)
	}
}

func TestGet_없는_id로_조홰시_오류가_발생한다(t *testing.T) {
	inMemoryAccessor := NewInMemoryAccessor()
	id := ID("1")

	_, err := inMemoryAccessor.Get(id)
	assert.EqualError(t, err, "task does not exist")
}
