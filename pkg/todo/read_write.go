package todo

import (
	"encoding/json"
	"errors"
	"os"
)

func (todoArray *TodoArray) addToFile() (err error) {
	var (
		todoStorage TodoStorage
		filename    string
	)
	filename, _, err = todoStorage.getStorage()
	if err != nil {
		return
	}
	var jsonData []byte
	jsonData, err = json.Marshal(todoArray)
	if err != nil {
		return errors.New("marshal " + filename + " error")
	}
	err = os.WriteFile("file/TodoList.json", jsonData, 0644)
	if err != nil {
		return errors.New("write " + filename + " error")
	}
	return
}

func (todoArray *TodoArray) ReadFromFile() (err error) {
	var (
		todoStorage TodoStorage
		filename    string
	)
	filename, _, err = todoStorage.getStorage()
	if err != nil {
		return
	}
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return errors.New("read " + filename + " error 'ReadFromFile() 41'")
	}
	err = json.Unmarshal(jsonData, &todoArray)
	if err != nil {
		return errors.New("unmarshal " + filename + " error 'ReadFromFile() 45'")
	}
	return
}
