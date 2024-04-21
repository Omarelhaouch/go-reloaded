package reloaded

import "strings"

func Punctuations(s []string) []string {
	speChars := []string{".", ",", "!", "?", ":", ";"}
	for idx, sp := range s {
		for _, char := range speChars {
			// Check if the current element is the same as the special character
			if sp == char {
				// Concatenate the special character with the previous element
				s[idx-1] = s[idx-1] + char
				// Remove the current element from the slice
				s = RmIndex(s, idx)
				// Check if the current element starts with the special character
			} else if strings.HasPrefix(sp, char) {
				// Concatenate the special character with the previous element
				s[idx-1] = s[idx-1] + char
				// Remove the first character (special character) from the current element
				s[idx] = s[idx][1:]
			}
		}
	}
	return s
}
