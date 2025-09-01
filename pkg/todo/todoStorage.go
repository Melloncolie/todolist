package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func (todoStorage *TodoStorage) getStorage() (filename string, lastID int, complTask int, err error) {
	var jsonData []byte
	jsonData, err = os.ReadFile("file/Storage.json")
	if err != nil {
		return "", 0, 0, fmt.Errorf("Read JSON file/Storage.json has error: %v", err)
	}
	err = json.Unmarshal(jsonData, &todoStorage)
	if err != nil {
		return "", 0, 0, fmt.Errorf("Unmarshal JSON file/Storage.json has error: %v", err)
	}
	return todoStorage.Filename, todoStorage.LastID, todoStorage.ComplTasks, err
}

func (todoStorage *TodoStorage) writeStorage() (err error) {
	var jsonData []byte
	todoStorage.LastID++
	jsonData, err = json.Marshal(todoStorage)
	if err != nil {
		return fmt.Errorf("Marshal JSON file/Storage.json has error: %v", err)
	}
	err = os.WriteFile("file/Storage.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Write JSON file/Storage.json has error: %v", err)
	}
	return
}

func (todoStorage *TodoStorage) Import() (err error) {
	todoStorage.Filename, todoStorage.LastID, todoStorage.ComplTasks, err = todoStorage.getStorage()
	if err != nil {
		return
	}

	todoStorage.TodoArray = &TodoArray{}
	err = todoStorage.TodoArray.ReadFromFile()
	if err != nil {
		return
	}
	return
}

func (todoStorage *TodoStorage) SuccecssRecord(ID int) (todoPointer *TodoObject, err error) {
	todoStorage.ComplTasks++
	todoStorage.writeStorage()
	return todoStorage.TodoArray.succecssRecordViaObject(TodoObject{ID: ID})
}
