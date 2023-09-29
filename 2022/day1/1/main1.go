package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func errorExit(e error) {
	if e != nil {
		panic(e)
	}
}

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

func sumSliceItems(lineSlice []string) []int { // Returns an int slice
	var summedItemsSlice []int
	sliceItem := 0                   // Initialising int variable to 0 so it can be used for additions
	for _, line := range lineSlice { // Loop over slice parameter
		if len(line) > 0 { // len() counts charactors in each line
			lineInt, err := strconv.Atoi(line) // converting each item to an int
			errorExit(err)

			sliceItem += lineInt                                   // adding the int to sliceItem
			summedItemsSlice = append(summedItemsSlice, sliceItem) // Appending the value to the summedItemsSlice slice
		} else { // When the line length is 0(empty)
			sliceItem = 0 // Reset sliceItem
		}
	}

	return summedItemsSlice // Return summedItemsSlice
}

func getBiggest(summedItemsSlice []int) int {
	biggestItem := summedItemsSlice[0] // Initialising int variable as first item of summedItemsSlice
	for _, item := range summedItemsSlice {
		if item > biggestItem { // Filter for the biggest item in summedItemsSlice
			biggestItem = item
		}
	}

	return biggestItem
}

func main() {
	inputFile := "../calories.txt"
	lineSlice := readFile(inputFile)

	summedItemsSlice := sumSliceItems(lineSlice)

	getBiggest := getBiggest(summedItemsSlice)
	fmt.Println(getBiggest)
}
