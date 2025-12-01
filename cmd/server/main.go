package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ltruchot/aoc2025-gods/internal/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	h := handler.New()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", h.Index)
	mux.HandleFunc("GET /day/{day}", h.Day)
	mux.HandleFunc("GET /adventofcode/2025/day/{day}/input", h.AoCInput)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
