package main

///This program reads each line of the file
// using a loop and a scanner, and uses a regular expression to search for common typos.
// It then uses the regexp.ReplaceAllString function to fix the typos and prints the fixed line to the console.

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the file
	file, err := os.Open("file_to_fix.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// Define a regular expression to match common typos
	re := regexp.MustCompile(`(?i)(teh|thier|their|thats|untill|until)`)

	// Read all of the lines in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Fix any typos found in the line
		fixedLine := re.ReplaceAllString(line, "the")

		// Print the fixed line
		fmt.Println(fixedLine)
	}
}
