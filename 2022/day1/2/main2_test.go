package main

import (
	"reflect"
	"testing"
)

// Testing main.sortSlice
func TestSortSlice(t *testing.T) {
	got := sortSlice([]int{1000, 3000, 6000, 4000, 5000, 11000, 7000, 15000, 24000, 10000})
	want := []int{1000, 3000, 4000, 5000, 6000, 7000, 10000, 11000, 15000, 24000}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// Testing main.getTopThree
func TestGetTopThree(t *testing.T) {
	got := getTopThree([]int{1000, 3000, 4000, 5000, 6000, 7000, 10000, 11000, 15000, 24000}, 3)
	want := []int{11000, 15000, 24000}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}

// Testing main.sumSlice
func TestSumSlice(t *testing.T) {
	got := sumSlice([]int{11000, 15000, 24000})
	want := 50000

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}
