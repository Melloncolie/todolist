package main

import (
	"fmt"
	"go-todolist/pkg/api"
	"go-todolist/pkg/todo"
)

func main() {

	var api api.Api
	err := api.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = api.CMD()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = api.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = api.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	va := todo.TodoStorage{}

	va.Filename = "123"
	va.TodoArray.Insert("", "", 1)

}
