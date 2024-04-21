package main

import (
	"os"
	"reloaded"
	"strings"
)

func main() {
	args := os.Args[1:] // Get command-line arguments excluding the program name

	// Read the content of the file specified in the first command-line argument
	text, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	// Split the content of the file into a slice of words
	strSlice := strings.Split(string(text), " ")

	// Apply a series of transformations to the word slice
	strSlice = reloaded.CaseN(strSlice)
	strSlice = reloaded.ConvertUpLowCap(strSlice)
	strSlice = reloaded.Punctuations(strSlice)
	strSlice = reloaded.MultSpeChars(strSlice)
	strSlice = reloaded.HexAndBinToDec(strSlice)
	strSlice = reloaded.Vowel(strSlice)
	strSlice = reloaded.Apostrophe(strSlice)

	// Join the modified word slice back into a string
	sentence := strings.Join(strSlice, " ") + "\n"

	// Create a new file named "result.txt" and write the modified string into it
	file, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the modified string to the file
	_, err = file.WriteString(sentence)
	if err != nil {
		panic(err)
	}
}
