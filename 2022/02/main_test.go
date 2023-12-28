package main

import (
	"os"
	"testing"
)

func TestParse02Part1(t *testing.T) {

	r, err := os.Open("./testdata/part2.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	wanted := 13005
	got := parsePart1(r)

	if got != wanted {
		t.Fatalf("expected %d but got %d", wanted, got)
	}
}

func TestParse02Part2(t *testing.T) {

	r, err := os.Open("./testdata/part2.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	got := parsePart2(r)

	expected := 11373

	if got != expected {
		t.Fatalf("expected %d but got %d", expected, got)
	}

}
