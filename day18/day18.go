package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getTokens(line string) ([]rune, error) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	var tokens []rune
	for scanner.Scan() {
		newTokens := []rune(scanner.Text())
		tokens = append(tokens, newTokens...)
	}
	if err := scanner.Err(); err != nil {
		return tokens, fmt.Errorf("Scanner error: %s", err)
	}

	return tokens, nil
}

func getExpression(tokens []rune) []interface{} {
	var expression []interface{}
	for _, token := range tokens {
		stringToken := string(token)
		value, err := strconv.Atoi(stringToken)
		if err != nil {
			expression = append(expression, stringToken)
			continue
		}

		expression = append(expression, value)
	}

	return expression
}

func readFile(file *os.File) [][]interface{} {
	var expressions [][]interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tokens, err := getTokens(line)
		if err != nil {
			log.Fatalf("Error scanning %s: %s", line, err)
		}

		expressions = append(expressions, getExpression(tokens))

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	return expressions
}

func getRPNFromExpression(expression []interface{}) []interface{} {
	var rpn []interface{}
	var operators []interface{}

	for _, token := range expression {
		switch token := token.(type) {
		case int:
			rpn = append(rpn, token)
		case string:
			switch token {
			case "(":
				operators = append(operators, token)
			case ")":
				for len(operators) > 0 {
					oper := operators[len(operators)-1]
					operators = operators[:len(operators)-1]

					if oper == "(" {
						break
					}

					rpn = append(rpn, oper)
				}
			default:
				for len(operators) > 0 {
					top := operators[len(operators)-1]

					if top == "(" {
						break
					}

					operators = operators[:len(operators)-1]
					rpn = append(rpn, top)
				}

				operators = append(operators, token)
			}

		}
	}
	for len(operators) > 0 {
		oper := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		rpn = append(rpn, oper)
	}

	return rpn
}

func doMath(operator string, arg1, arg2 int) int {
	switch operator {
	case "+":
		return arg1 + arg2
	case "*":
		return arg1 * arg2
	}

	return -1
}

func evaluateRPN(rpn []interface{}) int {
	var stack []int

	for _, token := range rpn {
		switch token := token.(type) {
		case int:
			stack = append(stack, token)
		case string:
			if len(stack) < 2 {
				log.Fatalf("Invalid expresion token %s in %s!", token, rpn)
			}

			arg1, arg2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			value := doMath(token, arg1, arg2)
			stack = append(stack, value)
		}
	}
	if len(stack) != 1 {
		log.Fatal("Bad stack!")
	}

	return stack[len(stack)-1]
}

func part1(expressions [][]interface{}) int {
	sum := 0

	for _, expression := range expressions {
		rpn := getRPNFromExpression(expression)
		sum += evaluateRPN(rpn)
	}

	return sum
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

	expressions := readFile(file)
	if err := file.Close(); err != nil {
		log.Fatalf("Failed to close file: %s", err)
	}

	fmt.Println("Part1:", part1(expressions))
}
