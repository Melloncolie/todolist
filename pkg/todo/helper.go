package todo

import (
	"fmt"
	"time"
)

type Object struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	TimeCreate  string       `json:"timeCreate"`
	TimeUpdate  string       `json:"timeUpdate"`
	Status      bool         `json:"status"`
	Snapshot    Array        `json:"snapshot"`
	Tag         *TodoTagUnit `json:"tag"`
}

type Array []Object

func String(times time.Time) string {
	return times.Format("2006-01-02 15:04:05")
}

func ToTime(str string) (t time.Time, err error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func (todoArray *TodoArray) StringArray() (array Array) {
	array = make(Array, len(*todoArray))
	for i, v := range *todoArray {
		array[i].ID = v.ID
		array[i].Title = v.Title
		array[i].Description = v.Description
		array[i].TimeCreate = String(v.TimeCreate)
		array[i].TimeUpdate = String(v.TimeUpdate)
		array[i].Status = v.Status
		if v.Snapshot != nil {
			array[i].Snapshot = v.Snapshot.StringArray()
		}
		array[i].Tag = v.Tag
	}
	return
}

func (array *Array) Todoarray() (todoArray *TodoArray, err error) {
	todoArray = &TodoArray{}
	*todoArray = make(TodoArray, len(*array))
	for i, v := range *array {
		(*todoArray)[i].ID = v.ID
		(*todoArray)[i].Title = v.Title
		(*todoArray)[i].Description = v.Description
		(*todoArray)[i].TimeCreate, err = ToTime(v.TimeCreate)
		if err != nil {
			return nil, fmt.Errorf("Wrong time format in ID %d: %v", v.ID, err)
		}
		(*todoArray)[i].TimeUpdate, err = ToTime(v.TimeUpdate)
		if err != nil {
			return nil, fmt.Errorf("Wrong time format in ID %d: %v", v.ID, err)
		}
		(*todoArray)[i].Status = v.Status
		if v.Snapshot != nil {
			(*todoArray)[i].Snapshot, err = v.Snapshot.Todoarray()
			if err != nil {
				return
			}
		}
		(*todoArray)[i].Tag = v.Tag
	}
	return
}
