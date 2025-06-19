package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var count int

func readCountFromFile() int {
	data, err := os.ReadFile("files/pong.txt")
	if err != nil {
		return 0
	}
	parts := strings.Split(string(data), ": ")
	if len(parts) != 2 {
		return 0
	}
	parsed, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0
	}
	return parsed
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	count = readCountFromFile()
	initialData := fmt.Sprintf("Ping / Pongs: %d", count)
	err := os.WriteFile("files/pong.txt", []byte(initialData), 0644)
	if err != nil {
		panic(err)
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

	http.ListenAndServe(":"+port, nil)
}
