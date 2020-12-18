package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getTokens(line string) ([]rune, error) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	var tokens []rune
	for scanner.Scan() {
		newTokens := []rune(scanner.Text())
		tokens = append(tokens, newTokens...)
	}
	if err := scanner.Err(); err != nil {
		return tokens, fmt.Errorf("Scanner error: %s", err)
	}

	return tokens, nil
}

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tokens, err := getTokens(line)
		if err != nil {
			log.Fatalf("Error scanning %s: %s", line, err)
		}

		for _, token := range tokens {
			fmt.Print(string(token), ",")
		}
		fmt.Println()

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
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

	readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}
}
