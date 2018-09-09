package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/tupkung/go-todo-service/pkg/forms"
	"github.com/tupkung/go-todo-service/pkg/models"
)

//Logger middleware
func (app *App) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer app.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

//Home for serving info content
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Todo Service"))
}

//Todo for serving tasks list
func (app *App) Todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.todoGet(w, r)
	case "POST":
		app.todoPost(w, r)
	case "PUT":
		app.todoPut(w, r)
	case "DELETE":
		app.todoDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (app *App) todoGet(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.dataBase.LatestTasks()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		app.logger.Printf("%v", err)
		return
	}
	app.writeAppJSONHeader(w, r)
	w.WriteHeader(http.StatusOK)
	type resultObj struct {
		Data models.Tasks `json:"data"`
	}
	result, _ := json.Marshal(&resultObj{
		Data: tasks,
	})
	w.Write([]byte(result))
	// w.Write([]byte("{\"data\":[{\"title\":\"Task-1\",\"complete\":false},{\"title\":\"Task-2\",\"complete\":false}]}"))
}

func (app *App) todoPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	form := &forms.NewTask{
		Title: r.PostForm.Get("title"),
	}

	if !form.Valid() {
		fmt.Fprint(w, form.Failures)
		return
	}

	uid, err := app.dataBase.InsertTask(form.Title)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	app.writeAppJSONHeader(w, r)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(uid))
}

func (app *App) todoPut(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	complete, err := strconv.ParseBool(r.PostForm.Get("complete"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	form := &forms.EditTask{
		UID:      r.PostForm.Get("uid"),
		Title:    r.PostForm.Get("title"),
		Complete: complete,
	}

	err = app.dataBase.UpdateTask(form.UID, form.Title, form.Complete)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (app *App) todoDelete(w http.ResponseWriter, r *http.Request) {

	uid := r.URL.Query().Get("uid")

	err := app.dataBase.DeleteTask(uid)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
