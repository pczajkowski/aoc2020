package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	name          string
	firstSegment  [2]int
	secondSegment [2]int
}

var rules []rule

func readRule(line string) {
	var newRule rule
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		log.Fatalf("Invalid line: %s", line)
	}

	newRule.name = parts[0]
	n, err := fmt.Sscanf(parts[1], "%d-%d or %d-%d\n", &newRule.firstSegment[0], &newRule.firstSegment[1], &newRule.secondSegment[0], &newRule.secondSegment[1])
	if err != nil || n != 4 {
		log.Fatalf("Error scanning '%s': %s", line, err)
	}

	rules = append(rules, newRule)
}

type ticket []int

var tickets []ticket

func readTicket(line string) {
	var newTicket ticket
	for _, item := range strings.Split(line, ",") {
		field, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("Error parsing field from %s: %s", item, err)
		}

		newTicket = append(newTicket, field)
	}

	tickets = append(tickets, newTicket)
}

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	currentFunction := readRule
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "your ticket:") {
			currentFunction = readTicket
			continue
		}

		if strings.Contains(line, "nearby tickets:") {
			continue
		}

		currentFunction(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
}

func checkAllRulesOnField(field int) bool {
	for _, currentRule := range rules {
		if (field >= currentRule.firstSegment[0] && field <= currentRule.firstSegment[1]) || (field >= currentRule.secondSegment[0] && field <= currentRule.secondSegment[1]) {
			return true
		}
	}

	return false
}

func sumBad() int {
	numberOfTickets := len(tickets)
	sum := 0

	for i := 1; i < numberOfTickets; i++ {
		for _, field := range tickets[i] {
			if !checkAllRulesOnField(field) {
				sum += field
			}
		}
	}

	return sum
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

	fmt.Println("Part1:", sumBad())
}
