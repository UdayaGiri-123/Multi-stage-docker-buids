package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to perform addition
func add(a, b float64) float64 {
	return a + b
}

// Function to perform subtraction
func subtract(a, b float64) float64 {
	return a - b
}

// Function to perform multiplication
func multiply(a, b float64) float64 {
	return a * b
}

// Function to perform division
func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Calculator")
	fmt.Println("------------------")

	for {
		// Read the first number
		fmt.Print("Enter the first number: ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		// Read the operator
		fmt.Print("Enter the operator (+, -, *, /): ")
		operator, _ := reader.ReadString('\n')
		operator = strings.TrimSpace(operator)

		// Read the second number
		fmt.Print("Enter the second number: ")
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)
		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		// Perform the calculation
		var result float64
		switch operator {
		case "+":
			result = add(num1, num2)
		case "-":
			result = subtract(num1, num2)
		case "*":
			result = multiply(num1, num2)
		case "/":
			result = divide(num1, num2)
		default:
			fmt.Println("Invalid operator. Please use one of +, -, *, /.")
			continue
		}

		// Display the result
		fmt.Printf("Result: %.2f\n", result)

		// Ask if the user wants to perform another calculation
		fmt.Print("Do you want to perform another calculation? (yes/no): ")
		again, _ := reader.ReadString('\n')
		again = strings.TrimSpace(again)
		if strings.ToLower(again) != "yes" {
			break
		}
	}
	fmt.Println("Thank you for using the calculator!")
}
