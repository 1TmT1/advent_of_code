package main

import (
	"adventOfCode/helpers"
	"fmt"
	"strconv"
)

func getFirstDigitFirstHalf(line string) *int {
	for _, value := range line {
		digit, err := strconv.Atoi(string(value))
		if err == nil {
			return &digit
		}
	}

	return nil
}

func getLastDigitFirstHalf(line string) *int {
	for i := len(line) - 1; i > -1; i-- {
		digit, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return &digit
		}
	}

	return nil
}

func twoNumToDecimal(first, second int) int {
	return first*10 + second
}

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

func getFirstDigitSecondHalf(line string, numbers []string) *int {
	var numberString string
	var possibilitiesIndexes []int

	for i, value := range line {
		digit, err := strconv.Atoi(string(value))
		if err == nil {
			return &digit
		}

		if possibilitiesIndexes == nil {
			for index, number := range numbers {
				if number[len(numberString)] == byte(value) {
					numberString += string(value)
					possibilitiesIndexes = append(possibilitiesIndexes, index)
				}
			}
		} else {
			found := false
			for _, indexNumber := range possibilitiesIndexes {
				if numbers[indexNumber][len(numberString)] == byte(value) {
					numberString += string(value)
					if numberString == numbers[indexNumber] {
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

					possibilitiesIndexes = []int{indexNumber}
					found = true
					break
				}
			}

			if !found {
				numberString = ""
				possibilitiesIndexes = nil
				i--
			}
		}
	}
	return nil
}

func getLastDigitSecondHalf(line string, numbers []string) *int {
	var numberString string
	var possibilitiesIndexes []int

	for i := len(line) - 1; i > -1; i-- {
		digit, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return &digit
		}

		if possibilitiesIndexes == nil {
			for index, number := range numbers {
				if number[len(number)-len(numberString)] == byte(line[i]) {
					numberString += string(line[i])
					possibilitiesIndexes = append(possibilitiesIndexes, index)
				}
			}
		} else {
			found := false
			for _, indexNumber := range possibilitiesIndexes {
				if numbers[indexNumber][len(numbers[indexNumber])-len(numberString)] == byte(line[i]) {
					numberString += string(line[i])
					if numberString == numbers[indexNumber] {
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

					possibilitiesIndexes = []int{indexNumber}
					found = true
					break
				}
			}

			if !found {
				numberString = ""
				possibilitiesIndexes = nil
				i++
			}
		}
	}
	return nil
}

func main() {
	fileInput := helpers.GetInput("./inputl.txt")

	fmt.Println(solveFirstHalf(fileInput))

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range fileInput {
		fmt.Println(*getFirstDigitSecondHalf(line, numbers))
		fmt.Println(*getLastDigitSecondHalf(line, numbers))
	}
}
