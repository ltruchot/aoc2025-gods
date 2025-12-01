package main

import (
	"log"
	"net/http"

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

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
