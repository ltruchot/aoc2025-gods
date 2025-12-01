package handler

import (
	"net/http"
	"strconv"

	"github.com/ltruchot/aoc2025-gods/templates"
	"github.com/ltruchot/aoc2025-gods/templates/day"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	templates.Index().Render(r.Context(), w)
}

func (h *Handler) Day(w http.ResponseWriter, r *http.Request) {
	dayStr := r.PathValue("day")
	dayNum, err := strconv.Atoi(dayStr)
	if err != nil || dayNum < 1 || dayNum > 25 {
		http.Error(w, "Invalid day", http.StatusBadRequest)
		return
	}
	switch dayNum {
	case 1:
		day.Day1().Render(r.Context(), w)
	default:
		day.Day(dayNum).Render(r.Context(), w)
	}
}
