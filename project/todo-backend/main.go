package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Todo struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

var (
	todos  []Todo
	nextId int
)

func getTodos(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	nextId++
	t.Id = nextId
	todos = append(todos, t)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func main() {
	port := os.Getenv("TODO_BACKEND_PORT")
	fmt.Printf("Starting todo-backend on port %s\n", port)
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w)
		case http.MethodPost:
			createTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
		os.Exit(1)
	}
}
