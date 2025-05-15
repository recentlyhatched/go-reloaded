package main

import (
	"strconv"
	"strings"
)

func convertHexBin(tokens []string) []string {
	var result []string
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(hex)" && i > 0 {
			val, err := strconv.ParseInt(tokens[i-1], 16, 64)
			if err == nil {
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
					result[j] = strings.Title(strings.ToLower(result[j]))
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

	for _, token := range tokens {
		if token == "'" {
			if inQuote {
				// End quote
				result = append(result, "'"+strings.Join(quoteBuffer, " ")+"'")
				inQuote = false
				quoteBuffer = []string{}
			} else {
				inQuote = true
			}
		} else if inQuote {
			quoteBuffer = append(quoteBuffer, token)
		} else {
			result = append(result, token)
		}
	}
	// If unmatched quote remains
	if len(quoteBuffer) > 0 {
		result = append(result, quoteBuffer...)
	}
	return result
}

func fixIndefiniteArticles(tokens []string) []string {
	vowels := "aeiouhAEIOUH"
	var result []string
	for i := 0; i < len(tokens); i++ {
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
