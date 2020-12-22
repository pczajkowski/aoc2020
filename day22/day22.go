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

	fmt.Println(decks)
}
