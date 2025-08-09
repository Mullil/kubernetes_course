package main

import (
	"fmt"
	"io"
	"log"
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
	if photoUrl == "" {
		return fmt.Errorf("PICSUM_URL env var not set")
	}

	resp, err := http.Get(photoUrl)
	if err != nil {
		log.Printf("Failed to download image: %v", err)
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(imagePath)
	if err != nil {
		log.Printf("Failed to create image file: %v", err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Printf("Failed to write image file: %v", err)
		return err
	}

	ts := time.Now().Format(time.RFC3339)
	err = os.WriteFile(timePath, []byte(ts), 0644)
	if err != nil {
		log.Printf("Failed to write timestamp file: %v", err)
		return err
	}

	log.Printf("Image updated successfully at %s", ts)
	return nil
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
