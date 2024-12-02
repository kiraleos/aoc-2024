package util

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func splitInputToLines(input string) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	return lines, nil
}

func fetchInput(day int) (string, error) {
	cachedInput, err := GetCachedInput(day)
	if err != nil {
		return "", fmt.Errorf("error checking cache: %w", err)
	}
	if cachedInput != "" {
		return cachedInput, nil
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return "", fmt.Errorf("AOC_SESSION environment variable is not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("failed to fetch input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch input: status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	input := string(body)

	if err := SaveInputToCache(day, input); err != nil {
		return "", fmt.Errorf("failed to save input to cache: %w", err)
	}

	return input, nil
}

func FetchInputLines(day int) []string {
	input, err := fetchInput(day)
	if err != nil {
		panic(err)
	}

	lines, err := splitInputToLines(input)
	if err != nil {
		panic(err)
	}
	return lines
}