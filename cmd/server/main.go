package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ltruchot/aoc2025-gods/internal/handler"
)

// Version is set at build time via -ldflags
var Version = "dev"

// cacheStatic wraps a handler to add cache headers for static assets
func cacheStatic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	h := handler.New(Version)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", h.Index)
	mux.HandleFunc("GET /day/{day}", h.Day)
	mux.HandleFunc("GET /adventofcode/2025/day/{day}/input", h.AoCInput)
	mux.Handle("GET /static/", cacheStatic(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
