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

// sort summedItemsSlice
func sortSlice(summedItemsSlice []int) []int {
	slice_len := len(summedItemsSlice)
	for i := 0; i < slice_len; i++ {
		for j := i + 1; j < slice_len; j++ {
			if summedItemsSlice[i] > summedItemsSlice[j] {
				temp := summedItemsSlice[i]
				summedItemsSlice[i] = summedItemsSlice[j]
				summedItemsSlice[j] = temp
			}
		}
	}

	return summedItemsSlice
}

// Get top 3 biggest items
func getTopThree(summedItemsSlice []int, limit int) []int {
	slice_len := len(summedItemsSlice)
	get_min := slice_len - limit
	var topThree []int
	for i := get_min; i < slice_len; i++ {
		topThree = append(topThree, summedItemsSlice[i])
	}

	return topThree
}

// Sum slice items
func sumSlice(sliceArg []int) int {
	sliceItem := 0
	for _, i := range sliceArg {
		sliceItem += i
	}

	return sliceItem
}

func main() {
	inputFile := "../calories.txt"
	lineSlice := readFile(inputFile)

	summedItemsSlice := sumSliceItems(lineSlice)

	sortedSlice := sortSlice(summedItemsSlice)

	topThree := getTopThree(sortedSlice, 3)

	summedTopThree := sumSlice(topThree)
	fmt.Println(summedTopThree)
}
