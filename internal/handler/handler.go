package handler

import (
	"net/http"
	"strconv"

	"github.com/ltruchot/aoc2025-gods/templates"
	"github.com/ltruchot/aoc2025-gods/templates/day"
)

type Handler struct {
	version string
}

func New(version string) *Handler {
	return &Handler{version: version}
}

func (h *Handler) setCache(w http.ResponseWriter, r *http.Request) bool {
	etag := `"` + h.version + `"`
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", "no-cache")

	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		return true
	}
	return false
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if h.setCache(w, r) {
		return
	}
	templates.Index().Render(r.Context(), w)
}

func (h *Handler) Day(w http.ResponseWriter, r *http.Request) {
	dayStr := r.PathValue("day")
	dayNum, err := strconv.Atoi(dayStr)
	if err != nil || dayNum < 1 || dayNum > 25 {
		http.Error(w, "Invalid day", http.StatusBadRequest)
		return
	}
	if h.setCache(w, r) {
		return
	}
	switch dayNum {
	case 1:
		day.Day1().Render(r.Context(), w)
	case 2:
		day.Day2().Render(r.Context(), w)
	default:
		day.Day(dayNum).Render(r.Context(), w)
	}
}
