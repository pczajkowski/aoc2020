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
	x int
	y int
}

func makeMove(path []string) position {
	currentPosition := position{x: 0, y: 0}

	for _, item := range path {
		switch item {
		case "e":
			currentPosition.x += 2
		case "w":
			currentPosition.x -= 2
		case "se":
			currentPosition.x += 1
			currentPosition.y -= 1
		case "sw":
			currentPosition.x -= 1
			currentPosition.y -= 1
		case "nw":
			currentPosition.x -= 1
			currentPosition.y += 1
		case "ne":
			currentPosition.x += 1
			currentPosition.y += 1
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

func part1(tiles map[position]int) map[position]int {
	blackTiles := make(map[position]int)

	for key, value := range tiles {
		if value%2 != 0 {
			blackTiles[key] = 1
		}
	}

	return blackTiles
}

func findNeighbours(tile position) map[position]int {
	neighbours := make(map[position]int)

	neighbours[position{x: tile.x - 2, y: tile.y}] = 0
	neighbours[position{x: tile.x + 2, y: tile.y}] = 0
	neighbours[position{x: tile.x + 1, y: tile.y - 1}] = 0
	neighbours[position{x: tile.x - 1, y: tile.y - 1}] = 0
	neighbours[position{x: tile.x - 1, y: tile.y + 1}] = 0
	neighbours[position{x: tile.x + 1, y: tile.y + 1}] = 0

	return neighbours
}

func numberOfBlackNeighbours(neighbours map[position]int, blackTiles map[position]int) int {
	black := 0
	for neighbour, _ := range neighbours {
		if _, ok := blackTiles[neighbour]; ok {
			black++
		}
	}

	return black
}

func flip(blackTiles map[position]int) map[position]int {
	newBlackTiles := make(map[position]int)

	for key, _ := range blackTiles {
		neighbours := findNeighbours(key)
		blackNeighbours := numberOfBlackNeighbours(neighbours, blackTiles)
		if blackNeighbours > 0 && blackNeighbours <= 2 {
			newBlackTiles[key] = 0
		}

		for neighbour, _ := range neighbours {
			if _, ok := blackTiles[neighbour]; ok {
				continue
			}

			anotherNeighbours := findNeighbours(neighbour)
			if numberOfBlackNeighbours(anotherNeighbours, blackTiles) == 2 {
				newBlackTiles[neighbour] = 0
			}
		}
	}

	return newBlackTiles
}

func do100Flips(blackTiles map[position]int) int {
	for i := 0; i < 100; i++ {
		blackTiles = flip(blackTiles)
	}
	return len(blackTiles)
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

	tiles := makeAllMoves(paths)
	blackTiles := part1(tiles)
	fmt.Println("Part1:", len(blackTiles))

	fmt.Println("Part2:", do100Flips(blackTiles))
}
