package main 

import (
	"fmt"
	"errors"
	"os"
	"strconv"
	"strings"
	"bufio"
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

}

// Learnings :- 
	// 1. errors.New and Errors lib
