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

func play(decks [2][]int) []int {
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
		return decks[1]
	}
	return decks[0]
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

	fmt.Println(play(decks))
}
