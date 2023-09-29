package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func errorExit(e error) {
	if e != nil {
		panic(e)
	}
}

// Create a slice of the lines in the file.
func readFile(inputFile string) []string { // Returns a string slice.
	fileData, err := os.Open(inputFile) // Open the file.
	errorExit(err)                      // Exit the program if opening the file fails.

	fileLines := bufio.NewScanner(fileData) // Get lines in the file.
	fileLines.Split(bufio.ScanLines)        // Split the lines. In place mutate fileLines.

	var lineSlice []string // An empty string slice.
	for fileLines.Scan() { // Scan() makes fileLines iterable.
		lineSlice = append(lineSlice, fileLines.Text()) // Append each line to lineSlice.
	}
	fileData.Close() // Close the file.

	return lineSlice // Return lineSlice.
}

// Remove duplicate alphabets in each slice.
func removeDuplicates(lines []string) []string {
	var deDupedSlice []string
	for _, line := range lines {
		var uniqueSlice []string
		for _, alphabet := range line {
			if !slices.Contains(uniqueSlice, string(alphabet)) { // 2. Ignore duplicates.
				uniqueSlice = append(uniqueSlice, string(alphabet)) // 1. Add unique.
			}
		}

		var sliceAsString string
		for _, alphabet := range uniqueSlice {
			sliceAsString += alphabet // 3. Make uniqueSlice1 alphabets into a single string.
		}

		deDupedSlice = append(deDupedSlice, sliceAsString)
	}

	return deDupedSlice
}

// Group the split items into a slice, add the slice to a slice.
func groupSlices(lineSlice []string) [][]string {
	var threes [][]string
	var three []string
	for _, line := range lineSlice {
		three = append(three, line) // Create a group of three lines.
		if len(three) == 3 {
			threes = append(threes, three) // Append the group of three to threes.
			three = nil                    // Reset the three slice.
		}
	}

	return threes
}

// Parse the whole dataset and make a slice of unique alphabets.
func getUniqueAlphabets(deDupedSlice []string) []rune {
	var alphabets []rune
	if len(alphabets) <= 36 { // There are 36 unique alphabets in the dataset.
		for _, line := range deDupedSlice {
			for _, alphabet := range line {
				if !slices.Contains(alphabets, alphabet) { // 2. If the alphabet isn't in alphabets yet.
					alphabets = append(alphabets, alphabet) // 1. Add the alphabet.
				}
			}
		}
	}

	return alphabets
}

// Find the repeating alphabet in each group of three.
func findRepeats(groupedSlices [][]string, alphabets []rune) []string {
	var commonAlphabets []string
	for _, slice := range groupedSlices {
		var commonInThree string
		line1 := slice[0]
		line2 := slice[1]
		line3 := slice[2]
		for _, alphabet := range alphabets {
			toString := string(alphabet) // Convert the alphabet from a rune to a string.
			if strings.Contains(line1, toString) && strings.Contains(line2, toString) && strings.Contains(line3, toString) {
				commonInThree += toString // Add the alphabet to commonInThree.
			}
		}
		commonAlphabets = append(commonAlphabets, commonInThree)
	}

	return commonAlphabets
}

// Loop over the alphabet slice and calculate each alphabet's ascii value.
func getAsciiCode(alphabetSlice []string) []int {
	var alphabetIndexes []int
	for _, alphabet := range alphabetSlice {
		lowercase := alphabet[0] - 96 // Convert lowercase alphabets.
		if lowercase <= 26 {
			alphabetIndexes = append(alphabetIndexes, int(lowercase))
		} else {
			uppercase := alphabet[0] - 38 // Convert uppercase alphabets.
			alphabetIndexes = append(alphabetIndexes, int(uppercase))
		}
	}

	return alphabetIndexes
}

// Sum the values of the slice.
func sumSlice(sliceArg []int) int {
	sliceItem := 0
	for _, i := range sliceArg {
		sliceItem += i
	}

	return sliceItem
}

func main() {
	inputFile := "../items_demo.txt"
	lineSlice := readFile(inputFile)

	removedDuplicates := removeDuplicates(lineSlice)

	groupedSlices := groupSlices(removedDuplicates)

	alphabets := getUniqueAlphabets(removedDuplicates)

	commonAlphabets := findRepeats(groupedSlices, alphabets)

	alphabetIndexSlice := getAsciiCode(commonAlphabets)

	summedItems := sumSlice(alphabetIndexSlice)
	fmt.Println(summedItems)
}
