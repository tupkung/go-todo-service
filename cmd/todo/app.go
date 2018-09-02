package main

import "log"

//App type for using to collect configuration
type App struct {
	todoServiceAddr     string
	todoServiceCertFile string
	todoServiceKeyFile  string
	logger              *log.Logger
}
