package main

import (
	ascii "ascii/banner"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No enough arguments: go run <main.go> [string]")
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		fmt.Println("Too many agurment: usage go run <main.go> [string] <bannerfile> ")
		os.Exit(1)
	}

	input := os.Args[1]

	// input = strings.ReplaceAll(input, "\\n", "\n")
	if input == "\n" {
		fmt.Println()
		return
	} else if input == "" {
		return
	}
	// Split the input into lines based on newline characters.

	Input := strings.Split(input, "\n")

	spaceCount := 0
	// Iterate over each line of the input.
	for _, word := range Input {
		if word == "" {
			spaceCount++
			if spaceCount < len(Input) {
				fmt.Println()
			}
		} else {
			// Print the banner for non-empty strings.
			ascii.PrinBanner(word)
		}
	}
}
