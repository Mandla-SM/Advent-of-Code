package main

import "testing"

type comboTest struct {
	arg1     string
	arg2     string
	expected int
}

var comboTests = []comboTest{
	comboTest{"A", "Y", 4},
	comboTest{"B", "X", 1},
	comboTest{"C", "Z", 7},
}

func TestPlay(t *testing.T) {
	for _, input := range comboTests {
		if output := play(input.arg1, input.arg2); output != input.expected {
			t.Errorf("Output %q not equal to expected %q", output, input.expected)
		}
	}
}
