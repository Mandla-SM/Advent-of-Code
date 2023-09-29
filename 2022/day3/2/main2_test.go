package main

import (
	"reflect"
	"testing"
)

var lines = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

// Remove duplicate alphabets in each slice.
func TestRemoveDuplicates(t *testing.T) {
	got := removeDuplicates(lines)
	want := []string{
		"vJrwpWtghcsFMf",
		"jqHRNzGDLrsFMfZS",
		"PmdzqrVvwTWBg",
		"wMqvLZHhjbcnSBTQF",
		"tgJRGQcTZ",
		"CrZsJPGzwLmpMD",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// Group the split items into a slice, add the slice to a slice.
var deDupedlines = []string{
	"vJrwpWtghcsFMf",
	"jqHRNzGDLrsFMfZS",
	"PmdzqrVvwTWBg",
	"wMqvLZHhjbcnSBTQF",
	"tgJRGQcTZ",
	"CrZsJPGzwLmpMD",
}

var groupedSlices = [][]string{
	[]string{"vJrwpWtghcsFMf", "jqHRNzGDLrsFMfZS", "PmdzqrVvwTWBg"},
	[]string{"wMqvLZHhjbcnSBTQF", "tgJRGQcTZ", "CrZsJPGzwLmpMD"},
}

func TestGroupSlices(t *testing.T) {
	got := groupSlices(deDupedlines)
	want := groupedSlices
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Output %q not equal to expected %q", got, want)
	}
}

// Parse the whole dataset and make a slice of unique alphabets.
// Needs improvement
func TestGetUniqueAlphabets(t *testing.T) {
	got := getUniqueAlphabets(deDupedlines)
	want := []rune{
		'v', 'J', 'r', 'w', 'p', 'W', 't', 'g', 'h', 'c',
		's', 'F', 'M', 'f', 'j', 'q', 'H', 'R', 'N', 'z',
		'G', 'D', 'L', 'Z', 'S', 'P', 'm', 'd', 'V', 'T',
		'B', 'b', 'n', 'Q', 'C',
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Output %q not equal to expected %q", got, want)
	}
}

// Find the repeating alphabet in each group of three.
var commonAlphabets = []string{"r", "Z"}

func TestFindRepeats(t *testing.T) {
	got := findRepeats(groupedSlices, getUniqueAlphabets(deDupedlines))
	want := commonAlphabets

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Output %q not equal to expected %q", want, got)
	}
}

// Loop over the alphabet slice and calculate each alphabet's ascii value.
func TestGetAsciiCode(t *testing.T) {
	got := getAsciiCode(commonAlphabets)
	want := []int{18, 52}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Output %q not equal to expected %q", want, got)
	}
}
