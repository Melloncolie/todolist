package api

import (
	"flag"
	"go-todolist/pkg/todo"
)

func (api *Api) CMD() (err error) {
	// var (
	// 	cursor = api.todoStorage.TodoArray
	// )

	ID := flag.Int("id", -1, "")
	title := flag.String("title", "", "")
	description := flag.String("description", "", "")
	tagID := flag.Int("tag", -1, "")

	add := flag.Bool("add", false, "")
	filename := flag.String("filename", "file/Storage.json", "")
	// search := flag.Bool("search", false, "")
	// del := flag.Bool("del", false, "")
	update := flag.Bool("update", false, "")
	// filter := flag.Bool("filter", false, "")
	// get := flag.Bool("get", false, "")
	// view := flag.Bool("view", false, "")
	// demon := flag.Bool("demon", false, "")
	flag.Parse()

	api.todoStorage = &todo.TodoStorage{}
	api.todoStorage.Filename = *filename
	err = api.New()
	if err != nil {
		return
	}

	if *add {
		_, err = api.todoStorage.AppendAndExport(*title, *description, *tagID)
		if err != nil {
			return
		}

		// } else if *search {
		// 	cursor = api.todoStorage.TodoArray.Search(*ID, *title, *description)
		// } else if *del {
		// 	api.todoStorage.Remove(*ID)
	} else if *update {
		_, err = api.todoStorage.UpdateAndExport(*ID, *title, *description)
		if err != nil {
			return
		}
	}
	// } else if *filter {
	// 	cursor = api.todoStorage.TodoArray.Filter(*title, *description)
	// } else if *view {
	// 	err = api.todoStorage.TodoArray.ReadFromFile()
	// 	if err != nil {
	// 		return
	// 	}
	// } else if *get {
	// 	record, err := api.todoStorage.TodoArray.Get(*ID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	pp.Println(record)
	// }

	// if *demon {
	// 	err = api.Init()
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}

	// 	err = api.Run()
	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// }

	if err != nil {
		return err
	}

	api.todoStorage.TodoArray.RenderTable()

	return
}
