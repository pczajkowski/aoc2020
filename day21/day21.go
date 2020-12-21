package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dish struct {
	ingredients []string
	allergens   []string
}

func readDish(line string) dish {
	var food dish
	parts := strings.Split(line, " (contains ")
	if len(parts) != 2 {
		log.Fatalf("Invalid line: %s", line)
	}

	for _, ingredient := range strings.Split(parts[0], " ") {
		food.ingredients = append(food.ingredients, ingredient)
	}

	cleanedPart2 := strings.TrimSuffix(parts[1], ")")
	for _, allergen := range strings.Split(cleanedPart2, ", ") {
		food.allergens = append(food.allergens, allergen)
	}

	return food
}

func readFile(file *os.File) []dish {
	var foods []dish
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		foods = append(foods, readDish(line))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	return foods
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

	foods := readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	fmt.Println(foods)
}
