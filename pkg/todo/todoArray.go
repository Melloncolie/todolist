package todo

import (
	"errors"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

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

func (todoArray *TodoArray) Search(ID int, todoTitle, todoDescription string) *TodoArray {
	result := &TodoArray{}

	for _, v := range *todoArray {
		matchID := ID == -1 || v.ID == ID
		matchTitle := todoTitle == "" || v.Title == todoTitle
		matchDesc := todoDescription == "" || v.Description == todoDescription

		if matchID && matchTitle && matchDesc {
			*result = append(*result, v)
		}
	}

	return result
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
