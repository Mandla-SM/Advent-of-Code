package main

import (
	"bufio"
	"fmt"
	"os"
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

func play(pOne string, pTwo string) int {
	rock := 1
	paper := 2
	scissors := 3
	lose := 0
	draw := 3
	win := 6
	var playerTwoPoints int

	if pOne == "A" && pTwo == "X" {
		playerTwoPoints = rock + draw
	} else if pOne == "A" && pTwo == "Y" {
		playerTwoPoints = paper + win
	} else if pOne == "A" && pTwo == "Z" {
		playerTwoPoints = scissors + lose
	} else if pOne == "B" && pTwo == "X" {
		playerTwoPoints = rock + lose
	} else if pOne == "B" && pTwo == "Y" {
		playerTwoPoints = paper + draw
	} else if pOne == "B" && pTwo == "Z" {
		playerTwoPoints = scissors + win
	} else if pOne == "C" && pTwo == "X" {
		playerTwoPoints = rock + win
	} else if pOne == "C" && pTwo == "Y" {
		playerTwoPoints = paper + lose
	} else if pOne == "C" && pTwo == "Z" {
		playerTwoPoints = scissors + draw
	}

	return playerTwoPoints
}

func main() {
	inputFile := "../plays_demo.txt"
	totalPoints := 0
	for _, i := range readFile(inputFile) {
		points := play(string(i[0]), string(i[2]))
		totalPoints += points
	}
	fmt.Println(totalPoints)
}
