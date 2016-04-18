package main

import (
	"log"
	"net/http"

	"./service"
)

func main() {
	// StartBasicWebServer()
	// StartWebServerWithRouter()
	// StartWebServerWithRouterAndModel()

	// /service has well-structured restful service
	/*
		    ../main.go
				handlers.go
				routes.go
				todo.go
	*/
	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
