package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readDirections(line string) []string {
	var path []string
	signs := []rune(line)
	size := len(signs)

	for i := 0; i < size; i++ {
		switch signs[i] {
		case 'e', 'w':
			path = append(path, string(signs[i]))
		case 's', 'n':
			path = append(path, string(signs[i:i+2]))
			i++
		}
	}

	return path
}

func readFile(file *os.File) [][]string {
	var paths [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		paths = append(paths, readDirections(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	return paths
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

	paths := readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	fmt.Println(paths)
}
