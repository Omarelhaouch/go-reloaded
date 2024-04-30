package reloaded

import (
	"strings"
)
// CustomSplit is a function that splits a string into fields based on certain delimiters.
// It returns a slice of strings, where each string is a field from the input string.
func CustomSplit(str string) []string {
	// fields will hold the fields from the input string.
	var fields []string
	// currentField is a strings.Builder that will be used to build each field.
	var currentField strings.Builder
	// Iterate over each character in the input string.
	for _, ch := range str {
		// If the character is a space, tab, period, exclamation mark, question mark, colon, or semicolon...
		if ch == ' ' || ch == '\t' || ch == '.' || ch == '!' || ch == '?' || ch == ':' || ch == ';' {
			// If currentField is not empty, append it to fields and reset it.
			if currentField.Len() != 0 {
				fields = append(fields, currentField.String())
				currentField.Reset()
			}
			// If the character is a period, exclamation mark, question mark, colon, or semicolon, append it to fields.
			if ch == '.' || ch == '!' || ch == '?' || ch == ':' || ch == ';' {
				fields = append(fields, string(ch))
			}
		} else if ch == '\n' { // If the character is a newline...
			// If currentField is not empty, append it to fields and reset it.
			if currentField.Len() != 0 {
				fields = append(fields, currentField.String())
				currentField.Reset()
			}
			// Append the newline to fields.
			fields = append(fields, "\n")
		} else { // If the character is not a delimiter...
			// Add it to currentField.
			currentField.WriteRune(ch)
		}
	}
	// If currentField is not empty after iterating over all characters, append it to fields.
	if currentField.Len() != 0 {
		fields = append(fields, currentField.String())
	}
	// Return the fields.
	return fields
}
// CustomJoin is a function that concatenates the fields in the input slice 'str' into a single string.
// It handles special cases where no space is added before the field.
func CustomJoin(str []string) string {
	// Initialize a strings.Builder to efficiently build the result string.
	var result strings.Builder
	// Iterate over the fields in the input slice.
	for i, field := range str {
		// Check if it is the first field, the previous field was a newline, the current field is a newline,
		// or the current field is a punctuation character.
		if i == 0 || field == "\n" || str[i-1] == "\n" || field == "." || field == "," || field == "!" || field == "?" || field == ":" || field == ";" || field == "'" {
			// If any of the conditions are met, write the field to the result without a leading space.
			result.WriteString(field)
		} else {
			// If none of the conditions are met, write the field to the result with a leading space.
			result.WriteString(" " + field)
		}
	}
	// Return the concatenated result string.
	return result.String()
}
