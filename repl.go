package main

import (
	"strings"
)

func cleanInput(text string) []string {
	// words := make([]string, 0)
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
