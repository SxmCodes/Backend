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
	input, _ := reader.ReadString('\n') // Will read string until a new line has been encountered.
	input = strings.TrimSpace(input) // Removes the leading and trailing white spaces from the string.

	// splitting the parts into 3, because we need 2 inputs and 1 string (to tell which action to perform). 
	parts := strings.Split(input, " ") // Splits the string into parts based on the space.
	if len(parts) != 3 {
		fmt.Println("Invalid input")
		return
	}

	a, err := strconv.Atoi(parts[0]) // Converts the string to integer.
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	
	b, err := strconv.Atoi(parts[0]) // Converts the string to integer for parts[2]
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	result, err := calc(a, b, parts[1]) // Calls the calc function with the input values.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %d\n", result)
}
