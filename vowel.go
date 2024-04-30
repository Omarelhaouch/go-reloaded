package reloaded

import "strings"

func Vowel(s []string) []string {
    for i, val := range s {
        // Check if the word is "a" or "A".
        if val == "a" || val == "A" {
            // Define a list of vowels and consonant-like letters that could influence the article.
            vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
            // Check each character in the list of vowels.
            for _, char := range vowels {
                // If the next word starts with a vowel or consonant-like letter,
                if i+1 < len(s) && strings.HasPrefix(s[i+1], char) {
                    // Add "n" to the word.
                    s[i] = val + "n"
                }
            }
        }
    }
    return s
}
