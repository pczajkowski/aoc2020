package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	id    int64
	index int64
}

type schedule struct {
	timestamp int64
	buses     []bus
}

func getIDs(busesString string) ([]bus, error) {
	buses := []bus{}
	var index int64 = 0
	for _, item := range strings.Split(busesString, ",") {
		if item != "x" {
			id, err := strconv.ParseInt(item, 10, 32)
			if err != nil {
				return buses, fmt.Errorf("Error parsing busID %s: %s", item, err)
			}
			newBus := bus{id: id, index: index}
			buses = append(buses, newBus)
		}

		index++
	}

	return buses, nil
}

func readData(file *os.File) (schedule, error) {
	scanner := bufio.NewScanner(file)
	data := schedule{}

	if !scanner.Scan() {
		return data, fmt.Errorf("Error reading timestamp!")
	}
	timestampString := scanner.Text()

	timestamp, err := strconv.ParseInt(timestampString, 10, 32)
	if err != nil {
		return data, fmt.Errorf("Error parsing timestamp %s: %s", timestampString, err)
	}
	data.timestamp = timestamp

	if !scanner.Scan() {
		return data, fmt.Errorf("Error reading buses!")
	}
	busesString := scanner.Text()
	if err := scanner.Err(); err != nil {
		return data, err
	}

	data.buses, err = getIDs(busesString)
	if err != nil {
		return data, err
	}

	return data, nil
}

func findEarliestBus(data schedule) int64 {
	var earliest int64 = data.timestamp
	var earliestID int64 = 0
	for _, item := range data.buses {
		value := item.id - (data.timestamp % item.id)
		if value < earliest {

			earliest = value
			earliestID = item.id
		}
	}

	return earliest * earliestID
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open %s!\n", os.Args[1])

	}

	data, err := readData(file)
	if err != nil {
		log.Fatalf("Failed to read data: %s\n", err)
	}

	fmt.Println("Part1:", findEarliestBus(data))
}
