package reloaded

import "strings"

func ConvertUpLowCap(s []string) []string {
	for i, tag := range s {
		if tag == "(low)" {
			// Convert the previous element to lowercase
			s[i-1] = strings.ToLower(s[i-1])
			// Remove the tag element from the slice
			s = RmIndex(s, i)
		} else if tag == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = RmIndex(s, i)
		} else if tag == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = RmIndex(s, i)
		}
	}
	return s
}
