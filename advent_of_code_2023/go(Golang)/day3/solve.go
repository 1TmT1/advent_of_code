package main

import (
	"adventOfCode/helpers"
	"fmt"
	"strconv"
)

func checkSpecialCharacter(character rune) bool {
	specialCharacters := []rune{'*', '/', '@', '&', '$', '=', '#', '-', '+', '%'}
	for _, value := range specialCharacters {
		if value == character {
			return true
		}
	}

	return false
}

// Helper function to find all special charactes except the dot(.)
func findingAllSpecialChars(inputLines []string) []string {
	var specialCharacters []string
	for _, line := range inputLines {
		for _, character := range line {
			if _, err := strconv.Atoi(string(character)); err != nil {
				if character != '.' {
					alreadyAdded := false
					for _, x := range specialCharacters {
						if x == string(character) {
							alreadyAdded = true

						}
					}
					if !alreadyAdded {
						specialCharacters = append(specialCharacters, string(character))
					}
				}
			}
		}
	}

	return specialCharacters
}

func checkLineBelow(strNumber string, letterIndex, i int, inputLines []string) int {
	//check line below for special characters
	for j := letterIndex - len(strNumber) - 1; j < letterIndex+1; j++ {
		if checkSpecialCharacter(rune(inputLines[i+1][j])) {
			return helpers.AddNumberToSumAfterConvertion(strNumber)
		}
	}

	return 0
}

func checkLineAbove(strNumber string, letterIndex, i int, inputLines []string) int {
	for j := letterIndex - len(strNumber) - 1; j < letterIndex+1; j++ {
		if checkSpecialCharacter(rune(inputLines[i-1][j])) {
			return helpers.AddNumberToSumAfterConvertion(strNumber)
		}
	}

	return 0
}

func checkFirstItemLineAbove(strNumber string, letterIndex, i int, inputLines []string) int {
	for j := letterIndex - len(strNumber); j < letterIndex+1; j++ {
		if checkSpecialCharacter(rune(inputLines[i-1][j])) {
			return helpers.AddNumberToSumAfterConvertion(strNumber)
		}
	}

	return 0
}

func checkFirstItemLineBelow(strNumber string, letterIndex, i int, inputLines []string) int {
	for j := letterIndex - len(strNumber); j < letterIndex+1; j++ {
		if checkSpecialCharacter(rune(inputLines[i+1][j])) {
			return helpers.AddNumberToSumAfterConvertion(strNumber)
		}
	}

	return 0
}

func main() {
	inputLines := helpers.GetInput("./input.txt")

	fmt.Println(findingAllSpecialChars(inputLines))

	// Engine scheme is 140 X 140

	sum := 0

	for i, line := range inputLines {
		var strNumber string
		for letterIndex, letter := range line {
			if _, err := strconv.Atoi(string(letter)); err == nil {
				// fmt.Printf("%v, ", string(letter))
				strNumber += string(letter)
				if letterIndex+1 > len(line)-1 {
					if i-1 < 0 {
						// check current line
						if checkSpecialCharacter(rune(inputLines[i][letterIndex-len(strNumber)])) {
							sum += helpers.AddNumberToSumAfterConvertion(strNumber)
							strNumber = ""

						}

						// Check line below
						if numToAdd := checkLineBelow(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}
					} else if i+1 > len(inputLines) {
						// check current line
						if checkSpecialCharacter(rune(inputLines[i][letterIndex-len(strNumber)])) {
							sum += helpers.AddNumberToSumAfterConvertion(strNumber)
							strNumber = ""

						}

						// check line above
						if numToAdd := checkLineAbove(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}
					} else {
						// check current line
						if checkSpecialCharacter(rune(inputLines[i][letterIndex-len(strNumber)])) {
							sum += helpers.AddNumberToSumAfterConvertion(strNumber)
							strNumber = ""

						}

						// check line above
						if numToAdd := checkLineAbove(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}

						// Check line below
						if numToAdd := checkLineBelow(strNumber, letterIndex, i, inputLines); sum != 0 {
							sum += numToAdd
							strNumber = ""
						}
					}
				}
			} else {
				if strNumber != "" {
					if i-1 < 0 {
						if letterIndex-len(strNumber)-1 < 0 {
							//check currect val if special character
							if checkSpecialCharacter(letter) {
								sum += helpers.AddNumberToSumAfterConvertion(strNumber)
								strNumber = ""
								break
							}

							// Check line below
							if numToAdd := checkLineBelow(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
								sum += numToAdd
								strNumber = ""
							}
						} else {
							//check currect val if special character
							if checkSpecialCharacter(letter) {
								sum += helpers.AddNumberToSumAfterConvertion(strNumber)
								strNumber = ""

							}

							//check currect line before strNumber if special character
							if checkSpecialCharacter(rune(inputLines[i][letterIndex-len(strNumber)-1])) {
								sum += helpers.AddNumberToSumAfterConvertion(strNumber)
								strNumber = ""

							}

							// Check line below
							if numToAdd := checkLineBelow(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
								sum += numToAdd
								strNumber = ""
							}
						}
					} else if i+1 > len(inputLines)-1 {
						if letterIndex-len(strNumber)-1 < 0 {
							//check currect val if special character
							if checkSpecialCharacter(letter) {
								sum += helpers.AddNumberToSumAfterConvertion(strNumber)
								strNumber = ""

							}

							// check line above
							if numToAdd := checkLineAbove(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
								sum += numToAdd
								strNumber = ""
							}
						} else {
							//check currect val if special character
							if checkSpecialCharacter(letter) {
								sum += helpers.AddNumberToSumAfterConvertion(strNumber)
								strNumber = ""

							}

							//check currect line before strNumber if special character
							if checkSpecialCharacter(rune(inputLines[i][letterIndex-len(strNumber)-1])) {
								sum += helpers.AddNumberToSumAfterConvertion(strNumber)
								strNumber = ""

							}

							// check line above
							if numToAdd := checkLineAbove(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
								sum += numToAdd
								strNumber = ""
							}
						}
					} else if letterIndex-len(strNumber)-1 < 0 {
						//check currect val if special character
						if checkSpecialCharacter(letter) {
							sum += helpers.AddNumberToSumAfterConvertion(strNumber)
							strNumber = ""

						}

						// check line above
						if numToAdd := checkFirstItemLineAbove(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}

						// Check line below
						if numToAdd := checkFirstItemLineBelow(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}
					} else {

						//check currect val if special character
						if checkSpecialCharacter(letter) {
							sum += helpers.AddNumberToSumAfterConvertion(strNumber)
							strNumber = ""

						}

						//check currect line before strNumber if special character
						if checkSpecialCharacter(rune(inputLines[i][letterIndex-len(strNumber)-1])) {
							sum += helpers.AddNumberToSumAfterConvertion(strNumber)
							strNumber = ""

						}

						// check line above
						if numToAdd := checkLineAbove(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}

						// Check line below
						if numToAdd := checkLineBelow(strNumber, letterIndex, i, inputLines); numToAdd != 0 {
							sum += numToAdd
							strNumber = ""
						}
					}
				}
				strNumber = ""
			}
		}
	}

	fmt.Println("The answer is", sum)
}
