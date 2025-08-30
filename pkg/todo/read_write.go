package todo

import (
	"encoding/json"
	"errors"
	"os"
)

func (todoArray *TodoArray) addToFile() (err error) {
	jsonData, err := json.Marshal(todoArray)
	if err != nil {
		return errors.New("marshal error")
	}
	err = os.WriteFile("pkg/file/TodoList.json", jsonData, 0644)
	if err != nil {
		return errors.New("write error")
	}
	return
}

func (todoArray *TodoArray) ReadFromFile() (err error) {

	jsonData, err := os.ReadFile("pkg/file/TodoList.json")
	if err != nil {
		return errors.New("read error")
	}
	err = json.Unmarshal(jsonData, &todoArray)
	if err != nil {
		return errors.New("unmarshal error")
	}
	return
}
