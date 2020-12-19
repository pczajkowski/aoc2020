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
	value   string
	mapping [][]int
}

var rules map[int]rule

func init() {
	rules = make(map[int]rule)
}

func readRule(line string) {
	var newRule rule
	parts := strings.Split(line, ": ")
	if len(parts) != 2 {
		log.Fatalf("Invalid line: %s", line)
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Error processing rule id for %s: %s", line, err)
	}

	var currentMapping []int
	for _, item := range strings.Split(parts[1], " ") {
		if strings.Contains(item, "\"") {
			newRule.value = strings.ReplaceAll(item, "\"", "")
			break
		}

		if item == "|" {
			newRule.mapping = append(newRule.mapping, currentMapping)
			currentMapping = []int{}
			continue
		}

		itemID, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("Error processing id for %s: %s", item, err)
		}

		currentMapping = append(currentMapping, itemID)
	}
	newRule.mapping = append(newRule.mapping, currentMapping)

	rules[id] = newRule
}

var messages []string

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	currentFunction := readRule
	changed := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if changed {
				break
			}

			currentFunction = func(line string) { messages = append(messages, line) }
			changed = true
			continue
		}

		currentFunction(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
}

func buildStringsForRule(ruleID int, stringsSoFar []string) []string {
	var newStrings []string

	if rules[ruleID].value != "" {
		for _, item := range stringsSoFar {
			item += rules[ruleID].value
			newStrings = append(newStrings, item)
		}
		return newStrings
	}

	for _, item := range stringsSoFar {
		for _, array := range rules[ruleID].mapping {
			partial := []string{item}
			for _, id := range array {
				partial = buildStringsForRule(id, partial)
			}
			newStrings = append(newStrings, partial...)
		}
	}
	return newStrings
}

var notMatched []string

func part1(validStrings []string) int {
	valid := 0

	for _, message := range messages {
		matched := false
		for _, item := range validStrings {
			if message == item {
				valid++
				matched = true
				break
			}
		}
		if !matched {
			notMatched = append(notMatched, message)
		}
	}

	return valid
}

func startsWith(toCheck []string, message string) bool {
	for _, item := range toCheck {
		if strings.HasPrefix(message, item) {
			return true
		}
	}

	return false
}

func endsWith(toCheck []string, message string) bool {
	for _, item := range toCheck {
		if strings.HasSuffix(message, item) {
			return true
		}
	}

	return false
}

func part2(max int) int {
	valid := 0
	thirtyOne := buildStringsForRule(31, []string{""})
	fortyTwo := buildStringsForRule(42, []string{""})
	length31 := len(thirtyOne[0])
	length42 := len(fortyTwo[0])

	currentLength := length31 + 2*length42
	for len(notMatched) > 0 && currentLength <= max {
		currentLength += length31
		var toCheck []string

		for _, message := range notMatched {
			if len(message)%currentLength == 0 {
				if startsWith(fortyTwo, message) && endsWith(thirtyOne, message) {
					valid++
				} else {
					toCheck = append(toCheck, message)
				}
			} else {
				toCheck = append(toCheck, message)
			}
		}

		notMatched = toCheck
	}

	return valid
}

func longest() int {
	max := 0
	for _, message := range notMatched {
		length := len(message)
		if length > max {
			max = length
		}
	}

	return max
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

	validStrings := buildStringsForRule(0, []string{""})
	resultPart1 := part1(validStrings)
	fmt.Println("Part1:", resultPart1)

	max := longest()
	fmt.Println("Part2:", resultPart1+part2(max))
}
