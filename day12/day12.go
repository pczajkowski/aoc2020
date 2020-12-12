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
var opposites map[string]string

func getDirections() {
	directions = []string{"E", "S", "W", "N"}

	directionsMap = make(map[string]int)
	for index, value := range directions {
		directionsMap[value] = index
	}

	opposites = make(map[string]string)
	opposites["E"] = "W"
	opposites["W"] = "E"
	opposites["N"] = "S"
	opposites["S"] = "N"
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
	opposite := opposites[item.action]
	if position[opposite] > 0 {
		position[opposite] = position[opposite] - item.value
		if position[opposite] < 0 {
			position[item.action] -= position[opposite]
			position[opposite] = 0
		}
	} else {
		position[item.action] += item.value
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

func rotateWaypoint(waypoint cooridinates, item sequence) cooridinates {
	newWaypoint := make(cooridinates)

	for key, value := range waypoint {
		newKey := rotate(key, item)
		newWaypoint[newKey] = value
	}

	return newWaypoint
}

func navigate2(sequences []sequence) int {
	currentPosition := make(cooridinates)
	currentPosition["E"] = 0
	currentPosition["W"] = 0
	currentPosition["N"] = 0
	currentPosition["S"] = 0

	currentWaypoint := make(cooridinates)
	currentWaypoint["E"] = 10
	currentWaypoint["W"] = 0
	currentWaypoint["N"] = 1
	currentWaypoint["S"] = 0

	for _, item := range sequences {
		switch item.action {
		case "L", "R":
			currentWaypoint = rotateWaypoint(currentWaypoint, item)
		case "E", "W", "N", "S":
			currentWaypoint = makeMove(item, currentWaypoint)
		case "F":
			for key, value := range currentWaypoint {
				currentPosition = makeMove(sequence{action: key, value: value * item.value}, currentPosition)
			}
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
	fmt.Println("Part2:", navigate2(sequences))
}
