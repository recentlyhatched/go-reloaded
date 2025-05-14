package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func CapitalizeWord(s string) string {
	result := ""

	stringParts := regexp.MustCompile(`\(cap[^)]+\)`).Split(s, -1)

	// stringParts := strings.Split(s, " (cap)")
	fmt.Println(stringParts)
	for _, stringPart := range stringParts {
		for i, ch := range stringPart {
			if i == 0 {
				result += strings.ToUpper(string(ch))
			} else {
				result += string(ch)
			}
		}
	}
	return result

	// s = "modified again"
	// return s
}

func main() {
	if len(os.Args) == 3 {
		fileName := os.Args[1]
		rawInputString := ReadFile(fileName)

		modifiedStr := CapitalizeWord(rawInputString)

		resultFile := os.Args[2]
		WriteFile(resultFile, modifiedStr)
	}
}

					func ReadFile(fileName string) string {
	result := ""
	file, err := os.Open("sample.txt")
	// check for errors opening the file
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	// creates a Scanner object that efficiently reads data from a file in a buffered manner by spliting input into lines
	scanner := bufio.NewScanner(file)

	// loop through each line in the file (while-style for loop)
	for scanner.Scan() {
		line := scanner.Text() // returns string
		// fmt.Println(line)
		result += line
	}

	// check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file %s", err)
	}

	return result
}

func WriteFile(resultFile, modifiedString string) {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer file.Close()

	_, err = file.WriteString(modifiedString)
	if err != nil {
		log.Fatalf("failed to write file: %s", err)
	}
}
