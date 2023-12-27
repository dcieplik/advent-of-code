package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type ScoreFunc func(player string, opponent string) int

func main() {

	r, err := os.Open("strategy.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	score := parsePart1(r)

	fmt.Printf("score %d", score)

}

func parse(r io.Reader, calc ScoreFunc) int {

	scan := bufio.NewScanner(r)
	scan.Split(bufio.ScanLines)

	score := 0

	for scan.Scan() {

		line := scan.Text()
		s := strings.Split(line, " ")

		opponent := s[0]
		player := s[1]

		score += calc(player, opponent)
	}
	return score
}

func parsePart1(r io.Reader) int {
	var score int

	// 'A' 'X' Rock     1
	// 'B' 'Y' Paper    2
	// 'C' 'Z' Scissors 3

	// Rock defeats Scissors,
	// Scissors defeats Paper
	// Paper defeats Rock.

	// DRAW    3
	// LOOSE   0
	// WIN     6

	data := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	scan := bufio.NewScanner(r)

	for scan.Scan() {
		score += data[scan.Text()]
	}

	return score
}

func parsePart2(r io.Reader) int {

	// 'A' Rock     1
	// 'B' Paper    2
	// 'C' Scissors 3

	// Rock defeats Scissors,
	// Scissors defeats Paper
	// Paper defeats Rock.

	// Y DRAW    3
	// X LOOSE   0
	// Z WIN     6

	var score int

	data := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	scan := bufio.NewScanner(r)

	for scan.Scan() {
		score += data[scan.Text()]
	}

	return score
}
