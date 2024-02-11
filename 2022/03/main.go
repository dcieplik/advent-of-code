package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dcieplik/advent-of-code/2022/03/tokenizer"
)

func main() {
	file, err := os.Open("rucksack.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var count int

	for scanner.Scan() {
		count += tokenizer.Calculate(scanner.Text())

	}

	fmt.Println(count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
