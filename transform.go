package main

import (
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func convertHexBin(tokens []string) []string {
	var result []string
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(hex)" && i > 0 {
			// convert from string representations of number to hexadecimal (base 16) as a 64 bit integer
			val, err := strconv.ParseInt(tokens[i-1], 16, 64)
			if err == nil {
				// FormatInt() turns value into string representation to specified base (which is base 10/decimal number here)
				// add to last index of results slice
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}
		} else if tokens[i] == "(bin)" && i > 0 {
			val, err := strconv.ParseInt(tokens[i-1], 2, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(val, 10)
			}
		} else {
			result = append(result, tokens[i])
		}
	}
	return result
}

func applyCaseModifiers(tokens []string) []string {
	var result []string
	i := 0

	caser := cases.Title(language.English)

	for i < len(tokens) {
		token := tokens[i]

		// Check if it's a command like (up), (low, 2), etc.
		if strings.HasPrefix(token, "(up") || strings.HasPrefix(token, "(low") || strings.HasPrefix(token, "(cap") {
			mod := strings.Trim(token, "()")
			parts := strings.Split(mod, ",")
			command := parts[0]
			count := 1
			if len(parts) > 1 {
				c, err := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err == nil {
					count = c
				}
			}
			for j := len(result) - 1; j >= 0 && count > 0; j-- {
				switch command {
				case "up":
					result[j] = strings.ToUpper(result[j])
				case "low":
					result[j] = strings.ToLower(result[j])
				case "cap":
					// result[j] = strings.Title(strings.ToLower(result[j]))
					result[j] = caser.String(strings.ToLower(result[j]))
				}
				count--
			}
		} else {
			result = append(result, token)
		}
		i++
	}
	return result
}

func fixApostrophes(tokens []string) []string {
	var result []string
	inQuote := false
	var quoteBuffer []string

	// read from 1. to 4.
	for _, token := range tokens {
		if token == "'" {
			if inQuote {
				// 4. End quote
				// add the single quote marks on either side of quoteBuffer
				result = append(result, "'"+strings.Join(quoteBuffer, " ")+"'")
				// reset inQuote and quoteBuffer
				inQuote = false
				quoteBuffer = []string{}
			} else {
				// 2. if token is a single quote mark, and inQuote is false
				// it means it is a starting quote
				// notice we don't add the single quote mark to []result yet
				inQuote = true
			}
		} else if inQuote {
			// 3. if token is within a quote, add token to []quoteBuffer
			quoteBuffer = append(quoteBuffer, token)
		} else {
			// 1. if token is not within a quote or not a single quote
			result = append(result, token)
		}
	}
	// If unmatched quote remains (optional according to program instrcutions)
	if len(quoteBuffer) > 0 {
		// quoteBuffer... has a variadic argument, which treats each element in qouteBuffer as a seperate argument
		// this avoids the need for a loop
		result = append(result, quoteBuffer...)
	}
	return result
}

func fixIndefiniteArticles(tokens []string) []string {
	vowels := "aeiouhAEIOUH"
	var result []string
	for i := 0; i < len(tokens); i++ {
		// avoid out or bounds errors is "a" is the last token
		// strings.ContainsRune() returns boolean- find second rune parameter in first string parameter
		if tokens[i] == "a" && i+1 < len(tokens) && strings.ContainsRune(vowels, rune(tokens[i+1][0])) {
			result = append(result, "an")
		} else {
			result = append(result, tokens[i])
		}
	}
	return result
}

func formatPunctuation(tokens []string) []string {
	var result []string
	i := 0

	for i < len(tokens) {
		current := tokens[i]

		if isPunctuation(current) {
			// Attach to previous word
			if len(result) > 0 {
				result[len(result)-1] += current
			} else {
				result = append(result, current)
			}
		} else if isEllipsisOrCombo(current) {
			// Example: "..." or "!?" — treat as group
			if len(result) > 0 {
				result[len(result)-1] += current
			} else {
				result = append(result, current)
			}
		} else {
			result = append(result, current)
		}
		i++
	}

	// Re-insert proper spacing
	// After punctuation, if next is not another punctuation, insert space
	final := []string{}
	for i := 0; i < len(result); i++ {
		final = append(final, result[i])
	}
	return final
}

func isPunctuation(s string) bool {
	return s == "." || s == "," || s == "!" || s == "?" || s == ":" || s == ";"
}

func isEllipsisOrCombo(s string) bool {
	return s == "..." || s == "!?" || s == "?!"
}
