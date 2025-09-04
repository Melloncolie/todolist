package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

func (todoStorage *TodoStorage) Import() (err error) {
	var (
		jsonData []byte
	)

	jsonData, err = os.ReadFile(todoStorage.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			todoStorage.TodoArray = &TodoArray{}
			return nil
		}
		return
	}

	err = json.Unmarshal(jsonData, todoStorage)

	if err != nil {
		todoStorage = nil
		return err
	}
	if todoStorage.TodoArray == nil {
		todoStorage.TodoArray = &TodoArray{}
	}

	return
}

func (todoStorage *TodoStorage) Export() (err error) {
	var (
		jsonData []byte
	)

	jsonData, err = json.Marshal(todoStorage)

	if err != nil {
		return
	}

	err = os.WriteFile(todoStorage.Filename, jsonData, 0644)

	return
}

func (todoStorage *TodoStorage) AppendAndExport(todoTitle, todoDescription string, tagID int) (todoPointer *TodoObject, err error) {
	var (
		tagPointer = &TodoTagUnit{}
		todoObject = TodoObject{}
	)

	if todoTitle == "" || todoDescription == "" {
		return nil, errors.New("All fields are not flush")
	}

	err = tagPointer.setTagUnit(tagID)
	if err != nil {
		return
	}

	todoObject = TodoObject{
		ID:          todoStorage.NextID,
		Title:       todoTitle,
		Description: todoDescription,
		Tag:         tagPointer,
	}
	now := time.Now()
	todoObject.TimeCreate.Time = &now
	todoObject.TimeUpdate = todoObject.TimeCreate

	*todoStorage.TodoArray = append(*todoStorage.TodoArray, todoObject)

	todoPointer = &todoObject
	todoStorage.NextID++

	todoStorage.Export()

	return
}

func (todoStorage *TodoStorage) UpdateAndExport(ID int, todoTitle, todoDescription string) (todoPointer *TodoObject, err error) {
	if todoTitle == "" && todoDescription == "" {
		return nil, errors.New("Fields are not flush")
	}

	todoPointer, err = todoStorage.TodoArray.getPointerID(ID)
	if err != nil {
		return nil, err
	}

	if todoTitle == "" {
		todoTitle = todoPointer.Title
	} else if todoDescription == "" {
		todoDescription = todoPointer.Description
	}

	todoObject := TodoObject{
		ID:          todoPointer.ID,
		Title:       todoTitle,
		Description: todoDescription,
		TimeCreate:  todoPointer.TimeCreate,
		Snapshot:    &TodoArray{},
	}

	todoPointer.update(todoObject)
	todoStorage.Export()
	return todoPointer, err
}

func (todoStorage *TodoStorage) SuccecssRecordAndExport(ID int) (todoPointer *TodoObject, err error) {
	todoPointer, err = todoStorage.getPointerID(ID)
	if err != nil {
		return nil, err
	}
	err = todoPointer.succecss()
	if err != nil {
		return
	}
	todoStorage.ComplTasks++
	err = todoStorage.Export()
	return
}

func (todoStorage *TodoStorage) RemoveAndExport(ID int) (todoPointer *TodoObject, err error) {
	todoPointer, err = todoStorage.TodoArray.getPointerID(ID)
	if err != nil {
		return nil, err
	}

	for i, todoObject := range *todoStorage.TodoArray {
		if todoObject.ID == todoPointer.ID {
			*todoStorage.TodoArray = append((*todoStorage.TodoArray)[:i], (*todoStorage.TodoArray)[i+1:]...)
			todoStorage.Export()
			return
		}
	}

	return nil, errors.New("Not found")
}
