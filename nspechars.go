package reloaded

import "strings"

// MultSpeChars is a function that combines consecutive special characters with adjacent words in a slice of strings.
func MultSpeChars(s []string) []string {
	// Define a list of special characters.
	spChars := []string{",", "!", "?", ":", ";", "."}

	// Iterate over each word in the slice along with its index.
	for i, val := range s {
		// Iterate over each special character.
		for _, v := range spChars {
			// Check if the word is longer than 1 character and both starts and ends with the current special character.
			if len(val) > 1 && strings.HasPrefix(val, v) && strings.HasSuffix(val, v) {
				// Combine the current word with the previous word and remove the current word from the slice.
				s[i-1] = s[i-1] + val
				s = RmIndex(s, i)
			}
		}
	}
	// Process the resulting slice to ensure proper punctuation usage.
	return Punctuations(s)
}
