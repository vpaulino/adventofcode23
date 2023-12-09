package main

import (
	"fmt"
	"regexp"
	"strconv"
)

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

	firstNumber := FindFirstNumberOccorrence(textToParse, numbersMap)
	lastNumber := FindLastNumberOccorrence(textToParse, numbersMap)

	numberOnString := fmt.Sprintf("%d%d", numbersMap[firstNumber], numbersMap[lastNumber])

	return numberOnString
}

func FindFirstNumberOccorrence(textToParse string, wordsToMatch map[string]int) string {
	var startIndex int = 0

	for startIndex <= len(textToParse)-1 {

		digitFound := textToParse[startIndex]
		foundNumber, found := wordsToMatch[string(digitFound)]

		if found {
			return fmt.Sprintf("%d", foundNumber)
		}

		var textsParsed = []string{}

		if startIndex+3 <= len(textToParse) {

			var threeWordsNumber = string(textToParse[startIndex : startIndex+3])
			textsParsed = append(textsParsed, threeWordsNumber)
		}

		if startIndex+4 <= len(textToParse) {
			var fourWordsNumber = string(textToParse[startIndex : startIndex+4])
			textsParsed = append(textsParsed, fourWordsNumber)
		}

		if startIndex+5 <= len(textToParse) {
			var fiveWordsNumber = string(textToParse[startIndex : startIndex+5])
			textsParsed = append(textsParsed, fiveWordsNumber)
		}

		for _, textParsed := range textsParsed {

			foundNumber, found := wordsToMatch[textParsed]
			if found {
				return fmt.Sprintf("%d", foundNumber)
			}
		}
		startIndex++
	}
	// Create a substring from the last character to the first character
	return ""
}

func FindLastNumberOccorrence(textToParse string, wordsToMatch map[string]int) string {

	var startIndex int = len(textToParse) - 1

	for startIndex >= 0 {

		digitFound := textToParse[startIndex]
		foundNumber, found := wordsToMatch[string(digitFound)]

		if found {
			return fmt.Sprintf("%d", foundNumber)
		}

		var textsParsed = []string{}

		if startIndex-2 >= 0 {

			var threeWordsNumber = string(textToParse[startIndex-2 : startIndex+1])
			textsParsed = append(textsParsed, threeWordsNumber)
		}

		if startIndex-3 >= 0 {
			var fourWordsNumber = string(textToParse[startIndex-3 : startIndex+1])
			textsParsed = append(textsParsed, fourWordsNumber)
		}

		if startIndex-4 >= 0 {
			var fiveWordsNumber = string(textToParse[startIndex-4 : startIndex+1])
			textsParsed = append(textsParsed, fiveWordsNumber)
		}

		if startIndex-5 >= 0 {
			var fiveWordsNumber = string(textToParse[startIndex-5 : startIndex+1])
			textsParsed = append(textsParsed, fiveWordsNumber)
		}

		for _, textParsed := range textsParsed {

			foundNumber, found := wordsToMatch[textParsed]
			if found {
				return fmt.Sprintf("%d", foundNumber)
			}
		}
		startIndex--
	}
	// Create a substring from the last character to the first character
	return ""
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
