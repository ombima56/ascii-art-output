package loadbanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(input string) map[rune]string {
	var height int
	banner := make(map[rune]string)
	currentChar := rune(32)
	charLine := []string{}
	filePath := "./bannerfile/" + input + ".txt"

	// File openning

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error openning file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// scanner.Scan()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if height == 8 {
			banner[currentChar] = strings.Join(charLine, "\n")
			currentChar++
			height = 0
			charLine = []string{}

		} else {
			charLine = append(charLine, line)
		}
	}
	if height > 0 {
		banner[currentChar] = strings.Join(charLine, "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scaning file", err)
		os.Exit(1)
	}
	return banner
}
