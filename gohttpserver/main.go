package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "foo")
	})

	http.ListenAndServe("localhost:8080", nil)
	//http.ListenAndServe(":8080", nil)
}

/*
http.ResponseWriter
http.Request
ServeHTTP
ListenAndServe
HandleFunc
Handle
*/
