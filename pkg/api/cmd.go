package api

import (
	"flag"

	"github.com/k0kubun/pp"
)

func (api *Api) CMD() (err error) {
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
	flag.Parse()

	if *add {
		_, err = api.todoArray.Insert(*title, *description, *tagID)
		if err != nil {
			return
		}
		api.todoArray.RenderTable()
	} else if *search {
		todoSearchArray, err := api.todoArray.Search(*ID, *title, *description)
		if err != nil {
			return err
		}
		todoSearchArray.RenderTable()
	} else if *del {
		api.todoArray.Remove(*ID)
	} else if *update {
		_, err = api.todoArray.UpdateRecord(*ID, *title, *description)
		if err != nil {
			return
		}
		api.todoArray.RenderTable()
	} else if *filter {
		todoFilterArray, err := api.todoArray.Filter(*title, *description)
		if err != nil {
			return err
		}
		todoFilterArray.RenderTable()
	} else if *view {
		err = api.todoArray.ReadFromFile()
		if err != nil {
			return
		}
		api.todoArray.RenderTable()
	} else if *get {
		record, err := api.todoArray.Get(*ID)
		if err != nil {
			return err
		}
		pp.Println(record)
	}
	return
}
