package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rule struct {
	name          string
	firstSegment  [2]int
	secondSegment [2]int
}

var rules []rule

func readRule(line string) {
	var newRule rule
	n, err := fmt.Sscanf(line, "%s %d-%d or %d-%d\n", &newRule.name, &newRule.firstSegment[0], &newRule.firstSegment[1], &newRule.secondSegment[0], &newRule.secondSegment[1])
	if err != nil || n != 5 {
		log.Fatalf("Error scanning '%s': %s", line, err)
	}

	rules = append(rules, newRule)
}

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		readRule(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
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

	readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	fmt.Println(rules)
}
