package todo

import (
	"errors"
	"time"
)

func (todoPointer *TodoObject) update(todoObject TodoObject) {
	if todoPointer.Snapshot == nil {
		todoPointer.Snapshot = &TodoArray{}
	}
	todoObject.Snapshot = todoPointer.Snapshot
	todoPointer.Snapshot = nil
	*todoObject.Snapshot = append(*todoObject.Snapshot, *todoPointer)

	todoPointer.Title = todoObject.Title
	todoPointer.Description = todoObject.Description

	todoObject.TimeCreate = todoPointer.TimeCreate
	todoObject.Tag = todoPointer.Tag
	*todoPointer = todoObject
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
