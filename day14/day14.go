package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func processLine(id int64, number int64) {
	number = setBits(number, mask)
	mem[id] = number
}

func readFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if strings.Contains(line, "mask") {
			n, err := fmt.Sscanf(line, "mask = %s\n", &mask)
			if err != nil || n != 1 {
				log.Fatalf("Error scanning '%s': %s", line, err)
			}

			continue
		}

		var id int64
		var number int64
		n, err := fmt.Sscanf(line, "mem[%d] = %d", &id, &number)
		if err != nil || n != 2 {
			log.Fatalf("Error scanning '%s': %s", line, err)
		}

		processLine(id, number)
		processLine2(id, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}
}

func sum(memory map[int64]int64) int64 {
	var sum int64
	for _, value := range memory {
		sum += value
	}

	return sum
}

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

func setBitsString(number string) string {
	newNumber := []byte(number)
	for i := 0; i <= maskMaxIndex; i++ {
		switch mask[i] {
		case 'X':
			newNumber[i] = 'X'
		case '1':
			newNumber[i] = '1'
		case '0':
			continue
		}
	}

	return string(newNumber)
}

var mem2 map[int64]int64

func processLine2(id int64, number int64) error {
	numberString := fmt.Sprintf("%036b", id)
	result := setBitsString(numberString)
	masks := permuteMask(maskMaxIndex, []string{result})
	for _, currentMask := range masks {
		currentID, err := strconv.ParseInt(currentMask, 2, 64)
		if err != nil {
			return fmt.Errorf("Error parsing current ID %s: %s", currentMask, err)
		}

		mem2[currentID] = number
	}

	return nil
}

func init() {
	mem = make(map[int64]int64)
	mem2 = make(map[int64]int64)
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

	fmt.Println("Part1:", sum(mem))
	fmt.Println("Part2:", sum(mem2))
}
