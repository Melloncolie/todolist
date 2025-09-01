package todo

import (
	"encoding/json"
	"errors"
	"os"
)

func (todoStorage *TodoStorage) getStorage() (filename string, lastID int, err error) {
	var jsonData []byte
	jsonData, err = os.ReadFile("file/Storage.json")
	if err != nil {
		return "", 0, errors.New("read file/Storage.json error")
	}
	err = json.Unmarshal(jsonData, &todoStorage)
	if err != nil {
		return "", 0, errors.New("unmarshal file/Storage.json error")
	}
	return todoStorage.Filename, todoStorage.LastID, err
}

func (todoStorage *TodoStorage) writeStorage() (err error) {
	var jsonData []byte
	todoStorage.LastID++
	jsonData, err = json.Marshal(todoStorage)
	if err != nil {
		return errors.New("marshal file/Storage.json error")
	}
	err = os.WriteFile("file/Storage.json", jsonData, 0644)
	if err != nil {
		return errors.New("write file/Storage.json error")
	}
	return
}
