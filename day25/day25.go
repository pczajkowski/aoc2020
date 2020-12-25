package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(file *os.File) []int {
	var keys []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		key, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatalf("Error processing key for %s: %s", line, err)
		}

		keys = append(keys, key)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	return keys
}

func establishLoopSize(key int) int {
	iterations := 0
	value := 1

	for value != key {
		value *= 7
		value %= 20201227

		iterations++
	}

	return iterations
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

	keys := readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	key1LoopSize := establishLoopSize(keys[0])
	key2LoopSize := establishLoopSize(keys[1])

	fmt.Println(key1LoopSize, key2LoopSize)
}
