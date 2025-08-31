package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type newTodo struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type Todo struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func getTodos(db *sql.DB, w http.ResponseWriter) {
	rows, err := db.Query("SELECT id, content, done FROM todos")
	if err != nil {
		http.Error(w, "Failed to query todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.Id, &t.Content, &t.Done); err != nil {
			http.Error(w, "Failed to scan todo", http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var t newTodo
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &t); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = db.QueryRow("INSERT INTO todos (content) VALUES ($1) RETURNING id", t.Content).Scan(&t.Id)
	if err != nil {
		http.Error(w, "Failed to insert todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func markDone(db *sql.DB, w http.ResponseWriter, id int) {
	var t Todo
	err := db.QueryRow(
		"UPDATE todos SET done = TRUE WHERE id = $1 RETURNING id, content, done", id).Scan(
		&t.Id, &t.Content, &t.Done)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func main() {
	port := os.Getenv("TODO_BACKEND_PORT")
	fmt.Printf("Starting todo-backend on port %s\n", port)
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("Failed to open DB: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		done BOOLEAN DEFAULT FALSE
	)`)
	if err != nil {
		log.Printf("Failed to create table: %v", err)
	}

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(db, w)
		case http.MethodPost:
			createTodo(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodPut:
			markDone(db, w, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			http.Error(w, fmt.Sprintf("db is not ready: %v", err), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Server failed: %v\n", err)
		os.Exit(1)
	}
}
