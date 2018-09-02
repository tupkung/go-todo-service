package main

import "net/http"

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
	app.writeAppJSONHeader(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{'data':[{'title':'Task-1','complete':false},{'title':'Task-2','complete':false}]}"))
}

func (app *App) todoPost(w http.ResponseWriter, r *http.Request) {
	app.writeAppJSONHeader(w, r)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{'data':[],}"))
}

func (app *App) todoPut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (app *App) todoDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
