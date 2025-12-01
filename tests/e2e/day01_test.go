package e2e

import (
	"testing"

	"github.com/playwright-community/playwright-go"
)

func TestDay01Part1(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("http://localhost:8080/day/1"); err != nil {
		t.Fatalf("could not goto: %v", err)
	}

	output := page.Locator("#part1")
	if err := output.WaitFor(); err != nil {
		t.Fatalf("part1 output not visible: %v", err)
	}

	text, err := output.TextContent()
	if err != nil {
		t.Fatalf("could not get text content: %v", err)
	}

	if text == "" || text == "Loading..." {
		t.Error("Part 1 output is empty or still loading")
	}
}

func TestDay01Part2(t *testing.T) {
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		t.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		t.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("http://localhost:8080/day/1"); err != nil {
		t.Fatalf("could not goto: %v", err)
	}

	output := page.Locator("#part2")
	if err := output.WaitFor(); err != nil {
		t.Fatalf("part2 output not visible: %v", err)
	}

	text, err := output.TextContent()
	if err != nil {
		t.Fatalf("could not get text content: %v", err)
	}

	if text == "" || text == "Loading..." {
		t.Error("Part 2 output is empty or still loading")
	}
}
