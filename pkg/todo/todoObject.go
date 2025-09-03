package todo

import (
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
	todoPointer.TimeUpdate = time.Now()
}

func (todoPointer *TodoObject) succecss() {
	todoPointer.Status = true
	todoPointer.TimeUpdate = time.Now()
}
