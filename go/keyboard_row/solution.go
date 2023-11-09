package keyboardrow

import "strings"

func canType(word string) bool {
	var (
		rows = []string{
			"qwertyuiop",
			"asdfghjkl",
			"zxcvbnm",
		}
	)

	rowIndex := 0

	for i, c := range word {
		if c >= 'A' && c <= 'Z' {
			c = 'a' + c - 'A'
		}

		if i == 0 {
			for j, row := range rows {
				if strings.ContainsRune(row, c) {
					rowIndex = j
				}
			}
		} else {
			if !strings.ContainsRune(rows[rowIndex], c) {
				return false
			}
		}

	}

	return true
}

func findWords(words []string) []string {
	result := make([]string, 0)

	for _, word := range words {
		if canType(word) {
			result = append(result, word)
		}
	}

	return result
}
