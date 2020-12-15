package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readNumbers(startingNumbers string) (int, map[int][2]int) {
	lastNumber := 0
	numbersSpoken := make(map[int][2]int)
	for i, item := range strings.Split(string(startingNumbers), ",") {
		var number int
		n, err := fmt.Sscanf(item, "%d", &number)
		if err != nil || n < 1 {
			log.Fatal(err)
		}

		lastNumber = number
		numbersSpoken[number] = [2]int{i + 1, 0}
	}

	return lastNumber, numbersSpoken
}

func playGame(currentRound, end, lastNumber int, numbersSpoken map[int][2]int) (int, map[int][2]int) {
	var currentNumber int
	for ; currentRound <= end; currentRound++ {
		if spoken, ok := numbersSpoken[lastNumber]; !ok {
			currentNumber = 0
		} else {
			if spoken[1] == 0 {
				currentNumber = 0

			} else {
				currentNumber = spoken[1] - spoken[0]
			}
		}

		if _, ok := numbersSpoken[currentNumber]; !ok {
			numbersSpoken[currentNumber] = [2]int{currentRound, 0}
		} else {
			if numbersSpoken[currentNumber][1] == 0 {
				numbersSpoken[currentNumber] = [2]int{numbersSpoken[currentNumber][0], currentRound}
			} else {
				numbersSpoken[currentNumber] = [2]int{numbersSpoken[currentNumber][1], currentRound}
			}
		}

		lastNumber = currentNumber
	}

	return currentNumber, numbersSpoken
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify starting numbers!")
	}

	lastNumber, numbersSpoken := readNumbers(os.Args[1])
	lastNumber, numbersSpoken = playGame(len(numbersSpoken)+1, 2020, lastNumber, numbersSpoken)
	fmt.Println("Part1:", lastNumber)

	lastNumber, numbersSpoken = playGame(2021, 30000000, lastNumber, numbersSpoken)
	fmt.Println("Part2:", lastNumber)
}
