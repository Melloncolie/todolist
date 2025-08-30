package api

import "go-todolist/pkg/todo"

type ApiInterface interface {
	Init() error
	Run() error
	New() error
}

type Api struct {
	todoArray *todo.TodoArray
}
