package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	filename := "Input.txt" // Replace with your file path

	lines, err := readLines(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0
	for _, line := range lines {

		intNums, err := ParseAndConvertToInt(line)
		total += intNums
		if err != nil {
			fmt.Println("error converting ints", err)
		}
		fmt.Printf(" %s - %d : Total: %d ", line, intNums, total)
	}

}

func readLines(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ParseAndConvertToInt(text string) (int, error) {

	numbersStr := parseInts(text)

	intNums, err := convertToIntegers(numbersStr)

	return intNums, err

}

func parseInts(textToParse string) string {

	regularExpression := regexp.MustCompile(`\d`)

	matches := regularExpression.FindAllString(textToParse, -1)

	firstNumber := matches[0]
	lastNumber := matches[len(matches)-1]

	return firstNumber + lastNumber

}

func convertToIntegers(numbers string) (int, error) {

	intNumbers, err := strconv.Atoi(numbers)

	return intNumbers, err
}
