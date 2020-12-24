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

type position struct {
	x float64
	y float64
}

func makeMove(path []string) position {
	currentPosition := position{x: 0, y: 0}

	for _, item := range path {
		switch item {
		case "e":
			currentPosition.x += 1
		case "w":
			currentPosition.x -= 1
		case "se":
			currentPosition.x += 0.5
			currentPosition.y -= 0.5
		case "sw":
			currentPosition.x -= 0.5
			currentPosition.y -= 0.5
		case "nw":
			currentPosition.x -= 0.5
			currentPosition.y += 0.5
		case "ne":
			currentPosition.x += 0.5
			currentPosition.y += 0.5
		}
	}

	return currentPosition
}

func makeAllMoves(paths [][]string) map[position]int {
	moves := make(map[position]int)

	for _, path := range paths {
		currentPosition := makeMove(path)
		moves[currentPosition] += 1
	}

	return moves
}

func part1(moves map[position]int) int {
	black := 0

	for _, value := range moves {
		if value%2 != 0 {
			black++
		}
	}

	return black
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

	moves := makeAllMoves(paths)
	fmt.Println("Part1:", part1(moves))
}
