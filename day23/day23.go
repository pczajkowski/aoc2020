package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func processSequence(input string) []int {
	var sequence []int

	for _, letter := range input {
		cup, err := strconv.Atoi(string(letter))
		if err != nil {
			log.Fatalf("Error processing cup for %s: %s", letter, err)
		}

		sequence = append(sequence, cup)
	}

	return sequence
}

func minMax(sequence []int) (int, int) {
	max := 0
	min := 9

	for _, cup := range sequence {
		if cup > max {
			max = cup
		}

		if cup < min {
			min = cup
		}
	}

	return min, max
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a sequence!")
	}

	sequence := processSequence(os.Args[1])
	min, max := minMax(sequence)

	fmt.Println(min, max)
}
