package main

import (
	"reflect"
	"testing"
)

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
	want := map[int]string{1: "M", 2: "C", 3: "PZND"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}
