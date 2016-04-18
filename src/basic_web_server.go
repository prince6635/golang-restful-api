package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func StartBasicWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// print 'Hello, "/foo/zi/1"' for "http://localhost:8080/foo/zi/1"
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
