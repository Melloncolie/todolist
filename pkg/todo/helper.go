package todo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func getTagUnit(tagID int) (tagPointer *TodoTagUnit, err error) {
	tagPointer = &TodoTagUnit{}
	switch tagID {
	case 0:
		tagPointer.Title = "Home"
		tagPointer.Class = "Green"
	case 1:
		tagPointer.Title = "Work"
		tagPointer.Class = "Red"
	case 2:
		tagPointer.Title = "Default"
		tagPointer.Class = "Blue"
	default:
		return nil, errors.New("unknow ID")
	}
	return
}
func (todoArray *TodoArray) addToArray(todoPointer *TodoObject) (err error) {
	for i, v := range *todoArray {
		if v.ID == (*todoPointer).ID {
			(*todoArray)[i] = *todoPointer
			return
		}
	}

	return errors.New("not found")
}

func getID() (ID int, err error) {
	input, err := os.Open("pkg/file/ID.csv")
	if err != nil {
		return 0, errors.New("open error")
	}
	defer input.Close()
	reader := csv.NewReader(input)
	data, err := reader.Read()
	if err != nil {
		return 0, errors.New("read error")
	}
	fmt.Sscanf(data[0], "%d", &ID)

	output, err := os.Create("pkg/file/ID.csv")
	if err != nil {
		return 0, errors.New("write error")
	}
	defer output.Close()

	writer := csv.NewWriter(output)

	data = []string{fmt.Sprintf("%d", ID+1)}
	err = writer.Write(data)
	if err != nil {
		return 0, errors.New("write error")
	}
	writer.Flush()
	return
}
