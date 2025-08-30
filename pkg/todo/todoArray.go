package todo

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

func (todoArray *TodoArray) Insert(todoTitle, todoDescription string, tagID int) (todoPointer *TodoObject, err error) {
	var (
		tagPointer *TodoTagUnit
	)

	if todoTitle == "" || todoDescription == "" {
		return nil, errors.New("all fields are not flush")
	}

	tagPointer, err = getTagUnit(tagID)
	if err != nil {
		return
	}

	return todoArray.insertViaObject(TodoObject{
		Title:       todoTitle,
		Description: todoDescription,
		Tag:         tagPointer,
	})
}

func (todoArray *TodoArray) insertViaObject(todoObject TodoObject) (todoPointer *TodoObject, err error) {
	todoObject.ID, err = getID()
	if err != nil {
		return
	}

	todoObject.TimeCreate = time.Now()
	todoObject.TimeUpdate = todoObject.TimeCreate
	*todoArray = append(*todoArray, todoObject)
	todoPointer = &todoObject

	err = todoArray.addToFile()
	if err != nil {
		todoPointer = nil
		return
	}

	return
}

func (todoArray *TodoArray) Filter(todoTitle, todoDescription string) (todoFilterArray *TodoArray, err error) {
	todoFilterArray = &TodoArray{}

	for _, v := range *todoArray {
		if strings.Contains(v.Title, todoTitle) && strings.Contains(v.Description, todoDescription) {
			*todoFilterArray = append(*todoFilterArray, v)
		}
	}

	if len(*todoFilterArray) == 0 {
		return nil, errors.New("not found")
	}

	return
}

func (todoArray *TodoArray) getPointer(todoObject TodoObject) (todoPointer *TodoObject, err error) {
	count := 0

	for _, v := range *todoArray {
		if v.ID == todoObject.ID {
			count++
			todoPointer = &v
			if count > 1 {
				return nil, errors.New("")
			}
		}
	}

	if count > 0 {
		return
	}

	return nil, errors.New("not found")
}

func (todoArray *TodoArray) UpdateRecord(ID int, todoTitle, todoDescription string) (todoPointer *TodoObject, err error) {
	return todoArray.updateRecordViaObject(TodoObject{
		ID:          ID,
		Title:       todoTitle,
		Description: todoDescription,
	})
}

func (todoArray *TodoArray) updateRecordViaObject(todoObject TodoObject) (todoPointer *TodoObject, err error) {
	todoPointer, err = todoArray.getPointer(todoObject)
	if err != nil {
		return
	}

	err = todoPointer.update(todoObject)
	if err != nil {
		return nil, err
	}

	err = todoArray.addToArray(todoPointer)
	if err != nil {
		return nil, err
	}
	err = todoArray.addToFile()
	return
}

func (todoArray *TodoArray) SuccecssRecord(ID int) (todoPointer *TodoObject, err error) {
	return todoArray.succecssRecordViaObject(TodoObject{ID: ID})
}

func (todoArray *TodoArray) succecssRecordViaObject(todoObject TodoObject) (todoPointer *TodoObject, err error) {
	todoPointer, err = todoArray.getPointer(todoObject)
	if err != nil {
		return nil, err
	}
	err = todoPointer.succecss()
	if err != nil {
		return nil, err
	}
	err = todoArray.addToArray(todoPointer)
	if err != nil {
		return nil, err
	}

	err = todoArray.addToFile()
	return
}

func (todoArray *TodoArray) Remove(ID int) (err error) {
	todoPointer, err := todoArray.getPointer(TodoObject{ID: ID})
	if err != nil {
		return
	}
	return todoArray.removeViaObject(*todoPointer)
}

func (todoArray *TodoArray) removeViaObject(todoObject TodoObject) (err error) {
	for i, v := range *todoArray {
		if v.ID == todoObject.ID {
			*todoArray = append((*todoArray)[:i], (*todoArray)[i+1:]...)
			err = todoArray.addToFile()
			return
		}
	}
	return errors.New("not found")
}

func (todoArray *TodoArray) Search(ID int, todoTitle, todoDescription string) (todoSearchArray *TodoArray, err error) {
	todoSearchArray = &TodoArray{}
	tr := false
	if ID != 0 {
		for _, v := range *todoArray {
			if v.ID == ID && (todoDescription == v.Description || todoDescription == "") && (todoTitle == v.Title || todoTitle == "") {
				*todoSearchArray = append(*todoSearchArray, v)
				tr = true
			}
		}
	} else if todoTitle != "" {
		for _, v := range *todoArray {
			if v.Title == todoTitle && (todoDescription == v.Description || todoDescription == "") {
				*todoSearchArray = append(*todoSearchArray, v)
				tr = true
			}
		}
	} else if todoDescription != "" {
		for _, v := range *todoArray {
			if todoDescription == v.Description {
				*todoSearchArray = append(*todoSearchArray, v)
				tr = true
			}
		}
	}
	if !tr {
		return &TodoArray{}, errors.New("not found")
	}
	return
}

func (todoArray *TodoArray) Get(ID int) (todoPointer *TodoObject, err error) {
	return todoArray.getPointer(TodoObject{ID: ID})
}

func (todoArray *TodoArray) RenderTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header("Title", "Description", "Create", "Update", "Status", "Tag")
	for _, v := range *todoArray {
		table.Append(v.ID, v.Description, v.TimeCreate, v.TimeUpdate, v.Status, v.Tag.Title)
	}
	table.Render()
}
