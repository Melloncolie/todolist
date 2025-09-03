package api

import "go-todolist/pkg/todo"

func (api *Api) New() (err error) {
	err = api.todoStorage.Import()
	if err != nil {
		return
	}
	*api = NewApi(api.todoStorage)
	return
}

func NewApi(todoStorage *todo.TodoStorage) Api {
	return Api{todoStorage: todoStorage}
}
