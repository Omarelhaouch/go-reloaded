package reloaded

import "strings"

func Vowel(s []string) []string {
	for i, val := range s {
		// Check if the word is "a", and if it's not the first or last word in the slice.
		if val == "a" && i > 0 && i+1 < len(s) {
			// Get the next word after "a".
			nextWord := s[i+1]
			// Define a list of vowels and consonant-like letters that could influence the article.
			vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
			// Check each character in the list of vowels.
			for _, char := range vowels {
				// If the next word starts with a vowel or consonant-like letter,
				// modify the current word "a" to "an".
				if strings.HasPrefix(nextWord, char) {
					s[i] = val + "n"
				}
			}
		}
	}
	return s
}
