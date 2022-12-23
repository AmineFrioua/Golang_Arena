package main

// Remove duplicate letter from a string
import "fmt"

func main() {
	input := "Hello, world!"
	output := ""

	seen := make(map[rune]bool)
	for _, c := range input {
		if !seen[c] {
			seen[c] = true
			output += string(c)
		}
	}

	fmt.Println(output) 
}