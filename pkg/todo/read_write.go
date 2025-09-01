package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func (todoArray *TodoArray) addToFile() (err error) {
	var (
		todoStorage TodoStorage
		filename    string
		array       Array
		jsonData    []byte
	)

	filename, _, _, err = todoStorage.getStorage()
	if err != nil {
		return
	}
	array = todoArray.StringArray()
	jsonData, err = json.Marshal(array)
	if err != nil {
		return fmt.Errorf("Marshal JSON %s has error: %v", filename, err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Write JSON %s has error: %v", filename, err)
	}
	return
}

func (todoArray *TodoArray) ReadFromFile() (err error) {
	var (
		todoStorage TodoStorage
		filename    string
		array       Array
		jsonData    []byte
		parse       *TodoArray
	)

	filename, _, _, err = todoStorage.getStorage()
	if err != nil {
		return
	}

	jsonData, err = os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Read JSON %s has error: %v", filename, err)
	}

	err = json.Unmarshal(jsonData, &array)
	if err != nil {
		return fmt.Errorf("Unmarshal JSON %s has error: %v", filename, err)
	}

	parse, err = array.Todoarray()
	if err != nil {
		return
	}
	*todoArray = *parse

	return
}
