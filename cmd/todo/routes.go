package main

import "net/http"

//Routes : generate all routes
func (app *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/todo", app.Todo)
	mux.HandleFunc("/", app.Home)
	return mux
}
