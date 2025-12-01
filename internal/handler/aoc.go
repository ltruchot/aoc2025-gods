package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	inputCache = make(map[int]string)
	cacheMu    sync.RWMutex
)

func (h *Handler) AoCInput(w http.ResponseWriter, r *http.Request) {
	dayStr := r.PathValue("day")
	dayNum, err := strconv.Atoi(dayStr)
	if err != nil || dayNum < 1 || dayNum > 25 {
		http.Error(w, "Invalid day", http.StatusBadRequest)
		return
	}

	input, err := getInput(2025, dayNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(input))
}

func getInput(year, day int) (string, error) {
	cacheMu.RLock()
	if cached, ok := inputCache[day]; ok {
		cacheMu.RUnlock()
		return cached, nil
	}
	cacheMu.RUnlock()

	input, err := fetchFromAoC(year, day)
	if err != nil {
		return "", err
	}

	cacheMu.Lock()
	inputCache[day] = input
	cacheMu.Unlock()

	return input, nil
}

func fetchFromAoC(year, day int) (string, error) {
	session := os.Getenv("SESSION")
	if session == "" {
		return "", fmt.Errorf("SESSION env var not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	req.Header.Set("User-Agent", "github.com/ltruchot/aoc2025-gods")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AoC returned %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
