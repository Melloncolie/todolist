package api

import (
	"flag"
	"go-todolist/pkg/todo"
	"log"

	"github.com/k0kubun/pp"
)

func (api *Api) CMD() (err error) {

	ID := flag.Int("id", -1, "")
	title := flag.String("title", "", "")
	description := flag.String("description", "", "")
	tagID := flag.Int("tag", -1, "")
	filename := flag.String("filename", "file/Storage.json", "")

	add := flag.Bool("add", false, "")
	succecss := flag.Bool("succecss", false, "")
	search := flag.Bool("search", false, "")
	del := flag.Bool("del", false, "")
	update := flag.Bool("update", false, "")
	filter := flag.Bool("filter", false, "")
	get := flag.Bool("get", false, "")
	port := flag.String("port", "", "")
	flag.Parse()

	api.todoStorage = &todo.TodoStorage{}

	err = api.New()
	if err != nil {
		return
	}

	var (
		cursor      = api.todoStorage.TodoArray
		todoPointer = &todo.TodoObject{}
	)

	api.port = *port
	api.todoStorage.Filename = *filename

	if api.port != "" {
		err = api.Init()
		if err != nil {
			return
		}

		err = api.Run()
		if err != nil {
			return
		}
	} else if *add {
		todoPointer, err = api.todoStorage.AppendAndExport(*title, *description, *tagID)
		if err != nil {
			return
		}
		log.Print("Task added ")
		pp.Println(todoPointer)
	} else if *search {
		cursor = api.todoStorage.TodoArray.Search(*ID, *title, *description)
	} else if *del {
		todoPointer, err = api.todoStorage.RemoveAndExport(*ID)
		if err != nil {
			return
		}
		log.Print("Task deleted ")
		pp.Println(todoPointer)
	} else if *update {
		todoPointer, err = api.todoStorage.UpdateAndExport(*ID, *title, *description)
		if err != nil {
			return
		}
		log.Print("Task has been updated ")
		pp.Println(todoPointer)
	} else if *succecss {
		todoPointer, err = api.todoStorage.SuccecssRecordAndExport(*ID)
		if err != nil {
			return
		}
		log.Print("Task is complited ")
		pp.Println(todoPointer)

	} else if *filter {
		cursor = api.todoStorage.TodoArray.Filter(*title, *description)
	} else if *get {
		todoPointer, err = api.todoStorage.TodoArray.Get(*ID)
		if err != nil {
			return err
		}
		pp.Println(todoPointer)
	}

	if err != nil {
		return err
	}

	cursor.RenderTable()

	return
}
