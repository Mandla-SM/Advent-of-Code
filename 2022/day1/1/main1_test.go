package main

import (
	"reflect"
	"testing"
)

// Testing main.readFile
func TestReadFile(t *testing.T) {
	inputFile := "../calories_demo.txt"
	got := readFile(inputFile)
	want := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// Testing main.sumSliceItems
func TestSumSliceItems(t *testing.T) {
	got := sumSliceItems([]string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"})
	want := []int{1000, 3000, 6000, 4000, 5000, 11000, 7000, 15000, 24000, 10000}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// Testing main.getBiggest

func TestGetBiggest(t *testing.T) {
	got := getBiggest([]int{1000, 3000, 6000, 4000, 5000, 11000, 7000, 15000, 24000, 10000})
	want := 24000

	if got != want {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}
