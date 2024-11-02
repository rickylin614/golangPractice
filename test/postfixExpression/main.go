package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	expression := "( 3 + 4 ) * 5 - 6 ^ ( 2 / 2 )"
	postfix := infixToPostfix(expression)
	fmt.Println("Postfix:", postfix)
	result := evaluatePostfix(postfix)
	fmt.Println("Result:", result)
}

func getPrecedence(op string) int {
	switch op {
	case "^":
		return 3
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	}
	return 0
}

func infixToPostfix(expression string) []string {
	var postfix []string
	var stack []string

	for _, token := range strings.Fields(expression) {
		switch token {
		case "(":
			stack = append(stack, token)
		case "^":
			for len(stack) > 0 && getPrecedence(stack[len(stack)-1]) > getPrecedence(token) {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case "+", "-", "*", "/":
			for len(stack) > 0 && getPrecedence(stack[len(stack)-1]) >= getPrecedence(token) {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		case ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Pop the "("
			}
		default:
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func evaluatePostfix(postfix []string) int {
	var stack []int

	for _, token := range postfix {
		value, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, value)
		} else {
			if len(stack) < 2 {
				fmt.Println("Error: Invalid expression")
				return 0
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			case "^":
				stack = append(stack, int(math.Pow(float64(a), float64(b))))
			}
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error: Invalid expression")
		return 0
	}
	return stack[0]
}
