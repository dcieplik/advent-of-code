package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dcieplik/advent-of-code/2022/03/tokenizer"
)

var mapping = make(map[string]int)

func main() {

	for i, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		mapping[string(char)] = i + 1
	}

	file, err := os.Open("rucksack.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		left, right := tokenizer.Half(scanner.Text())
		fmt.Printf("left: %s %s", left, right)
		duplicates := tokenizer.FindDuplicateChars(left, right)
		fmt.Println("dd:", duplicates)

		for _, char := range duplicates {
			count += mapping[string(char)]
		}

		// fmt.Printf("left %s right %s\n", left, reight)
	}

	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
