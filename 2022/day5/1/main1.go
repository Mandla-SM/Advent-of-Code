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

// Convert stacks into a dictionary
func stackConverter(lineSlice2 []string) map[int]string {
	generatedStacks := make(map[int]string)
	for index, line := range lineSlice2 {
		if index+1 == len(lineSlice2) { // len(lineSlice2) is the slice with the column numbers
			for _, num := range line {
				if num != 32 { // 32 = white space
					key, err := strconv.Atoi(string(num))
					errorExit(err)
					generatedStacks[key] = "" // set the column number as a key with an empty value
				}
			}
		}
	}
	for index, line := range lineSlice2 { // Run again since generatedStacks needs to already have all keys
		for column, item := range line {
			if item != 32 && index+1 != len(lineSlice2) {
				if column == 1 { // "column" is the array index of the slice string
					generatedStacks[1] += string(item)
				}
				if column == 5 { // 1, 5, 9... are the columns that have letters
					generatedStacks[2] += string(item)
				}
				if column == 9 {
					generatedStacks[3] += string(item)
				}
				if column == 13 {
					generatedStacks[4] += string(item)
				}
				if column == 17 {
					generatedStacks[5] += string(item)
				}
				if column == 21 {
					generatedStacks[6] += string(item)
				}
				if column == 25 {
					generatedStacks[7] += string(item)
				}
				if column == 29 {
					generatedStacks[8] += string(item)
				}
				if column == 33 {
					generatedStacks[9] += string(item)
				}
			}
		}
	}

	return generatedStacks
}

// Reverse the order of the values in the stacks
func orderedValues(value string) string {
	newValues := []byte(value)
	i := 0
	j := len(newValues) - 1
	for i < j {
		newValues[i], newValues[j] = newValues[j], newValues[i]
		i++
		j--
	}

	return string(newValues)
}

// Order the items according to the instructions
// 1. Read each line...
func moveItems(lineSlice []string, initialStacks map[int]string) map[int]string {
	givenMap := initialStacks
	for _, instruction := range lineSlice {
		splitLine := strings.Split(instruction, " ")         // 2. Split it at the spaces
		loopCount, err := strconv.Atoi(string(splitLine[1])) // 3. Get the number of items to move
		from, err := strconv.Atoi(string(splitLine[3]))      // 4. Get the source
		to, err := strconv.Atoi(string(splitLine[5]))        // 5. Get the destination
		errorExit(err)

		for i := 1; i <= loopCount; i++ { // Loop with loopCount as limit
			getTop := givenMap[from][len(givenMap[from])-1:]        // 1. Grab the top item
			givenMap[from] = givenMap[from][:len(givenMap[from])-1] // 2. Set the new pile minus the top item
			givenMap[to] = givenMap[to] + getTop                    // 3. Append the grabbed top item to the destination
		}

	}

	return givenMap
}

// Get the top item in each stack
func getTopItem(stackedItems map[int]string) string {
	keys := len(stackedItems)
	var asString string
	for i := 1; i <= keys; i++ {
		if stackedItems[i] != "" {
			asString += stackedItems[i][len(stackedItems[i])-1:]
		}
	}

	return asString
}

func main() {
	inputFile1 := "../instructions_demo.txt"
	lineSlice1 := readFile(inputFile1)
	inputFile2 := "../stacks_demo.txt"
	lineSlice2 := readFile(inputFile2)

	generatedStacks := stackConverter(lineSlice2)

	initialStacks := make(map[int]string)
	for key, value := range generatedStacks {
		initialStacks[key] = orderedValues(value)
	}

	rearrangedItems := moveItems(lineSlice1, initialStacks)
	fmt.Println(rearrangedItems)

	topItems := getTopItem(rearrangedItems)
	fmt.Println(topItems)
}
