package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		http.FileServer(http.Dir("./frontend/build")).ServeHTTP(w, r)
	})

	fmt.Printf("Server started in port %s\n", port)

	http.ListenAndServe(":"+port, nil)
}
