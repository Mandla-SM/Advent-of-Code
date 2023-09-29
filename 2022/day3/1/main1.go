package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

// split slice item into two halves
func splitInTwo(item string) (string, string) {
	itemLength := len(item)
	item1 := item[:itemLength/2]
	item2 := item[itemLength/2:]

	return item1, item2
}

// Group the split items into a slice, add the slice to a slice
func groupSlices(lineSlice []string) [][]string {
	splitItems := [][]string{}
	for _, item := range lineSlice {
		itemSlice := []string{}
		item1, item2 := splitInTwo(item)
		itemSlice = append(itemSlice, item1, item2)
		splitItems = append(splitItems, itemSlice)
	}

	return splitItems
}

// Remove duplicate letters in each half.
func cleanDuplicates(twoHalves [][]string) [][]string {
	var splitUniqueItems [][]string

	for _, items := range twoHalves {
		var newTwoHalves []string

		slice1 := items[0]
		var uniqueSlice1 []string
		for _, item1 := range slice1 {
			if !slices.Contains(uniqueSlice1, string(item1)) { // 2. ignore duplicates
				uniqueSlice1 = append(uniqueSlice1, string(item1)) // 1. add unique
			}
		}
		var slice1AsString string
		for _, letter1 := range uniqueSlice1 {
			slice1AsString += letter1 // 3. make uniqueSlice1 items into a single string
		}

		slice2 := items[1]
		var uniqueSlice2 []string
		for _, item2 := range slice2 {
			if !slices.Contains(uniqueSlice2, string(item2)) {
				uniqueSlice2 = append(uniqueSlice2, string(item2))
			}
		}
		var slice2AsString string
		for _, letter2 := range uniqueSlice2 {
			slice2AsString += letter2
		}

		newTwoHalves = append(newTwoHalves, slice1AsString, slice2AsString)
		splitUniqueItems = append(splitUniqueItems, newTwoHalves)
	}

	return splitUniqueItems
}

// Find the repeating letter in each half.
func findRepeats(twoHalves [][]string) []string {
	var commonItem []string
	for _, item := range twoHalves {
		for _, letter0 := range item[0] {
			for _, letter1 := range item[1] {
				if string(letter0) == string(letter1) {
					commonItem = append(commonItem, string(letter0))
					break
				}
			}
		}
	}

	return commonItem
}

// loop over the duplicate letter slice and calculate each letter's ascii value
func getAsciiCode(commonItemSlice []string) []int {
	var alphabetIndexes []int
	for _, i := range commonItemSlice {
		lowercase := i[0] - 96 // conver lowercase lett
		if lowercase <= 26 {
			alphabetIndexes = append(alphabetIndexes, int(lowercase))
		} else {
			uppercase := i[0] - 38
			alphabetIndexes = append(alphabetIndexes, int(uppercase))
		}
	}

	return alphabetIndexes
}

// sum the slice of values
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

	groupedSlices := groupSlices(lineSlice)

	removedDuplicates := cleanDuplicates(groupedSlices)

	commonItemSlice := findRepeats(removedDuplicates)

	alphabetIndexSlice := getAsciiCode(commonItemSlice)

	summedItems := sumSlice(alphabetIndexSlice)
	fmt.Println(summedItems)
}
