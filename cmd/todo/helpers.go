package main

import "net/http"

func (app *App) writeAppJSONHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
