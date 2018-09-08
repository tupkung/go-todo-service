package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/tupkung/go-todo-service/pkg/models"
	"github.com/tupkung/go-todo-service/server"

	_ "github.com/go-sql-driver/mysql"
)

var (
	todoServiceAddr     = os.Getenv("TODO_SERVICE_ADDR")
	todoServiceCertFile = os.Getenv("TODO_SERVICE_CERT_FILE")
	todoServiceKeyFile  = os.Getenv("TODO_SERVICE_KEY_FILE")
	todoMySQLDSN        = os.Getenv("TODO_MYSQL_DSN")
)

func main() {
	db := connect(todoMySQLDSN)
	defer db.Close()

	app := &App{
		todoServiceAddr:     todoServiceAddr,
		todoServiceCertFile: todoServiceCertFile,
		todoServiceKeyFile:  todoServiceKeyFile,
		logger:              log.New(os.Stdout, "todo", log.LstdFlags|log.Lshortfile),
		dataBase:            &models.Database{db},
	}

	err := app.dataBase.Migrate()
	if err != nil {
		app.logger.Fatalf("db migration failed: %v", err)
	}

	app.logger.Println("server starting...")
	serv := server.NewServer(app.Routes(), app.todoServiceAddr)
	err = serv.ListenAndServeTLS(app.todoServiceCertFile, app.todoServiceKeyFile)
	if err != nil {
		app.logger.Fatalf("server failed to start: %v", err)
	}
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
