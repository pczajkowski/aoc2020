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

	for _, ing := range strings.Split(parts[0], " ") {
		food.ingredients = append(food.ingredients, ing)
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

type alerg struct {
	count       int
	ingredients map[string]int
}

func processFoods(foods []dish) map[string]alerg {
	allergensPossibleForIngredients := make(map[string]alerg)

	for _, food := range foods {
		for _, allergen := range food.allergens {
			var currentAllergen alerg
			if _, ok := allergensPossibleForIngredients[allergen]; !ok {
				currentAllergen = alerg{count: 0, ingredients: make(map[string]int)}
			} else {
				currentAllergen = allergensPossibleForIngredients[allergen]
			}

			currentAllergen.count++
			for _, ingredient := range food.ingredients {
				currentAllergen.ingredients[ingredient]++
			}

			allergensPossibleForIngredients[allergen] = currentAllergen
		}
	}

	return allergensPossibleForIngredients
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

	allergens := processFoods(foods)
	fmt.Println(allergens)
}
