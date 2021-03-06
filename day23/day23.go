package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func getThreeCups(sequence []int, index int) []int {
	count := 0
	length := len(sequence) - 1
	var cups []int
	for count < 3 {
		if index > length {
			index = 0
		}

		cups = append(cups, sequence[index])
		count++
		index++
	}

	return cups
}

func getSequence(sequence []int, min, max, iterations int) []int {
	index := 0
	size := len(sequence)
	length := size - 1

	for iteration := 0; iteration < iterations; iteration++ {
		pickup := getThreeCups(sequence, index+1)

		x := 0
		y := index + 1
		if y > length {
			y = 0
		}
		for x < 3 {
			sequence[y] = 0
			x++
			y++
			if y > length {
				y = 0
			}
		}

		destination := sequence[index] - 1
		for {
			if destination < min {
				_, newDestination := minMax(sequence)
				destination = newDestination
				break
			}

			if !inSequence(destination, pickup) {
				break
			}

			destination--
		}

		newSequence := make([]int, size)
		i := index + 1
		j := index + 4
		if j > length {
			j = j - size
		}
		count := 0
		for count < size {
			count++
			if i > length {
				i = 0
			}

			if j > length {
				j = 0
			}

			newSequence[i] = sequence[j]

			if sequence[j] == destination {
				i++

				for _, cup := range pickup {
					if i > length {
						i = 0
					}

					newSequence[i] = cup
					i++
				}

				count += 3

				j++
				continue
			}

			j++
			i++
		}

		sequence = newSequence
		index++
		if index > length {
			index = 0
		}
	}

	return sequence
}

func part1(sequence []int) string {
	size := len(sequence)
	indexOfOne := indexOf(1, sequence)

	var result []string
	for i := indexOfOne + 1; i < size; i++ {
		result = append(result, fmt.Sprintf("%d", sequence[i]))
	}

	for i := 0; i < indexOfOne; i++ {
		result = append(result, fmt.Sprintf("%d", sequence[i]))
	}

	return strings.Join(result, "")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a sequence!")
	}

	sequence := processSequence(os.Args[1])
	min, max := minMax(sequence)
	finalSequence := getSequence(sequence, min, max, 100)

	fmt.Println("Part1:", part1(finalSequence))
}
