package todo

import (
	"errors"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// func (todoArray *TodoArray) Insert(todoTitle, todoDescription string, tagID int) (todoPointer *TodoObject, err error) {

// 	tagPointer := &TodoTagUnit{}
// 	if todoTitle == "" || todoDescription == "" {
// 		return nil, errors.New("All fields are not flush")
// 	}

// 	err = tagPointer.getTagUnit(tagID)
// 	if err != nil {
// 		return
// 	}

// 	return todoArray.insertViaObject(TodoObject{
// 		Title:       todoTitle,
// 		Description: todoDescription,
// 		Tag:         tagPointer,
// 	})
// }

// func (todoArray *TodoArray) insertViaObject(todoObject TodoObject) (todoPointer *TodoObject, err error) {
// 	var todoStorage TodoStorage
// 	_, todoObject.ID, _, err = todoStorage.getStorage()
// 	if err != nil {
// 		return
// 	}

// 	todoObject.TimeCreate = time.Now()
// 	todoObject.TimeUpdate = todoObject.TimeCreate
// 	*todoArray = append(*todoArray, todoObject)

// 	err = todoArray.addToFile()
// 	if err != nil {
// 		return
// 	}

// 	err = todoStorage.writeStorage()
// 	if err != nil {
// 		return
// 	}

// 	todoPointer = &(*todoArray)[len(*todoArray)-1]

// 	return
// }

func (todoArray *TodoArray) Filter(todoTitle, todoDescription string) (todoFilterArray *TodoArray) {
	todoFilterArray = &TodoArray{}

	for _, v := range *todoArray {
		if strings.Contains(v.Title, todoTitle) && strings.Contains(v.Description, todoDescription) {
			*todoFilterArray = append(*todoFilterArray, v)
		}
	}

	if len(*todoFilterArray) == 0 {
		return nil
	}

	return
}

func (todoArray *TodoArray) getPointerID(todoID int) (todoPointer *TodoObject, err error) {
	count := 0

	for i, v := range *todoArray {
		if v.ID == todoID {
			count++
			todoPointer = &(*todoArray)[i]
			if count > 1 {
				return nil, errors.New("To much Object")
			}
		}
	}

	if count > 0 {
		return
	}

	return nil, errors.New("Pointer not found")
}

// func (todoArray *TodoArray) UpdateRecord(ID int, todoTitle, todoDescription string) (todoPointer *TodoObject, err error) {
// 	if todoTitle == "" && todoDescription == "" {
// 		return nil, errors.New("Fields are not flush")
// 	}

// 	todoPointer, err = todoArray.getPointerID(ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if todoTitle == "" {
// 		todoTitle = todoPointer.Title
// 	} else if todoDescription == "" {
// 		todoDescription = todoPointer.Description
// 	}
// 	return todoArray.updateRecordViaObject(TodoObject{
// 		ID:          ID,
// 		Title:       todoTitle,
// 		Description: todoDescription,
// 	})
// }

// func (todoArray *TodoArray) updateRecordViaObject(todoObject TodoObject) (todoPointer *TodoObject, err error) {
// 	todoPointer, err = todoArray.getPointerID(todoObject.ID)
// 	if err != nil {
// 		return
// 	}

// 	todoPointer.update(todoObject)

// 	err = todoArray.addToArray(todoPointer)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = todoArray.addToFile()
// 	return
// }

// func (todoArray *TodoArray) succecssRecordViaObject(todoObject TodoObject) (todoPointer *TodoObject, err error) {
// 	todoPointer, err = todoArray.getPointerID(todoObject.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	todoPointer.succecss()

// 	err = todoArray.addToArray(todoPointer)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = todoArray.addToFile()
// 	return
// }

// func (todoArray *TodoArray) Remove(ID int) (err error) {
// 	todoPointer := &TodoObject{}
// 	todoPointer, err = todoArray.getPointerID(ID)
// 	if err != nil {
// 		return
// 	}
// 	return todoArray.removeViaObject(*todoPointer)
// }

// func (todoArray *TodoArray) removeViaObject(todoObject TodoObject) (err error) {
// 	for i, v := range *todoArray {
// 		if v.ID == todoObject.ID {
// 			*todoArray = append((*todoArray)[:i], (*todoArray)[i+1:]...)
// 			err = todoArray.addToFile()
// 			return
// 		}
// 	}
// 	return errors.New("Not found Object 'removeViaObject 154'")
// }

func (todoArray *TodoArray) Search(ID int, todoTitle, todoDescription string) (todoSearchArray *TodoArray) {
	todoSearchArray = &TodoArray{}
	tr := false
	if ID != -1 {
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
		return &TodoArray{}
	}
	return
}

func (todoArray *TodoArray) Get(ID int) (todoPointer *TodoObject, err error) {
	return todoArray.getPointerID(ID)
}

func (todoArray *TodoArray) RenderTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Title", "Description", "Create", "Update", "Status", "Tag")
	for _, v := range *todoArray {
		table.Append(v.ID, v.Title, v.Description, v.TimeCreate.String(), v.TimeUpdate.String, v.Status, v.Tag.Title)
	}
	table.Render()
}

// func (todoArray *TodoArray) addToArray(todoPointer *TodoObject) (err error) {
// 	for i, v := range *todoArray {
// 		if v.ID == (*todoPointer).ID {
// 			(*todoArray)[i] = *todoPointer
// 			return
// 		}
// 	}

// 	return errors.New("Not found in 'addToArray 224'")
// }
