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
		count++
		fmt.Fprintf(w, "Pong %d\n", count)
		data := fmt.Sprintf("Ping / Pongs: %d", count)
		err := os.WriteFile("files/pong.txt", []byte(data), 0644)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/pings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d\n", count)
	})

	http.ListenAndServe(":"+port, nil)
}
