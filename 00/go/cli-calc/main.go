package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: <num1> <operation> <num2>\nExample: 1 + 2")
	}

	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return
	}
	b, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return
	}
	op := args[1]

	switch op {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(substract(a, b))
	case "*":
		fmt.Println(multiply(a, b))
	case "/":
		fmt.Println(divide(a, b))
	}

}

func add(a, b float64) float64 {
	return a + b
}

func substract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("divide by zero")
		return -1
	}
	return a / b
}
