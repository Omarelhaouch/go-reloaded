package main

import (
    // "fmt"
    "os"
    "reloaded"
    "strings"
)

func main() {
    Args := os.Args[1:] // Get command-line arguments excluding the program name

    // Open the file specified in the first command-line argument
    file, err := os.Open(Args[0])
    if err != nil {
        panic(err)
    }
    defer file.Close() // Ensure the file is closed after processing

    // Read the content of the file
    data, err := os.ReadFile(Args[0])
    if err != nil {
        panic(err)
    }

    // Split the content of the file into a slice of words
    strSlice := strings.Split(string(data), " ")

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
    // fmt.Println("\n" + sentence)

    // Write the modified string to the output file
    err = os.WriteFile(Args[1], []byte(sentence), 0644)
    if err != nil {
        panic(err)
    }
}
