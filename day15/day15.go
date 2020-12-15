package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var rounds map[int]int
var numbersSpoken map[int][2]int

func readNumbers(startingNumbers string) {
	for i, item := range strings.Split(string(startingNumbers), ",") {
		var number int
		n, err := fmt.Sscanf(item, "%d", &number)
		if err != nil || n < 1 {
			log.Fatal(err)
		}

		rounds[i+1] = number
		numbersSpoken[number] = [2]int{i + 1, 0}
	}
}

func playGame(limit int) int {
	currentRound := len(rounds) + 1

	var currentNumber int
	for ; currentRound <= limit; currentRound++ {
		lastNumber := rounds[currentRound-1]

		if spoken, ok := numbersSpoken[lastNumber]; !ok {
			currentNumber = 0
		} else {
			if spoken[1] == 0 {
				currentNumber = 0

			} else {
				currentNumber = spoken[1] - spoken[0]
			}
		}

		rounds[currentRound] = currentNumber

		if _, ok := numbersSpoken[currentNumber]; !ok {
			numbersSpoken[currentNumber] = [2]int{currentRound, 0}
		} else {
			if numbersSpoken[currentNumber][1] == 0 {
				numbersSpoken[currentNumber] = [2]int{numbersSpoken[currentNumber][0], currentRound}
			} else {
				numbersSpoken[currentNumber] = [2]int{numbersSpoken[currentNumber][1], currentRound}
			}
		}
	}

	return currentNumber
}

func init() {
	rounds = make(map[int]int)
	numbersSpoken = make(map[int][2]int)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify starting numbers!")
	}

	readNumbers(os.Args[1])
	fmt.Println("Part1:", playGame(2020))
	fmt.Println("Part2:", playGame(30000000))
}
