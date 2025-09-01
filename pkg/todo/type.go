package todo

import "time"

type TodoObject struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	TimeCreate  time.Time    `json:"timeCreate"`
	TimeUpdate  time.Time    `json:"timeUpdate"`
	Status      bool         `json:"status"`
	Snapshot    TodoArray    `json:"snapshot"`
	Tag         *TodoTagUnit `json:"tag"`
}

type TodoTagUnit struct {
	ID    int
	Title string
	Class string
}

type TodoArray []TodoObject

type TodoStorage struct {
	Filename string `json:"filename"`
	LastID   int    `json:"id"`
}
