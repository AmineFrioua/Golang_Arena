package main

// a script that reverse an integer if inputted correctly
import (
	"fmt"
	"strconv"
)

func reverseInt(input interface{}) interface{} {
	switch v := input.(type) {
	case int:
		inputStr := strconv.Itoa(v)
		outputStr := ""
		for i := len(inputStr) - 1; i >= 0; i-- {
			outputStr += string(inputStr[i])
		}
		output, _ := strconv.Atoi(outputStr)
		return output
	default:
		return "not an integer"
	}
}

func main() {
	fmt.Println(reverseInt(12345)) // prints "54321"
	fmt.Println(reverseInt("hello")) // prints "not an integer"
}