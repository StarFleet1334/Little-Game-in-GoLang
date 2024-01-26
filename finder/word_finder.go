package main

import (
	"fmt"
	"os"
	"strings"
)

const text = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(text)
	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			// Here we also take into consideration case-sensitive searching
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}
