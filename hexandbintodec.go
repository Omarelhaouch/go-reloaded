package reloaded

import (
	"fmt"
	"strconv"
)

func HexAndBinToDec(s []string) []string {
	for i, tag := range s {
		if tag == "(hex)" {
			// Convert the previous element from hexadecimal to decimal
			dec, err := strconv.ParseInt(s[i-1], 16, 64)
			s[i-1] = strconv.Itoa(int(dec))
			// Remove the tag element from the slice
			s = RmIndex(s, i)
			// Print an error message if conversion fails
			if err != nil {
				fmt.Printf("ERROR")
			}
		} else if tag == "(bin)" {
			dec, err := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = strconv.Itoa(int(dec))
			s = RmIndex(s, i)
			if err != nil {
				fmt.Printf("ERROR")
			}
		}
	}
	return s
}
