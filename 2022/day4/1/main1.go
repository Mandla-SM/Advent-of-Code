package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errorExit(e error) {
	if e != nil {
		panic(e)
	}
}

// Create a slice of the lines in the file
func readFile(inputFile string) []string { // Returns a string slice
	fileData, err := os.Open(inputFile) // Open the file
	errorExit(err)                      // Exit the program if opening the file fails

	fileLines := bufio.NewScanner(fileData) // Get lines in the file
	fileLines.Split(bufio.ScanLines)        // Split the lines. In place mutate fileLines

	var lineSlice []string // An empty string slice
	for fileLines.Scan() { // Scan() makes fileLines iterable
		lineSlice = append(lineSlice, fileLines.Text()) // Append each line to lineSlice
	}
	fileData.Close() // Close the file

	return lineSlice // Return lineSlice
}

// Determine if sectiona contains sectionb or vice versa
func getBiggest(lineSlice []string) int {
	biggest := 0
	for _, line := range lineSlice {
		commaSplit := strings.Split(line, ",")                    // Split each line at the comma
		hyphenSplit1 := strings.Split(string(commaSplit[0]), "-") // Split half at "-", get index 0
		hyphenSplit2 := strings.Split(string(commaSplit[1]), "-") // Split half at "-", get index 1
		section1First, err := strconv.Atoi(hyphenSplit1[0])       // Get item by index and convert it to int.
		section1Last, err := strconv.Atoi(hyphenSplit1[1])        // Get item by index and convert it to int.
		section2First, err := strconv.Atoi(hyphenSplit2[0])       // Get item by index and convert it to int.
		section2Last, err := strconv.Atoi(hyphenSplit2[1])        // Get item by index and convert it to int.
		errorExit(err)
		if (section1First <= section2First && section1Last >= section2Last) || (section1First >= section2First && section1Last <= section2Last) {
			biggest += 1
		}
	}

	return biggest
}

func main() {
	inputFile := "../sections.txt"
	lineSlice := readFile(inputFile)

	biggest := getBiggest(lineSlice)
	fmt.Println(biggest)
}
