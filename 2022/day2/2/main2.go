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
	fileData, err := os.Open("../plays.txt") // Open the file
	errorExit(err)                           // Exit the program if opening the file fails

	fileLines := bufio.NewScanner(fileData) // Get lines in the file
	fileLines.Split(bufio.ScanLines)        // Split the lines. In place mutate fileLines

	var lineSlice []string // An empty string slice
	for fileLines.Scan() { // Scan() makes fileLines iterable
		lineSlice = append(lineSlice, fileLines.Text()) // Append each line to lineSlice
	}
	fileData.Close() // Close the file

	return lineSlice // Return lineSlice
}

func play(pOne string, condition string) int {
	rock := 1
	paper := 2
	scissors := 3
	lose := 0
	draw := 3
	win := 6
	var playerTwoPoints int

	if pOne == "A" && condition == "X" {
		playerTwoPoints = scissors + lose
	} else if pOne == "A" && condition == "Y" {
		playerTwoPoints = rock + draw
	} else if pOne == "A" && condition == "Z" {
		playerTwoPoints = paper + win
	} else if pOne == "B" && condition == "X" {
		playerTwoPoints = rock + lose
	} else if pOne == "B" && condition == "Y" {
		playerTwoPoints = paper + draw
	} else if pOne == "B" && condition == "Z" {
		playerTwoPoints = scissors + win
	} else if pOne == "C" && condition == "X" {
		playerTwoPoints = paper + lose
	} else if pOne == "C" && condition == "Y" {
		playerTwoPoints = scissors + draw
	} else if pOne == "C" && condition == "Z" {
		playerTwoPoints = rock + win
	}

	return playerTwoPoints
}

func main() {
	inputFile := "../plays.txt"
	totalPoints := 0
	for _, i := range readFile(inputFile) {
		points := play(string(i[0]), string(i[2]))
		totalPoints += points
	}
	fmt.Println(totalPoints)
}
