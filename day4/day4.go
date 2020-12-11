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

func checkByr(line string) int {
	if strings.Contains(line, "byr:") {
		re := regexp.MustCompile(`byr:(\d{4})`)
		result := re.FindStringSubmatch(line)
		if len(result) != 2 {
			return 0
		}

		byrString := result[1]
		if byrString == "" {
			return 0
		}

		byr, err := strconv.ParseInt(byrString, 10, 32)
		if err != nil {
			return 0
		}

		if byr >= 1920 && byr <= 2002 {
			return 1
		}
	}
	return 0
}

func checkIyr(line string) int {
	if strings.Contains(line, "iyr:") {
		re := regexp.MustCompile(`iyr:(\d{4})`)
		result := re.FindStringSubmatch(line)
		if len(result) != 2 {
			return 0
		}

		iyrString := result[1]
		if iyrString == "" {
			return 0
		}

		iyr, err := strconv.ParseInt(iyrString, 10, 32)
		if err != nil {
			return 0
		}

		if iyr >= 2010 && iyr <= 2020 {
			return 1
		}
	}
	return 0
}

func checkEyr(line string) int {
	if strings.Contains(line, "eyr:") {
		re := regexp.MustCompile(`eyr:(\d{4})`)
		result := re.FindStringSubmatch(line)
		if len(result) != 2 {
			return 0
		}

		eyrString := result[1]
		if eyrString == "" {
			return 0
		}

		eyr, err := strconv.ParseInt(eyrString, 10, 32)
		if err != nil {
			return 0
		}

		if eyr >= 2020 && eyr <= 2030 {
			return 1
		}
	}
	return 0
}

func checkHgt(line string) int {
	if strings.Contains(line, "hgt:") {
		re := regexp.MustCompile(`hgt:(\d+)(\w{2})`)
		result := re.FindStringSubmatch(line)
		if len(result) != 3 {
			return 0
		}

		hgtString := result[1]
		if hgtString == "" {
			return 0
		}

		tp := result[2]
		if tp == "" {
			return 0
		}

		hgt, err := strconv.ParseInt(hgtString, 10, 32)
		if err != nil {
			return 0
		}

		switch tp {
		case "cm":
			if hgt >= 150 && hgt <= 193 {
				return 1
			}
		case "in":
			if hgt >= 59 && hgt <= 76 {
				return 1
			}
		}
	}

	return 0
}

func checkHcl(line string) int {
	if strings.Contains(line, "hcl:#") {
		re := regexp.MustCompile(`hcl:#([0-9a-f]{6})`)
		result := re.FindStringSubmatch(line)
		if len(result) != 2 {
			return 0
		}

		hcl := result[1]
		if hcl != "" {
			return 1
		}
	}
	return 0
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func checkEcl(line string) int {
	if strings.Contains(line, "ecl:") {
		re := regexp.MustCompile(`ecl:(\w{3})`)
		result := re.FindStringSubmatch(line)
		if len(result) != 2 {
			return 0
		}

		ecl := result[1]
		if ecl == "" {
			return 0
		}

		for _, color := range eyeColors {
			if color == ecl {
				return 1
			}
		}
	}
	return 0
}

func checkPid(line string) int {
	if strings.Contains(line, "pid:") {
		re := regexp.MustCompile(`pid:(\d+)`)
		result := re.FindStringSubmatch(line)
		if len(result) != 2 {
			return 0
		}

		pid := result[1]
		if pid == "" {
			return 0
		}

		if len(pid) != 9 {
			return 0
		}

		_, err := strconv.ParseInt(pid, 10, 32)
		if err == nil {
			return 1
		}
	}
	return 0
}

func getChecks() map[string]func(line string) int {
	toCheck := make(map[string]func(line string) int)
	toCheck["byr:"] = checkByr
	toCheck["iyr:"] = checkIyr
	toCheck["eyr:"] = checkEyr
	toCheck["hgt:"] = checkHgt
	toCheck["hcl:"] = checkHcl
	toCheck["ecl:"] = checkEcl
	toCheck["pid:"] = checkPid

	return toCheck
}

func performCheck(line string, name string, check func(line string) int) (int, int) {
	check1 := 0
	check2 := 0
	if strings.Contains(line, name) {
		check1 = 1
		check2 = check(line)
	}

	return check1, check2
}

func checkLine(line string, toCheck map[string]func(line string) int) (int, int) {
	checks1 := 0
	checks2 := 0

	for name, check := range toCheck {
		check1, check2 := performCheck(line, name, check)
		checks1 += check1
		checks2 += check2
	}

	return checks1, checks2
}

func countChecks(check1 int, check2 int, valid1 *int, valid2 *int) {
	if check1 == 7 {
		*valid1++
	}

	if check2 == 7 {
		*valid2++
	}

}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open")

	}

	var (
		valid1, valid2 int
		text           string
	)

	toCheck := getChecks()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			check1, check2 := checkLine(text, toCheck)
			countChecks(check1, check2, &valid1, &valid2)

			text = ""
			continue
		}

		text += line + " "

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	check1, check2 := checkLine(text, toCheck)
	countChecks(check1, check2, &valid1, &valid2)

	fmt.Println("Part1: ", valid1)
	fmt.Println("Part2: ", valid2)
}
