package main

import (
	"reflect"
	"testing"
)

var lines = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestGetBiggest(t *testing.T) {
	got := getBiggest(lines)
	want := 4
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %q. wanted %q", got, want)
	}
}
