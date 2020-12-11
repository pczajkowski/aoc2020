package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	count int64
	name  string
}

func getBags(line string, bags map[string][]Bag) {
	bugNameRegex := regexp.MustCompile(`(\w+\s\w+) bags contain `)
	result := bugNameRegex.FindStringSubmatch(line)
	if len(result) != 2 {
		return
	}

	bagName := result[1]

	bugContainsRegex := regexp.MustCompile(`(\d+)\s(\w+\s\w+)`)
	subStrings := strings.Split(line, ", ")
	for _, sub := range subStrings {
		subResult := bugContainsRegex.FindStringSubmatch(sub)
		if len(subResult) != 3 {
			continue
		}

		number, err := strconv.ParseInt(subResult[1], 10, 32)
		if err != nil {
			continue
		}

		bags[bagName] = append(bags[bagName], Bag{name: subResult[2], count: number})
	}
}

func readFile(filePath string, bags map[string][]Bag) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		getBags(line, bags)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func checkBag(phrase string, bagName string, bags map[string][]Bag) bool {
	for _, bag := range bags[bagName] {
		if bag.name == phrase {
			return true
		}

		if checkBag(phrase, bag.name, bags) {
			return true
		}
	}

	return false
}

func countBagsContaining(phrase string, bags map[string][]Bag) int {
	count := 0

	for name, _ := range bags {
		if name == phrase {
			continue
		}

		if checkBag(phrase, name, bags) {
			count++
		}
	}

	return count
}

func countBagsIn(bag string, bags map[string][]Bag) int64 {
	var count int64 = 0

	for _, value := range bags[bag] {
		count += value.count

		count += value.count * countBagsIn(value.name, bags)
	}

	return count
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	var bags = make(map[string][]Bag)

	readFile(os.Args[1], bags)
	fmt.Println("Part1:", countBagsContaining("shiny gold", bags))
	fmt.Println("Part2:", countBagsIn("shiny gold", bags))
}
