# AoC 2025 - Go + Datastar

Advent of Code 2025 solutions with Go backend and Datastar frontend.

## Quickstart

```bash
make install
echo "SESSION=your_aoc_cookie" > .env
make dev
```

Open http://localhost:7331

## Commands

| Command | Description |
|---------|-------------|
| `make dev` | Hot reload on :7331 |
| `make serve` | Server on :8080 |
| `make test` | E2E tests |
| `make lint` | Format code |

## Stack

Go + Templ + Datastar + Playwright-go
