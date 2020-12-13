package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type schedule struct {
	timestamp int64
	buses     map[int64]int64
}

func getIDs(busesString string) (map[int64]int64, error) {
	buses := make(map[int64]int64)

	for _, bus := range strings.Split(busesString, ",") {
		if bus == "x" {
			continue
		}

		busID, err := strconv.ParseInt(bus, 10, 32)
		if err != nil {
			return buses, fmt.Errorf("Error parsing busID %s: %s", bus, err)
		}
		buses[busID] = 0
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
	fmt.Println(data)
}
