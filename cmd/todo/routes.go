package main

import "net/http"

//Routes : generate all routes
func (app *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/todo", app.Logger(app.Todo))
	mux.HandleFunc("/", app.Logger(app.Home))
	return mux
}
