package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-todolist/pkg/todo"
	"log"
	"net/http"
	"text/template"
)

func (api *Api) New() (err error) {
	api.todoStorage = &todo.TodoStorage{}
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
func (api *Api) Init() (err error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "error loading template", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		data := struct {
			Tasks      *todo.TodoArray
			LastID     int
			ComplTasks int
		}{
			Tasks:      api.todoStorage.TodoArray,
			LastID:     api.todoStorage.LastID,
			ComplTasks: api.todoStorage.ComplTasks,
		}
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		_, err = buf.WriteTo(w)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var tagID int
		fmt.Sscanf(record.Get("tag"), "%d", &tagID)
		todoPointer, err := api.todoStorage.TodoArray.Insert(record.Get("title"), record.Get("description"), tagID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Println("Task added")
		fmt.Printf("Task: %+v\n", todoPointer)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		todoPointer, err := api.todoStorage.TodoArray.UpdateRecord(ID, record.Get("title"), record.Get("description"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Println("Task updated")
		fmt.Printf("Task: %+v\n", todoPointer)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/success/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		todoPointer, err := api.todoStorage.SuccecssRecord(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Println("Task complited")
		fmt.Printf("Task: %+v\n", todoPointer)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/del/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		err := api.todoStorage.TodoArray.Remove(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println("Task removed")
	})

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		todoPointer, err := api.todoStorage.TodoArray.Get(ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		json.NewEncoder(w).Encode(todoPointer)
		log.Print("Task received ")
		fmt.Printf("Task: %+v\n", todoPointer)
	})

	http.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		var ID int
		fmt.Sscanf(record.Get("id"), "%d", &ID)
		todoSearchArray := api.todoStorage.TodoArray.Search(ID, record.Get("title"), record.Get("description"))
		json.NewEncoder(w).Encode(todoSearchArray)
		log.Printf("Search completed. Found %d tasks", len(*todoSearchArray))
	})

	http.HandleFunc("/filter/", func(w http.ResponseWriter, r *http.Request) {
		record := r.URL.Query()
		todoFilterArray := api.todoStorage.TodoArray.Filter(record.Get("title"), record.Get("description"))
		json.NewEncoder(w).Encode(todoFilterArray)
		log.Printf("Filtering completed. Found %d tasks.", len(*todoFilterArray))
	})

	return
}

var port string = ":5000"

func (api *Api) Run() (err error) {
	log.Println("Server is running")
	return http.ListenAndServe(port, nil)
}
