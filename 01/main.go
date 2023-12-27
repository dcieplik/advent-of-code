package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("elves.txt")
	if err != nil {
		panic(err)
	}

	func() {
		err := readFile.Close()
		if err != nil {
			os.Exit(1)
		}
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var calories []int

	tmp := 0

	for fileScanner.Scan() {

		line := fileScanner.Text()
		if line == "" {
			calories = append(calories, tmp)
			tmp = 0
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		tmp += i
	}

	calories = append(calories, tmp)

	// sort calories
	sort.Ints(calories)

	length := len(calories)
	calories = calories[length-3 : length]

	var top3 int
	for _, v := range calories {
		top3 += v
	}

	fmt.Println("Top calories:", top3)
}
