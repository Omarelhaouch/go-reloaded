package reloaded

import "fmt"

func RmIndex(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		fmt.Println("index out of range")
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func RmIndexmore2(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		fmt.Println("index out of range")
		return s
	}
	return append(s[:index], s[index+2:]...)
}
