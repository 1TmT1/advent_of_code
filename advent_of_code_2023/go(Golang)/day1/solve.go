package main

// Imports
import (
	"adventOfCode/helpers"
	"fmt"
	"strconv"
)

// Get first digit of the string - only numeric
func getFirstDigitFirstHalf(line string) *int {
	for _, value := range line {
		digit, err := strconv.Atoi(string(value))
		if err == nil {
			return &digit
		}
	}

	return nil
}

// Get last digit of the string - only numeric
func getLastDigitFirstHalf(line string) *int {
	for i := len(line) - 1; i > -1; i-- {
		digit, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return &digit
		}
	}

	return nil
}

// Take two integers (ones, tens) and return the proper integer combined (5, 7 => 57)
func twoNumToDecimal(first, second int) int {
	return first*10 + second
}

// Solve first part of day 1
func solveFirstHalf(fileInput []string) int {
	sum := 0

	for _, line := range fileInput {
		firstDigit := *getFirstDigitFirstHalf(line)
		lastDigit := *getLastDigitFirstHalf(line)
		num := twoNumToDecimal(firstDigit, lastDigit)
		sum += num
	}

	return sum
}

// Get first digit doesn't matter if numeric or text (5, 'five')
func getFirstDigitSecondHalf(line string, numbers []string) *int {
	var possibilitiesIndexes []int

	for i, value := range line {
		digit, err := strconv.Atoi(string(value))
		if err == nil {
			return &digit
		}

		if possibilitiesIndexes == nil {
			for index, number := range numbers {
				if number[0] == byte(value) {
					possibilitiesIndexes = append(possibilitiesIndexes, index)
				}
			}
		}
		for posIndex := 0; posIndex < len(possibilitiesIndexes); posIndex++ {
			numberString := string(line[i])
			for letterIndex := 0; letterIndex < len(numbers[possibilitiesIndexes[posIndex]]); letterIndex++ {
				if line[i+len(numberString)] == numbers[possibilitiesIndexes[posIndex]][len(numberString)] {
					numberString += string(line[i+len(numberString)])
					if numberString == numbers[possibilitiesIndexes[posIndex]] {
						var numToReturn int
						switch numberString {
						case "one":
							numToReturn = 1
						case "two":
							numToReturn = 2
						case "three":
							numToReturn = 3
						case "four":
							numToReturn = 4
						case "five":
							numToReturn = 5
						case "six":
							numToReturn = 6
						case "seven":
							numToReturn = 7
						case "eight":
							numToReturn = 8
						case "nine":
							numToReturn = 9
						}

						return &numToReturn
					}
				} else {
					break
				}
			}
		}
		possibilitiesIndexes = nil
	}
	return nil
}

// Get last digit doesn't matter if numeric or text (5, 'five')
func getLastDigitSecondHalf(line string, numbers []string) *int {
	var possibilitiesIndexes []int

	for i := len(line) - 1; i > -1; i-- {
		digit, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return &digit
		}

		if possibilitiesIndexes == nil {
			for index, number := range numbers {
				if number[len(number)-1] == byte(line[i]) {
					possibilitiesIndexes = append(possibilitiesIndexes, index)
				}
			}
		}
		for posIndex := 0; posIndex < len(possibilitiesIndexes); posIndex++ {
			numberString := string(line[i])
			for letterIndex := len(numbers[possibilitiesIndexes[posIndex]]); letterIndex > -1; letterIndex-- {
				if line[i-len(numberString)] == numbers[possibilitiesIndexes[posIndex]][len(numbers[possibilitiesIndexes[posIndex]])-len(numberString)-1] {
					numberString = string(line[i-len(numberString)]) + numberString
					if numberString == numbers[possibilitiesIndexes[posIndex]] {
						var numToReturn int
						switch numberString {
						case "one":
							numToReturn = 1
						case "two":
							numToReturn = 2
						case "three":
							numToReturn = 3
						case "four":
							numToReturn = 4
						case "five":
							numToReturn = 5
						case "six":
							numToReturn = 6
						case "seven":
							numToReturn = 7
						case "eight":
							numToReturn = 8
						case "nine":
							numToReturn = 9
						}

						return &numToReturn
					}
				} else {
					break
				}
			}
		}
		possibilitiesIndexes = nil
	}
	return nil
}

// Solve second part of day 1
func solveSecondHalf(fileInput, numbers []string) int {
	sum := 0

	for _, line := range fileInput {
		firstDigit := *getFirstDigitSecondHalf(line, numbers)
		lastDigit := *getLastDigitSecondHalf(line, numbers)
		num := twoNumToDecimal(firstDigit, lastDigit)
		sum += num
	}

	return sum
}

// Main function
func main() {
	fileInput := helpers.GetInput("./input.txt")

	fmt.Println("First half answer: ", solveFirstHalf(fileInput))

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	fmt.Println("Second half answer: ", solveSecondHalf(fileInput, numbers))
}
