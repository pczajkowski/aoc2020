package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for _, item := range sequences {
		fmt.Println(item)
	}

	file.Close()
}
