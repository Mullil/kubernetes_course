package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	imagePath = os.Getenv("IMAGE_PATH")
	timePath  = os.Getenv("TIME_PATH")
)

func isValid() bool {
	ts, err := os.ReadFile(timePath)
	if err != nil {
		return false
	}
	t, err := time.Parse(time.RFC3339, string(ts))
	if err != nil {
		return false
	}
	return time.Since(t) < 10*time.Minute
}

func updateImage() error {
	photoUrl := os.Getenv("PICSUM_URL")
	resp, err := http.Get(photoUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	ts := time.Now().Format(time.RFC3339)
	return os.WriteFile(timePath, []byte(ts), 0644)
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	if !isValid() {
		err := updateImage()
		if err != nil {
			http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, imagePath)
}

func main() {
	port := os.Getenv("SERVER_PORT")

	http.HandleFunc("/files/image", serveImage)

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./frontend/build")).ServeHTTP(w, r)
	})

	fmt.Printf("Server started in port %s\n", port)

	http.ListenAndServe(":"+port, nil)
}
