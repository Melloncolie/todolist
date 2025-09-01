package todo

import "time"

func (todoPointer *TodoObject) update(todoObject TodoObject) {
	todoObject.Snapshot = todoPointer.Snapshot
	todoObject.TimeCreate = todoPointer.TimeCreate
	todoObject.Tag = todoPointer.Tag
	todoPointer.Snapshot = nil
	*todoObject.Snapshot = append(*todoObject.Snapshot, *todoPointer)
	*todoPointer = todoObject
	todoPointer.TimeUpdate = time.Now()
}

func (todoPointer *TodoObject) succecss() {
	todoPointer.Status = true
	todoPointer.TimeUpdate = time.Now()
}
