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

func inSequence(value int, sequence []int) bool {
	for _, item := range sequence {
		if item == value {
			return true
		}
	}

	return false
}

func indexOf(value int, sequence []int) int {
	for i, item := range sequence {
		if item == value {
			return i
		}
	}

	return -1
}

func part1(sequence []int, min, max int) {
	index := 0

	for i := 0; i < 10; i++ {
		pickup := sequence[index+1 : index+4]
		var partialSequence []int
		for j, item := range sequence {
			if j > index && j < index+4 {
				continue
			}

			partialSequence = append(partialSequence, item)
		}

		destination := sequence[index] - 1
		fmt.Println(destination)
		for {
			if !inSequence(destination, pickup) {
				break
			}

			destination--
			if destination < min {
				_, newDestination := minMax(partialSequence)
				destination = newDestination
				break
			}
		}

		fmt.Println(destination)
		destinationIndex := indexOf(destination, partialSequence)
		if destinationIndex < 0 {
			log.Fatalf("Wrong destinationIndex: %d", destinationIndex)
		}

		var newSequence []int
		for j, item := range partialSequence {
			newSequence = append(newSequence, item)

			if j == destinationIndex {
				for _, cup := range pickup {
					newSequence = append(newSequence, cup)
				}
			}
		}

		sequence = newSequence
		index++
		if index < 0 {
			index = 0
		}
		fmt.Println(sequence)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a sequence!")
	}

	sequence := processSequence(os.Args[1])
	min, max := minMax(sequence)

	part1(sequence, min, max)
}
