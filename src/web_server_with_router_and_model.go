package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// use struct as Model
// !!! attributes need to be upper letters, otherwise json.NewEncoder will return empty
// Add `json:"xxx"` after since it’s not idiomatic JSON to have uppercased keys
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func StartWebServerWithRouterAndModel() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexNew)
	router.HandleFunc("/todos", TodoIndexNew)
	router.HandleFunc("/todos/{todoId}", TodoShowNew)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func IndexNew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndexNew(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Todo Index!")

	// send back JSON data
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	fmt.Printf("%v", todos)
	json.NewEncoder(w).Encode(todos)
}

func TodoShowNew(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show: ", todoId)
}
