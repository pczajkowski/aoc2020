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
	id      string
	idValue int64
	value   int64
}

type schedule struct {
	timestamp int64
	buses     []bus
}

func getIDs(busesString string) ([]bus, error) {
	buses := []bus{}
	for _, item := range strings.Split(busesString, ",") {
		newBus := bus{id: item, value: 0}
		if newBus.id == "x" {
			newBus.idValue = 0
		} else {
			var err error
			newBus.idValue, err = strconv.ParseInt(item, 10, 32)
			if err != nil {
				return buses, fmt.Errorf("Error parsing busID %s: %s", item, err)
			}
		}

		buses = append(buses, newBus)
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

func calculateTimes(data schedule) schedule {
	for i, item := range data.buses {
		if item.id == "x" {
			continue
		}
		data.buses[i].value = item.idValue - (data.timestamp % item.idValue)
	}

	return data
}

func findEarliestBus(data schedule) int64 {
	var earliest int64 = data.timestamp
	var earliestID int64 = 0
	for _, item := range data.buses {
		if item.value < earliest {
			if item.id == "x" {
				continue
			}

			earliest = item.value
			earliestID = item.idValue
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

	data = calculateTimes(data)
	fmt.Println("Part1:", findEarliestBus(data))
}
