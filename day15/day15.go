package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type numberSpoken struct {
	number int
	rounds []int
}

func readFile(filePath string) []numberSpoken {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var numbersSpoken []numberSpoken
	for i, item := range strings.Split(string(content), ",") {
		var number int
		n, err := fmt.Sscanf(item, "%d", &number)
		if err != nil || n < 1 {
			log.Fatal(err)
		}

		numbersSpoken = append(numbersSpoken, numberSpoken{number: number, rounds: []int{i + 1}})
	}

	return numbersSpoken
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	numbersSpoken := readFile(os.Args[1])
	fmt.Println(numbersSpoken)
}
