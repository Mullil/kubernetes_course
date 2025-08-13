package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ping_count (
		id SERIAL PRIMARY KEY,
		value INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	_, err = db.Exec(`INSERT INTO ping_count (value) 
		SELECT 0 WHERE NOT EXISTS (SELECT 1 FROM ping_count)`)
	if err != nil {
		log.Fatalf("Failed to insert initial count: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := db.Exec(`UPDATE ping_count SET value = value + 1 WHERE id = 1`)
		if err != nil {
			http.Error(w, "DB update failed", http.StatusInternalServerError)
			return
		}

		var count int
		err = db.QueryRow(`SELECT value FROM ping_count WHERE id = 1`).Scan(&count)
		if err != nil {
			http.Error(w, "DB read failed", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Pong %d\n", count)
		data := fmt.Sprintf("Ping / Pongs: %d", count)
		err = os.WriteFile("files/pong.txt", []byte(data), 0644)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/pings", func(w http.ResponseWriter, r *http.Request) {
		var count int
		err := db.QueryRow(`SELECT value FROM ping_count WHERE id = 1`).Scan(&count)
		if err != nil {
			http.Error(w, "DB read failed", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%d\n", count)
	})

	http.ListenAndServe(":"+port, nil)
}
