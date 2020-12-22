package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(file *os.File) [2][]int {
	var decks [2][]int
	index := 0
	changed := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if changed {
				break
			}

			continue
		}

		if line == "Player 1:" {
			continue
		}

		if line == "Player 2:" {
			index++
			changed = true
			continue
		}

		card, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Error processing card for %s: %s", line, err)
		}

		decks[index] = append(decks[index], card)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	return decks
}

func play1(decks [2][]int) (int, []int) {
	for {
		if len(decks[0]) == 0 || len(decks[1]) == 0 {
			break
		}

		player1Hand := decks[0][0]
		decks[0] = decks[0][1:len(decks[0])]

		player2Hand := decks[1][0]
		decks[1] = decks[1][1:len(decks[1])]

		if player1Hand > player2Hand {
			decks[0] = append(decks[0], player1Hand)
			decks[0] = append(decks[0], player2Hand)
		} else {
			decks[1] = append(decks[1], player2Hand)
			decks[1] = append(decks[1], player1Hand)
		}
	}

	if len(decks[0]) == 0 {
		return 1, decks[1]
	}
	return 0, decks[0]
}

func checkDeck(deck []int, deckFromRound []int) bool {
	for i, card := range deck {
		if card != deckFromRound[i] {
			return false
		}
	}
	return true
}

func checkDecks(deck1, deck2 []int, previousRounds []previous) bool {
	for _, round := range previousRounds {
		if len(deck1) != len(round.deck1) || len(deck2) != len(round.deck2) {
			continue
		}
		if checkDeck(deck1, round.deck1) || checkDeck(deck2, round.deck2) {
			return true
		}
	}

	return false
}

type previous struct {
	deck1 []int
	deck2 []int
}

func play2(decks [2][]int) (int, []int) {
	var previousRounds []previous

	for {
		if len(decks[0]) == 0 || len(decks[1]) == 0 {
			break
		}

		if len(previousRounds) > 0 {
			if checkDecks(decks[0], decks[1], previousRounds) {
				return 0, decks[0]
			}
		}

		previousRounds = append(previousRounds, previous{deck1: decks[0], deck2: decks[1]})

		player1Hand := decks[0][0]
		decks[0] = decks[0][1:len(decks[0])]

		player2Hand := decks[1][0]
		decks[1] = decks[1][1:len(decks[1])]

		if len(decks[0]) >= player1Hand && len(decks[1]) >= player2Hand {
			var newDecks [2][]int
			for i, card := range decks[0] {
				if i >= player1Hand {
					break
				}

				newDecks[0] = append(newDecks[0], card)
			}

			for i, card := range decks[1] {
				if i >= player2Hand {
					break
				}

				newDecks[1] = append(newDecks[1], card)
			}

			winner, _ := play2(newDecks)

			if winner == 0 {
				decks[0] = append(decks[0], player1Hand)
				decks[0] = append(decks[0], player2Hand)
			} else {
				decks[1] = append(decks[1], player2Hand)
				decks[1] = append(decks[1], player1Hand)
			}
		} else {
			if player1Hand > player2Hand {
				decks[0] = append(decks[0], player1Hand)
				decks[0] = append(decks[0], player2Hand)
			} else {
				decks[1] = append(decks[1], player2Hand)
				decks[1] = append(decks[1], player1Hand)
			}
		}

	}

	if len(decks[0]) == 0 {
		return 1, decks[1]
	}
	return 1, decks[0]
}

func calculate(deck []int) int {
	result := 0
	multiplyBy := 1
	index := len(deck) - 1

	for ; index >= 0; index-- {
		result += deck[index] * multiplyBy
		multiplyBy++
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)

	}

	decks := readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	_, winningDeck := play1(decks)
	fmt.Println("Part1:", calculate(winningDeck))

	_, winningDeck2 := play2(decks)
	fmt.Println("Part2:", calculate(winningDeck2))
}
