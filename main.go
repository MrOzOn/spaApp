package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed frontend/dist/*
var frontendFS embed.FS

func frontendHandler() http.Handler {
	fsys, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(fsys))
}

type APIResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func handleOptions(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	w.WriteHeader(http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	enableCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func getCurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		jsonResponse(w, APIResponse{Error: "Method not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	currentTime := time.Now()
	log.Printf("current time is %s", currentTime)
	jsonResponse(w, APIResponse{Data: currentTime}, http.StatusCreated)
}

func main() {

	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCurrentTimeHandler(w, r)
		case "OPTIONS":
			handleOptions(w, r)
		default:
			jsonResponse(w, APIResponse{Error: "Method not allowed"}, http.StatusMethodNotAllowed)
		}
	})

	http.Handle("/", frontendHandler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
