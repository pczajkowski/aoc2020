package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var rounds map[int]int
var numbersSpoken map[int][]int

func readFile(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for i, item := range strings.Split(string(content), ",") {
		var number int
		n, err := fmt.Sscanf(item, "%d", &number)
		if err != nil || n < 1 {
			log.Fatal(err)
		}

		rounds[i+1] = number
		numbersSpoken[number] = []int{i + 1}
	}
}

func init() {
	rounds = make(map[int]int)
	numbersSpoken = make(map[int][]int)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	readFile(os.Args[1])
	fmt.Println(numbersSpoken, rounds)
}
