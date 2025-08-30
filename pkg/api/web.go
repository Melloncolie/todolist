package api

import (
	"encoding/json"
	"fmt"
	"go-todolist/pkg/todo"
	"net/http"
	"text/template"
)

func (api *Api) New() (err error) {
	api.todoArray = &todo.TodoArray{}
	err = api.todoArray.ReadFromFile()
	if err != nil {
		return
	}
	*api = NewApi(api.todoArray)
	return
}

func NewApi(todoArray *todo.TodoArray) Api {
	return Api{todoArray: todoArray}
}
func (api *Api) Init() (err error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("pkg/templates/index.html")
		if err != nil {
			http.Error(w, "error loading template", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		data := struct {
			Tasks *todo.TodoArray
		}{
			Tasks: api.todoArray,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	})

	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var tagID int
		fmt.Sscanf(record.Get("tag"), "%d", &tagID)
		_, err := api.todoArray.Insert(record.Get("title"), record.Get("description"), tagID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		_, err := api.todoArray.UpdateRecord(ID, record.Get("title"), record.Get("description"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/success/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		_, err := api.todoArray.SuccecssRecord(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/del/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		err := api.todoArray.Remove(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		todoArray, err := api.todoArray.Get(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(todoArray)
	})

	http.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		todoSearchArray, err := api.todoArray.Search(ID, record.Get("title"), record.Get("description"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(todoSearchArray)
	})

	http.HandleFunc("/filter/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		todoFilterArray, err := api.todoArray.Filter(record.Get("title"), record.Get("description"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(todoFilterArray)
	})

	return
}

func (api *Api) Run() (err error) {
	return http.ListenAndServe(":5000", nil)
}
