package main

import (
	"fmt"
	"net/http"
	"os"
)

var count int

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong %d\n", count)
		count++
	})

	http.ListenAndServe(":"+port, nil)
}
