# Advent of Code 2025 - Go + Datastar

A minimal web scaffold for solving Advent of Code 2025 puzzles in JavaScript with Datastar, served by Go.

## Architecture

```
.
├── cmd/server/main.go       # HTTP server (routes + AoC proxy)
├── internal/handler/        # Request handlers
├── templates/               # Go Templ templates
│   ├── layout.templ         # Shared layout with Datastar CDN
│   ├── index.templ          # Home page
│   └── day/day.templ        # Day page template
├── static/                  # JavaScript puzzle solvers
│   └── day01.js             # Day 1 solution (edit this!)
├── tests/e2e/               # Playwright tests
└── Makefile
```

**Go server:** Serves pages + proxies AoC inputs (with caching)

**Datastar + JS:** Puzzle solving happens in the browser via `static/dayXX.js`

## Quick Start

```bash
# Install tools
make install

# Configure AoC session
cp .env.example .env
# Edit .env with your session cookie from adventofcode.com

# Run with hot reload
make dev
```

Open http://localhost:8080

## Routes

| Route | Description |
|-------|-------------|
| `/` | Home with day navigation |
| `/day/{n}` | Day n page (1-25) |
| `/adventofcode/2025/day/{n}/input` | Proxied AoC input |

## Solving a Day

Edit `static/dayXX.js`:

```javascript
function solvePart1(input) {
  // Your solution here
  return "answer";
}

function solvePart2(input) {
  // Your solution here
  return "answer";
}
```

Results display automatically via Datastar signals.

## Commands

| Command | Description |
|---------|-------------|
| `make dev` | Start with hot reload |
| `make serve` | Start once |
| `make build` | Build binary |
| `make lint` | Format code |
| `make test` | Run e2e tests |
| `make clean` | Remove artifacts |

## Stack

- **Go** - HTTP server, AoC input proxy
- **Go Templ** - HTML templates
- **Datastar** - Reactive UI bindings
- **Playwright** - E2E testing
