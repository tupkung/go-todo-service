package main

import (
	"log"
	"os"

	"github.com/tupkung/go-todo-service/server"
)

var (
	todoServiceAddr     = os.Getenv("TODO_SERVICE_ADDR")
	todoServiceCertFile = os.Getenv("TODO_SERVICE_CERT_FILE")
	todoServiceKeyFile  = os.Getenv("TODO_SERVICE_KEY_FILE")
)

func main() {
	app := &App{
		todoServiceAddr:     todoServiceAddr,
		todoServiceCertFile: todoServiceCertFile,
		todoServiceKeyFile:  todoServiceKeyFile,
		logger:              log.New(os.Stdout, "todo", log.LstdFlags|log.Lshortfile),
	}

	app.logger.Println("server starting...")
	serv := server.NewServer(app.Routes(), app.todoServiceAddr)
	err := serv.ListenAndServeTLS(app.todoServiceCertFile, app.todoServiceKeyFile)
	if err != nil {
		app.logger.Fatalf("server failed to start: %v", err)
	}
}
