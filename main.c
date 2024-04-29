package main
import (
	"fmt"
	"go-reloaded/customFuncStr"
	"go-reloaded/process"
	"log"
	"os"
	"strings"
)
// main is the entry point of the program
func main() {
	Args := os.Args
	// Check if the number of command-line arguments is correct
	if len(Args) < 3 {
		// Print usage instructions if arguments file name are missing
		fmt.Println("\033[31mFile name missing \nExample : go run . [input_file] [output file] \033[0m")
		log.Fatalln()
		return
	} else if len(Args) > 3 {
		// Print usage instructions if too many arguments are provided
		fmt.Println("\x1b[38;5;197mtoo many arguments are provided \nExample : go run . [input_file] [output file] \x1b[38;5;197m")
		log.Fatalln()
		return
	}
	input_file := Args[1]
	output_file := Args[2]
	if (!(strings.HasSuffix(input_file, ".txt"))) || (!(strings.HasSuffix(output_file, ".txt"))) {
		fmt.Println("\x1b[38;5;154m[!] Error: incorrect file extension (.txt).\x1b[38;5;154m")
		fmt.Println("Program exited.")
		log.Fatalln()
	}
	// Process the input file and handle any errors
	err := processFile(input_file, output_file)
	if err != nil {
		fmt.Println(err)
	}
}
// processFile processes the input file and writes the result to the output file
func processFile(inputFileName, outputFileName string) error {
	// Open the input file for reading
	file, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after processing
	// Read the content of the input file into a byte slice
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		return err
	}
	// Split the content into words
	splitTxt := customFuncStr.CustomSplit(string(data))
	// Initialize variables for processing quotes and skipping triggers
	nextSkip := 0
	quoteFlag := 0
	// Process triggers and quotes for each word in the split text
	for index, val := range splitTxt {
		process.ProcessTriggers(index, val, splitTxt)
		quoteFlag, nextSkip = process.ProcessQuote(index, splitTxt, quoteFlag, nextSkip)
	}
	// Process punctuation and reconstruct the text
	resultTxt := customFuncStr.CustomJoin(customFuncStr.CustomSplit(string(process.ProcessPunctuation(splitTxt))))
	// Write the processed text to the output file
	err = os.WriteFile(outputFileName, []byte(resultTxt), 0644)
	if err != nil {
		return err
	}
	return nil
}
