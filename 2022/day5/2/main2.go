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

	var lineSlice1 []string // An empty string slice
	for fileLines.Scan() {  // Scan() makes fileLines iterable
		lineSlice1 = append(lineSlice1, fileLines.Text()) // Append each line to lineSlice1
	}
	fileData.Close() // Close the file

	return lineSlice1 // Return lineSlice1
}

// Convert stacks to dictionary
func stackConverter(lineSlice2 []string) map[int]string {
	generatedStacks := make(map[int]string)
	for index, line := range lineSlice2 {
		if index+1 == len(lineSlice2) {
			for _, num := range line {
				if num != 32 {
					key, err := strconv.Atoi(string(num))
					errorExit(err)
					generatedStacks[key] = ""
				}
			}
		}
	}
	for index, line := range lineSlice2 {
		for column, item := range line {
			if item != 32 && index+1 != len(lineSlice2) {
				if column == 1 {
					generatedStacks[1] += string(item)
				}
				if column == 5 {
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

// 1. Read each line...
func moveItems(lineSlice1 []string, initialStacks map[int]string) map[int]string {
	givenMap := initialStacks
	for _, instruction := range lineSlice1 {
		splitLine := strings.Split(instruction, " ")
		index, err := strconv.Atoi(string(splitLine[1]))
		from, err := strconv.Atoi(string(splitLine[3]))
		to, err := strconv.Atoi(string(splitLine[5]))
		grab := string(givenMap[from][len(givenMap[from])-index:])
		errorExit(err)

		givenMap[from] = givenMap[from][:len(givenMap[from])-index]
		givenMap[to] = givenMap[to] + grab
	}

	return givenMap
}

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
