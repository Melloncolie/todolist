package todo

import "time"

type TodoObject struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	TimeCreate  TodoTime     `json:"timeCreate"`
	TimeUpdate  TodoTime     `json:"timeUpdate"`
	Status      bool         `json:"status"`
	Snapshot    *TodoArray   `json:"snapshot"`
	Tag         *TodoTagUnit `json:"tag"`
}

type TodoTime struct {
	*time.Time
}

type TodoTagUnit struct {
	ID    int
	Title string
	Class string
}

type TodoArray []TodoObject

type TodoStorage struct {
	Filename   string `json:"filename"`
	NextID     int    `json:"id"`
	ComplTasks int    `json:"complTasks"`
	*TodoArray `json:"todoArray"`
}
