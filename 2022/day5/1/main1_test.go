package main

import (
	"reflect"
	"testing"
)

func TestStackConverter(t *testing.T) {
	inputFile := "../stacks_demo.txt"
	lineSlice := readFile(inputFile)
	got := stackConverter(lineSlice)
	want := map[int]string{1: "NZ", 2: "DCM", 3: "P"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

func TestOrderedValues(t *testing.T) {
	got := orderedValues("DCM")
	want := "MCD"

	if got != want {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

func TestMoveItems(t *testing.T) {
	inputFile1 := "../instructions_demo.txt"
	lineSlice1 := readFile(inputFile1)
	inputFile2 := "../stacks_demo.txt"
	lineSlice2 := readFile(inputFile2)

	generatedStacks := stackConverter(lineSlice2)

	initialStacks := make(map[int]string)
	for key, value := range generatedStacks {
		initialStacks[key] = orderedValues(value)
	}

	got := moveItems(lineSlice1, initialStacks)
	want := map[int]string{1: "C", 2: "M", 3: "PDNZ"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

func TestGetTopItem(t *testing.T) {
	got := getTopItem(map[int]string{1: "C", 2: "M", 3: "PDNZ"})
	want := "CMZ"

	if got != want {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}
