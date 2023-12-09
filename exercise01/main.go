package main

import (
	"fmt"
)

func main() {

	filename := "Input.txt" // Replace with your file path

	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0
	for _, line := range lines {

		//intNums, err := ParseAndConvertToInt(line)
		intNums, err := ParseAndConvertToTextInt(line)
		total += intNums

		if err != nil {
			fmt.Println("error converting ints", err)
		}

		fmt.Printf(" { \"%s\", %d }, \n", line, intNums)
	}

	fmt.Printf(" Total is: %d \n", total)
}
