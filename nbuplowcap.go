package reloaded

import (
	"strconv"
	"strings"
)

func CaseN(s []string) []string {
	for idx, word := range s {
		// Check if the word is "(up," and the index is within bounds.
		if word == "(up," && idx < len(s)-1 {
			// Extract the value after "(up," and convert it to an integer.
			endWord := strings.TrimSuffix(s[idx+1], ")")
			nb, _ := strconv.Atoi(endWord)
			// Iterate for the specified number of times.
			for i := 1; i <= nb; i++ {
				// Calculate the index of the word to be modified.
				calcu := (idx - i)
				// Check if the calculated index is within bounds.
				if calcu >= 0 && calcu < len(s) {
					// Convert the word to uppercase.
					s[calcu] = strings.ToUpper(s[calcu])
				}
			}
			// Remove the specified index and the next one from the slice.
			s = RmIndexmore2(s, idx)

		} else if word == "(low," {
			endWord := strings.TrimSuffix(s[idx+1], ")")
			nb, _ := strconv.Atoi(endWord)
			for i := 1; i <= nb; i++ {
				calcu := (idx - i)
				if calcu >= 0 && calcu < len(s) {
					s[calcu] = strings.ToLower(s[calcu])
				}
			}
			s = RmIndexmore2(s, idx)
		} else if word == "(cap," {
			endWord := strings.TrimSuffix(s[idx+1], ")")
			nb, _ := strconv.Atoi(endWord)
			for i := 1; i <= nb; i++ {
				calcu := (idx - i)
				if calcu >= 0 && calcu < len(s) {
					s[calcu] = strings.Title(s[calcu])
				}
			}
			s = RmIndexmore2(s, idx)
		}
	}
	return s
}
