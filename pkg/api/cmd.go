package api

import (
	"flag"
	"log"

	"github.com/k0kubun/pp"
)

func (api *Api) CMD() (err error) {
	var (
		cursor = api.todoStorage.TodoArray
	)

	ID := flag.Int("id", 0, "")
	title := flag.String("title", "", "")
	description := flag.String("description", "", "")
	tagID := flag.Int("tag", -1, "")

	add := flag.Bool("add", false, "")
	search := flag.Bool("search", false, "")
	del := flag.Bool("del", false, "")
	update := flag.Bool("update", false, "")
	filter := flag.Bool("filter", false, "")
	get := flag.Bool("get", false, "")
	view := flag.Bool("view", false, "")
	demon := flag.Bool("demon", false, "")
	flag.Parse()

	if *add {
		_, err = api.todoStorage.TodoArray.Insert(*title, *description, *tagID)
		if err != nil {
			return
		}
	} else if *search {
		cursor = api.todoStorage.TodoArray.Search(*ID, *title, *description)
	} else if *del {
		api.todoStorage.Remove(*ID)
	} else if *update {
		_, err = api.todoStorage.TodoArray.UpdateRecord(*ID, *title, *description)
		if err != nil {
			return
		}
	} else if *filter {
		cursor = api.todoStorage.TodoArray.Filter(*title, *description)
	} else if *view {
		err = api.todoStorage.TodoArray.ReadFromFile()
		if err != nil {
			return
		}
	} else if *get {
		record, err := api.todoStorage.TodoArray.Get(*ID)
		if err != nil {
			return err
		}
		pp.Println(record)
	}

	if *demon {
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

	if err != nil {
		return err
	}

	cursor.RenderTable()

	return
}
