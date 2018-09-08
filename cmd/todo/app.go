package main

import (
	"log"

	"github.com/tupkung/go-todo-service/pkg/models"
)

//App type for using to collect configuration
type App struct {
	todoServiceAddr     string
	todoServiceCertFile string
	todoServiceKeyFile  string
	logger              *log.Logger
	dataBase            *models.Database
}
