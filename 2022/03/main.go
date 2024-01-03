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

	fmt.Println(mapping["A"])

	fmt.Println("vim-go")
	file, err := os.Open("rucksack.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		left, reight := tokenizer.Half(scanner.Text())

		fmt.Printf("left %s right %s\n", left, reight)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
