package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	// Add a router by the mux router from the Gorilla Web Toolkit.
	"github.com/gorilla/mux"
)

func StartWebServerWithRouter() {
	router := mux.NewRouter().StrictSlash(true)
	/*
		    If only have the following line:
			  Before we could ask for http://localhost:8080/foo and that worked.
			  That will no longer work now as there is no route defined.
			  Only http://localhost:8080 will be a valid response.
	*/
	router.HandleFunc("/", Index)

	// add more basic routes
	// This is the Todo Index route: http://localhost:8080/todos
	router.HandleFunc("/todos", TodoIndex)
	// This is the Todo Show route: http://localhost:8080/todos/{todoId}
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show: ", todoId)
}
