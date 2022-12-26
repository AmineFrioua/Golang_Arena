package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("operation.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read the first line of the file
	// Read all of the lines in the file
	for scanner.Scan() {
		operation := scanner.Text()

	// Split the operation into its component parts
	parts := strings.Split(operation, " ")

	// Check the operation type
	if parts[0] == "add" {
		// Parse the operands
		operand1, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		operand2, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Perform the operation
		result := operand1 + operand2

		// Print the result
		fmt.Printf("%d + %d = %d\n", operand1, operand2, result)
	} else if parts[0] == "subtract" {
		// Parse the operands
		operand1, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		operand2, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Perform the operation
		result := operand1 - operand2

		// Print the result
		fmt.Printf("%d - %d = %d\n", operand1, operand2, result)
	}else if parts[0] == "multiply" {
		// Parse the operands
		operand1, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		operand2, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Perform the operation
		result := operand1 * operand2

		// Print the result
		fmt.Printf("%d * %d = %d\n", operand1, operand2, result)
	} else if parts[0] == "divide" {
		// Parse the operands
		operand1, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		operand2, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		if (operand2==0 ){
			fmt.Println("Can't divide by 0")
		} else{
	
		// Perform the operation
		result := operand1 / operand2

		// Print the result
		fmt.Printf("%d / %d = %d\n", operand1, operand2, result)}
	}else {
		fmt.Println("Unknown operation")
	}
}
}
