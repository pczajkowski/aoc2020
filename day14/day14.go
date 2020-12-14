package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var mem map[int64]int64
var mask string

const (
	maskMaxIndex = 35
)

func setBits(number int64, mask string) int64 {
	for i := 0; i <= maskMaxIndex; i++ {
		switch mask[maskMaxIndex-i] {
		case 'X':
			continue
		case '1':
			number |= (1 << i)
		case '0':
			var tempMask int64
			tempMask = ^(1 << i)
			number &= tempMask
		}
	}

	return number
}

func processLine(line string) error {
	if strings.Contains(line, "mask") {
		n, err := fmt.Sscanf(line, "mask = %s\n", &mask)
		if err != nil || n != 1 {
			return fmt.Errorf("Error scanning '%s': %s", line, err)
		}

		return nil
	}

	var id int64
	var number int64
	n, err := fmt.Sscanf(line, "mem[%d] = %d", &id, &number)
	if err != nil || n != 2 {
		return fmt.Errorf("Error scanning '%s': %s", line, err)
	}

	number = setBits(number, mask)
	mem[id] = number

	return nil
}

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if err := processLine(line); err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
}

func init() {
	mem = make(map[int64]int64)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open %s!\n", os.Args[1])

	}

	readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	fmt.Println(mem)
}
