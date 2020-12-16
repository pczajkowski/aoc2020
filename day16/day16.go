package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

var validTickets []ticket

func sumBad() int {
	numberOfTickets := len(tickets)
	sum := 0

	for i := 1; i < numberOfTickets; i++ {
		validTicket := true
		for _, field := range tickets[i] {
			if !checkAllRulesOnField(field) {
				sum += field
				validTicket = false
			}
		}

		if validTicket {
			validTickets = append(validTickets, tickets[i])
		}
	}

	validTickets = append(validTickets, tickets[0])
	return sum
}

func checkRuleOnField(currentRule rule, field int) bool {
	if (field >= currentRule.firstSegment[0] && field <= currentRule.firstSegment[1]) || (field >= currentRule.secondSegment[0] && field <= currentRule.secondSegment[1]) {
		return true
	}

	return false
}

type fieldRule struct {
	id    int
	size  int
	rules []rule
}

type bySize []fieldRule

func (a bySize) Len() int           { return len(a) }
func (a bySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySize) Less(i, j int) bool { return a[i].size < a[j].size }

func establishOrder() {
	numberOfFields := len(rules)
	var rulesByField []fieldRule

	for i := 0; i < numberOfFields; i++ {
		current := fieldRule{id: i, size: 0}
		for _, currentRule := range rules {
			ruleValid := true
			for _, item := range validTickets {
				if !checkRuleOnField(currentRule, item[i]) {
					ruleValid = false
					break
				}
			}

			if ruleValid {
				current.rules = append(current.rules, currentRule)
				current.size++
			}
		}
		rulesByField = append(rulesByField, current)
	}

	sort.Sort(bySize(rulesByField))
	for _, item := range rulesByField {
		fmt.Println(item)
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

	fmt.Println("Part1:", sumBad())

	establishOrder()
}
