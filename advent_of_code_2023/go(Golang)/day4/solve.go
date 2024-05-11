package main

import (
	"adventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

func extractInfoFromCards(line string) ([]string, []string) {

	winning := strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[0])
	allNumbers := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[1]), " ")

	var winningSlice []string
	var currentNumber string
	for _, value := range winning {
		if _, err := strconv.Atoi(string(value)); err != nil {
			winningSlice = append(winningSlice, currentNumber)
			currentNumber = ""
			continue
		}

		currentNumber += string(value)
	}

	var allNumbersSlice []string
	currentNumber = ""
	for _, value := range allNumbers {
		if _, err := strconv.Atoi(string(value)); err != nil {
			allNumbersSlice = append(allNumbersSlice, currentNumber)
			currentNumber = ""
			continue
		}

		currentNumber += string(value)
	}

	return winningSlice, allNumbersSlice
}

func countWinningMatches(winning, allNumbers []string) int {
	winningNumbers := helpers.ConvertStringToIntegerSlice(winning)
	allNumbersInt := helpers.ConvertStringToIntegerSlice(allNumbers)
	helpers.Introsort(winningNumbers, 2)
	helpers.Quicksort(allNumbersInt, 0, len(allNumbersInt)-1)
	fmt.Println(winning)
	fmt.Println(allNumbers)

	return 0
}

func main() {
	inputLines := helpers.GetInput("./input.txt")

	for _, line := range inputLines {
		winning, allNumbers := extractInfoFromCards(line)
		countWinningMatches(winning, allNumbers)
	}
}
