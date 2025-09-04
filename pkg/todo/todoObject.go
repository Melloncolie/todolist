package todo

import (
	"errors"
	"time"
)

func (todoPointer *TodoObject) update(title, description string) {
	if todoPointer.Snapshot == nil {
		todoPointer.Snapshot = &TodoArray{}
	}
	*todoPointer.Snapshot = append(*todoPointer.Snapshot, *todoPointer)

	todoPointer.Title = title
	todoPointer.Description = description
	now := time.Now()
	todoPointer.TimeUpdate.Time = &now
}

func (todoPointer *TodoObject) succecss() (err error) {
	if todoPointer.Status {
		return errors.New("Task has alredy been completed")
	}
	todoPointer.Status = true
	now := time.Now()
	todoPointer.TimeUpdate.Time = &now
	return
}
