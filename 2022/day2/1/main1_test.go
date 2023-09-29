package main

import (
	"reflect"
	"testing"
)

// Testing main.sortSlice
func TestReadFile(t *testing.T) {
	inputFile := "../plays_demo.txt"
	got := readFile(inputFile)
	want := []string{"A Y", "B X", "C Z"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

type comboTest struct {
	arg1     string
	arg2     string
	expected int
}

var comboTests = []comboTest{
	comboTest{"A", "Y", 8},
	comboTest{"B", "X", 1},
	comboTest{"C", "Z", 6},
}

func TestPlay(t *testing.T) {
	for _, input := range comboTests {
		if output := play(input.arg1, input.arg2); output != input.expected {
			t.Errorf("Output %q not equal to expected %q", output, input.expected)
		}
	}
}
