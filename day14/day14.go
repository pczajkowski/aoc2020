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

func sum() int64 {
	var sum int64
	for _, value := range mem {
		sum += value
	}

	return sum
}

var masks []string

func permuteMask(index int64, masksSoFar []string) []string {
	if index < 0 {
		return masksSoFar
	}

	var newMasks []string
	for _, mask := range masksSoFar {
		if mask[index] != 'X' {
			newMasks = append(newMasks, mask)
			continue
		}

		newMask1 := []byte(mask)
		newMask1[index] = '1'
		newMasks = append(newMasks, string(newMask1))

		newMask2 := []byte(mask)
		newMask2[index] = '0'
		newMasks = append(newMasks, string(newMask2))
	}

	return permuteMask(index-1, newMasks)
}

func processLine2(line string) error {
	if strings.Contains(line, "mask") {
		var currentMask string
		n, err := fmt.Sscanf(line, "mask = %s\n", &currentMask)
		if err != nil || n != 1 {
			return fmt.Errorf("Error scanning '%s': %s", line, err)
		}

		masks = permuteMask(maskMaxIndex, []string{currentMask})
		fmt.Println(masks)

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

func readFile2(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if err := processLine2(line); err != nil {
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

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)

	}

	readFile(file)
	fmt.Println("Part1:", sum())
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	file, err = os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)

	}

	readFile2(file)

	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}
}
