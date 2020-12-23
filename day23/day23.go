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
	size := len(sequence)

	for i := 0; i < 10; i++ {
		pickup := sequence[index+1 : index+4]
		for j, _ := range sequence {
			if j > index && j < index+4 {
				sequence[j] = 0
			}
		}

		destination := sequence[index] - 1
		fmt.Println(destination)
		for {
			if !inSequence(destination, pickup) {
				break
			}

			destination--
			if destination < min {
				_, newDestination := minMax(sequence)
				destination = newDestination
				break
			}
		}

		fmt.Println(sequence)
		fmt.Println(destination)

		partialSequence := make([]int, size)
		i := index + 1
		for {
			partialSequence[i] = sequence[i]
			if i == index {
				i++
				break
			}
			i++
			if i > size-1 {
				i = 0
			}
		}

		destinationIndex := indexOf(destination, partialSequence)
		if destinationIndex < 0 {
			log.Fatalf("Wrong destinationIndex: %d", destinationIndex)
		}

		newSequence := make([]int, size)
		i = destinationIndex
		for j := destinationIndex; ; j++ {
			newSequence[j] = partialSequence[i]

			if j == destinationIndex {
				j++
				for _, cup := range pickup {
					newSequence[j] = cup
					j++
				}

				i++
				continue
			}

			j++
			i++
		}

		sequence = newSequence
		index++
		if index < 0 {
			index = 0
		}
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
