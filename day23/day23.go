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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a sequence!")
	}

	fmt.Println(processSequence(os.Args[1]))
}
