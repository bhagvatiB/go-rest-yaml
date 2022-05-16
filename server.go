package main

import (
	"fmt"
	"net/http"
	"payloadrest/src/logger"
)

type Server struct {
}

/**
* Start and initialize server
**/
func main() {
	fmt.Println("Starting server -->")
	port := "8000"

	// initialze logger
	logger.InitLogger()
	log := logger.SugarLogger

	defer log.Sync()

	// initialize server application
	controller := InitializeServer()

	log.Infof("Invoking handlers")
	controller.HandleRequest()

	// setup Listening port
	log.Infof("Application listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Infof("Application started")
}
