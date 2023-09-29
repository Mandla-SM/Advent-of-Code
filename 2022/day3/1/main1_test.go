package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	inputFile := "../items_demo.txt"
	got := readFile(inputFile)
	want := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// split list item into parts
type splitInTwoTest struct {
	arg1      string
	expected1 string
	expected2 string
}

var splitInTwoTests = []splitInTwoTest{
	splitInTwoTest{"vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp"},
	splitInTwoTest{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
	splitInTwoTest{"PmmdzqPrVvPwwTWBwg", "PmmdzqPrV", "vPwwTWBwg"},
	splitInTwoTest{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
	splitInTwoTest{"ttgJtRGJQctTZtZT", "ttgJtRGJ", "QctTZtZT"},
	splitInTwoTest{"CrZsJsPPZsGzwwsLwLmpwMDw", "CrZsJsPPZsGz", "wwsLwLmpwMDw"},
}

func TestSplitInTwo(t *testing.T) {
	for _, input := range splitInTwoTests {
		if output1, output2 := splitInTwo(input.arg1); output1 != input.expected1 && output2 != input.expected2 {
			t.Errorf("Output %q and %q not equal to expected %q and %q", output1, output2, input.expected1, input.expected2)
		}
	}
}

// Group items into slices, add slices to a slice
func TestGroupSlices(t *testing.T) {
	got := groupSlices([]string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"})
	want := [][]string{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
		{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
		{"ttgJtRGJ", "QctTZtZT"},
		{"CrZsJsPPZsGz", "wwsLwLmpwMDw"}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// Remove duplicate letters in each half.
type cleanDuplicatesTest struct {
	arg1     []string
	expected [][]string
}

var cleanDuplicatesTests = []cleanDuplicatesTest{
	cleanDuplicatesTest{[]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, [][]string{{"vJrwpWtg", "hcsFMfp"}}},
	cleanDuplicatesTest{[]string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}, [][]string{{"jqHRNzGDL", "rsFMfZSL"}}},
	cleanDuplicatesTest{[]string{"PmmdzqPrV", "vPwwTWBwg"}, [][]string{{"PmdzqrV", "vPwTWBg"}}},
	cleanDuplicatesTest{[]string{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"}, [][]string{{"wMqvLZHh", "jbvcnSBTQF"}}},
	cleanDuplicatesTest{[]string{"ttgJtRGJ", "QctTZtZT"}, [][]string{{"tgJRG", "QctTZ"}}},
	cleanDuplicatesTest{[]string{"CrZsJsPPZsGz", "wwsLwLmpwMDw"}, [][]string{{"CrZsJPGz", "wsLmpMD"}}},
}

func TestCleanDuplicates(t *testing.T) {
	for _, input := range cleanDuplicatesTests {
		singleItem := input.arg1
		twoDSlice := [][]string{singleItem}
		if !reflect.DeepEqual(cleanDuplicates(twoDSlice), input.expected) {
			t.Errorf("Got %q. wanted %q", cleanDuplicates(twoDSlice), input.expected)
		}
	}
}

// Find the repeating letter in each half.
type findRepeatsTest struct {
	arg1     []string
	expected []string
}

var findRepeatsTests = []findRepeatsTest{
	findRepeatsTest{[]string{"vJrwpWtg", "hcsFMfp"}, []string{"p"}},
	findRepeatsTest{[]string{"jqHRNzGDL", "rsFMfZSL"}, []string{"L"}},
	findRepeatsTest{[]string{"PmdzqrV", "vPwTWBg"}, []string{"P"}},
	findRepeatsTest{[]string{"wMqvLZHh", "jbvcnSBTQF"}, []string{"v"}},
	findRepeatsTest{[]string{"tgJRG", "QctTZ"}, []string{"t"}},
	findRepeatsTest{[]string{"CrZsJPGz", "wsLmpMD"}, []string{"s"}},
}

func TestFindRepeats(t *testing.T) {
	for _, input := range findRepeatsTests {
		singleItem := input.arg1
		twoDSlice := [][]string{singleItem}
		if !reflect.DeepEqual(findRepeats(twoDSlice), input.expected) {
			t.Errorf("Got %q. wanted %q", findRepeats(twoDSlice), input.expected)
		}
	}
}

// loop over the duplicate letter slice and calculate each letter's ascii value
type getAsciiCodeTest struct {
	arg1     []string
	expected []int
}

var getAsciiCodeTests = []getAsciiCodeTest{
	getAsciiCodeTest{[]string{"p", "L", "P", "v", "t", "s"}, []int{16, 38, 42, 22, 20, 19}},
}

func TestGetAsciiCode(t *testing.T) {
	for _, input := range getAsciiCodeTests {
		// fmt.Println(getAsciiCode(input.arg1))
		if !reflect.DeepEqual(getAsciiCode(input.arg1), input.expected) {
			t.Errorf("Output of %q not equal to the expected %q", getAsciiCode(input.arg1), input.expected)
		}
	}
}

// sum the slice of values
func TestSumSlice(t *testing.T) {
	got := sumSlice([]int{16, 38, 42, 22, 20, 19})
	want := 157

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}
