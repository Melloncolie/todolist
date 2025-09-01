package main

import (
	"go-todolist/pkg/api"
	"log"
)

func main() {

	var api api.Api
	err := api.New()
	if err != nil {
		log.Println(err)
		return
	}

	err = api.CMD()
	if err != nil {
		log.Println(err)
		return
	}
	err = api.Init()
	if err != nil {
		log.Println(err)
		return
	}

	err = api.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
