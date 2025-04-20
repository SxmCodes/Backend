package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b==0{
		return 0, errors.New("Division by zero is not allowed")
		}
		return a / b, nil
	default:
		return 0, errors.New("Invalid operator")
	}
}

func main(){
	// Creating a simple command line calculator. 
	reader := bufio.NewReader(os.Stdin) // Gets the input from console.
	fmt.Println("Enter the first number: ")
	// we took the input as string and then converted it to integer. 
	// input, _ := reader.ReadString('\n') is a line of code that reads a string from a bufio.Reader until a newline character (\n) is encountered. 
	inputA, _ := reader.ReadString('\n') // Will read string until a new line has been encountered.
	inputA = strings.TrimSpace(inputA) // Removes the leading and trailing white spaces from the string.
	
	a, err := strconv.Atoi(inputA) // Converts the string to integer.
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	// Taking operator as an input. 
	fmt.Println("Enter the operator: ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	fmt.Println("Enter the second number: ")
	inputB, _ := reader.ReadString('\n')
	inputB = strings.TrimSpace(inputB)
	b, err := strconv.Atoi(inputB)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	// Calling the calc function.
	result, err := calc(a, b, op) // Calls the calc function with the input values.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %d\n", result)
}