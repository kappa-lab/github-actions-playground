package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, getWorld())
	})
	http.ListenAndServe(":8088", nil)
}

func getWorld() string {
	return "world is mine"
}
