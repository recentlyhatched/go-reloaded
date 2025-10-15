package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	contentString := ""

	for _, b := range content {
		contentString += string(b)
	}

	// match "range a-z", optionally ", number)" then match ",", then match non-whitespace (^ = not)
	re := regexp.MustCompile(`\([a-z]+(,\s*\d+)?\)|[,]+|[^\s()]+`)
	tokens := re.FindAllString(contentString, -1)

	tokens = convertHexBin(tokens)
	tokens = applyCaseModifiers(tokens)
	tokens = fixApostrophes(tokens)
	tokens = fixIndefiniteArticles(tokens)
	tokens = formatPunctuation(tokens)

	output := strings.Join(tokens, " ")
	os.WriteFile(outputFile, []byte(output), 0644)
}
