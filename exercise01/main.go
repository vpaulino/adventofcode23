package main

import (
	"fmt"
	"regexp"
	"strconv"
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

		fmt.Printf(" { \"%s\", %d - %d } \n", line, intNums, total)
	}
}

func ParseAndConvertToInt(text string) (int, error) {

	numbersStr := parseInts(text)

	intNums, err := convertToIntegers(numbersStr)

	return intNums, err

}

func ParseAndConvertToTextInt(text string) (int, error) {

	numberOnString := parseIntsFromMap(text)

	intNums, err := convertToIntegers(numberOnString)

	return intNums, err
}

func parseIntsFromMap(textToParse string) string {

	numbersMap := make(map[string]int)

	numbersMap["zero"] = 0
	numbersMap["one"] = 1
	numbersMap["two"] = 2
	numbersMap["three"] = 3
	numbersMap["four"] = 4
	numbersMap["five"] = 5
	numbersMap["six"] = 6
	numbersMap["seven"] = 7
	numbersMap["eight"] = 8
	numbersMap["nine"] = 9
	numbersMap["0"] = 0
	numbersMap["1"] = 1
	numbersMap["2"] = 2
	numbersMap["3"] = 3
	numbersMap["4"] = 4
	numbersMap["5"] = 5
	numbersMap["6"] = 6
	numbersMap["7"] = 7
	numbersMap["8"] = 8
	numbersMap["9"] = 9

	regularExpression := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|\d)`)

	matches := regularExpression.FindAllString(textToParse, -1)

	firstNumber := matches[0]
	lastNumber := matches[len(matches)-1]

	numberOnString := fmt.Sprintf("%d%d", numbersMap[firstNumber], numbersMap[lastNumber])

	return numberOnString
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
