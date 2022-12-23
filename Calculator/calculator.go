package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an equation (or 'q' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		tokens := strings.Split(input, " ")
		if len(tokens) != 3 {
			fmt.Println("Invalid input. Please enter a valid equation in the format 'x operator y'.")
			continue
		}

		x, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid equation in the format 'x operator y'.")
			continue
		}

		y, err := strconv.ParseFloat(tokens[2], 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid equation in the format 'x operator y'.")
			continue
		}

		var result float64
		isCorrect :=true

		switch tokens[1] {
		case "+":
			result = x + y
		case "-":
			result = x - y
		case "*":
			result = x * y
		case "/":
			if y==0{
			isCorrect = false
			break
			}
			result = x / y
		default:
			fmt.Println("Invalid operator. Please use one of the following operators: +, -, *, /.")
			continue
		}

		if isCorrect== true{
		fmt.Printf("Result: %.2f\n", result)
		}else{
		fmt.Println("Shame on you for dividing by 0, Shame. Ring the bells")
		}
	}
}
