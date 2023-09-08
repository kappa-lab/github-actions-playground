package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, getWorld())
	})

	http.HandleFunc("/404", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(404)
	})

	http.ListenAndServe(":8088", nil)
}

func getWorld() string {
	return "world is mine"
}
