package todo

import (
	"math/rand"
	"time"
)

// Todo struct
type Todo struct {
	ID          uint64    `json:",omitempty"`
	Task        string    `json:""`
	IsCompleted bool      `json:""`
	CreatedAt   time.Time `json:",omitempty"`
}

// NewTodo creates a new Todo struct and sets the CreatedAt field to the current time
func NewTodo() *Todo {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Todo{ID: rand.Uint64(), CreatedAt: time.Now()}
}

// Update updates a the Task and IsCompleted fields of a Todo struct
func (t *Todo) Update(s string, i bool) {
	t.Task = s
	if t.IsCompleted != i {
		t.IsCompleted = i
	}
}
