package main

import (
	loadbanner "asciiartoutput/banner"
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

	input = strings.ReplaceAll(input, "\\n", "\n")
	output := loadbanner.LoadBanner(input)
	fmt.Println(output)
}
