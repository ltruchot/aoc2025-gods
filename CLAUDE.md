# Advent of Code 2025 - GO + Datastar

Claude is in charge to provide a minimalistic (but UX friendly) scaffold for the Advent of Code 2025.

You NEVER provide help, hints, or solutions to the puzzles. You are only a scaffold developer.

## Architecture

**Go server role (backend):**
- Serve static HTML pages via Go Templ templates
- Proxy AoC inputs with caching (fetch once, serve from cache)
- Routes: `/day/1`, `/day/2`, etc. for each day's page (12 days)
- Route: `/adventofcode/2025/day/{n}/input` to proxy and cache puzzle inputs

**Datastar role (frontend):**
- All puzzle solving logic is written in JavaScript within each day's template
- Datastar provides reactive data binding to display results
- The input is fetched client-side from the Go proxy endpoint
- Solutions are computed in the browser, not on the server

## Stack

- Go - Simple HTTP server + AoC input proxy with cache
- Go Templ - HTML templating (compile-time)
- Datastar - Reactive UI via CDN, puzzle solving in JS
- Each day is a separate page based on the same layout
- Makefile to serve the pages locally
- Hot reload: `templ generate --watch --proxy` (browser auto-refresh)
- playwright-go for e2e tests (2 tests per day: part1 + part2)
- Linter/formatter: go fmt + templ fmt

## CSS Strategy

- `static/common.css` - Shared styles (body, header, nav, buttons, inputs, outputs)
- `static/dayXX.css` - Day-specific styles (loaded only on that day's page)
- View Transitions API for smooth navigation between pages
- Navigation highlights active page in violet

## No Node.js

This project is 100% Go. No npm, no package.json, no node_modules.
- E2E tests: playwright-go (not Playwright with npm)
- Formatting: go fmt + templ fmt (not Prettier)
- Only JS allowed: browser-side scripts for puzzle solving (vanilla JS, no build step)

## Page Structure

Each day page contains:
- `<output id="part1" data-text="$part1Output">` - displays Part 1 result
- `<output id="part2" data-text="$part2Output">` - displays Part 2 result
- JavaScript that fetches input and computes both solutions
- Datastar signals to bind results to outputs

## Environment

`.env` file (never committed):
```
SESSION=your_aoc_session_cookie
```

Frontend fetches input from: `http://localhost:8080/adventofcode/2025/day/1/input`

## Makefile Tasks

- `make install` - Install templ, air, and dependencies
- `make install-browsers` - Install Playwright browsers
- `make generate` - Generate templ files
- `make build` - Build the server binary
- `make serve` - Run server (no hot reload)
- `make dev` - Run with hot reload (templ watch + proxy on :7331)
- `make lint` - Format Go and templ files
- `make test` - Run e2e tests
- `make clean` - Remove build artifacts

## Claude Good Practices

- Go server is minimal: just routing + proxy, no solving logic
- All puzzle logic lives in JS within templ files
- Code is idiomatic for each language (Go, JS, HTML)
- Latest stable versions of all tools

Datastar CDN:
```html
<script type="module" src="https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar.js"></script>
```

Datastar docs: https://data-star.dev/
