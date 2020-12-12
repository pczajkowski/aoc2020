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
	case 0, 360:
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

type cooridinates map[string]int

func makeMove(item sequence, position cooridinates) cooridinates {
	switch item.action {
	case "E":
		if position["W"] > 0 {
			position["W"] = position["W"] - item.value
			if position["W"] < 0 {
				position["E"] -= position["W"]
				position["W"] = 0
			}
		} else {
			position["E"] += item.value
		}
	case "W":
		if position["E"] > 0 {
			position["E"] = position["E"] - item.value
			if position["E"] < 0 {
				position["W"] -= position["E"]
				position["E"] = 0
			}
		} else {
			position["W"] += item.value
		}
	case "N":
		if position["S"] > 0 {
			position["S"] = position["S"] - item.value
			if position["S"] < 0 {
				position["N"] -= position["S"]
				position["S"] = 0
			}
		} else {
			position["N"] += item.value
		}
	case "S":
		if position["N"] > 0 {
			position["N"] = position["N"] - item.value
			if position["N"] < 0 {
				position["S"] -= position["N"]
				position["N"] = 0
			}
		} else {
			position["S"] += item.value
		}
	}

	return position
}

func navigate(sequences []sequence) int {
	currentDirection := "E"
	currentPosition := make(cooridinates)
	currentPosition["E"] = 0
	currentPosition["W"] = 0
	currentPosition["N"] = 0
	currentPosition["S"] = 0

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

	sum := 0
	for _, value := range currentPosition {
		sum += value
	}
	return sum
}

func rotateWaypoint(waypoint cooridinates) cooridinates {

	return waypoint
}

func navigate2(sequences []sequence) int {
	currentPosition := make(cooridinates)
	currentPosition["E"] = 0
	currentPosition["W"] = 0
	currentPosition["N"] = 0
	currentPosition["S"] = 0
	//currentWaypoint := cooridinates{north: 1, south: 0, east: 10, west: 0}

	for _, item := range sequences {
		switch item.action {
		case "L", "R":
			//currentDirection = rotate(currentDirection, item)
		case "E", "W", "N", "S":
			//currentPosition = makeMove(item, currentPosition)
		case "F":
			//item.action = currentDirection
			//currentPosition = makeMove(item, currentPosition)
		}
	}

	sum := 0
	for _, value := range currentPosition {
		sum += value
	}
	return sum
}

func init() {
	getDirections()
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

	fmt.Println("Part1:", navigate(sequences))
}
