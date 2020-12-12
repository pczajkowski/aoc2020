package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	numberOfDirecrtions = 4
)

type sequence struct {
	action string
	value  int
}

func readFile(file *os.File) ([]sequence, error) {
	var sequences []sequence
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p sequence
		line := scanner.Text()
		if line == "" {
			break
		}

		n, err := fmt.Sscanf(line, "%1s%d\n", &p.action, &p.value)
		if err != nil || n != 2 {
			return sequences, fmt.Errorf("Error scanning '%s': %s", line, err)
		}

		sequences = append(sequences, p)
	}
	if err := scanner.Err(); err != nil {
		return sequences, fmt.Errorf("Scanner error: %s", err)
	}

	return sequences, nil
}

var directions []string
var directionsMap map[string]int

func getDirections() {
	directions = []string{"E", "S", "W", "N"}

	directionsMap = make(map[string]int)
	for index, value := range directions {
		directionsMap[value] = index
	}
}

func rotate(currentDirection string, item sequence) string {
	change := 0

	switch item.value {
	case 0:
	case 360:
		return currentDirection
	case 90:
		change = 1
	case 180:
		change = 2
	case 270:
		change = 3
	}

	if item.action == "L" {
		change = -change
	}

	newDirectionIndex := (directionsMap[currentDirection] + change) % 4
	if newDirectionIndex < 0 {
		newDirectionIndex = numberOfDirecrtions + newDirectionIndex
	}

	return directions[newDirectionIndex]
}

type cooridinates struct {
	north int
	south int
	east  int
	west  int
}

func makeMove(item sequence, position cooridinates) cooridinates {
	switch item.action {
	case "E":
		if position.west > 0 {
			position.west = position.west - item.value
			if position.west < 0 {
				position.east -= position.west
				position.west = 0
			}
		} else {
			position.east += item.value
		}
	case "W":
		if position.east > 0 {
			position.east = position.east - item.value
			if position.east < 0 {
				position.west -= position.east
				position.east = 0
			}
		} else {
			position.west += item.value
		}
	case "N":
		if position.south > 0 {
			position.south = position.south - item.value
			if position.south < 0 {
				position.north -= position.south
				position.south = 0
			}
		} else {
			position.north += item.value
		}
	case "S":
		if position.north > 0 {
			position.north = position.north - item.value
			if position.north < 0 {
				position.south -= position.north
				position.north = 0
			}
		} else {
			position.south += item.value
		}
	}

	return position
}

func navigate(sequences []sequence) int {
	currentDirection := "E"
	currentPosition := cooridinates{north: 0, south: 0, east: 0, west: 0}

	for _, item := range sequences {
		switch item.action {
		case "L", "R":
			currentDirection = rotate(currentDirection, item)
		case "E", "W", "N", "S":
			currentPosition = makeMove(item, currentPosition)
		case "F":
			item.action = currentDirection
			currentPosition = makeMove(item, currentPosition)
		}
	}

	return currentPosition.north + currentPosition.south + currentPosition.east + currentPosition.west
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open %s!\n", os.Args[1])

	}

	sequences, err := readFile(file)
	if err != nil {
		log.Fatalf("Can't read sequences!\n%s", err)
	}
	file.Close()

	getDirections()
	fmt.Println("Part1:", navigate(sequences))
}
