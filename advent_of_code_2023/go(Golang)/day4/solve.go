package main

import (
	"adventOfCode/helpers"
	"fmt"
	"strconv"
	"strings"
)

func extractInfoFromCards(line string) ([]string, []string) {

	winning := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[0]), " ")
	allNumbers := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[1]), " ")

	var winningSlice []string
	for _, value := range winning {
		if _, err := strconv.Atoi(string(value)); err == nil {
			winningSlice = append(winningSlice, value)
		}
	}

	var allNumbersSlice []string
	for _, value := range allNumbers {
		if _, err := strconv.Atoi(string(value)); err == nil {
			allNumbersSlice = append(allNumbersSlice, value)
		}
	}

	return winningSlice, allNumbersSlice
}

func countWinningMatches(winning, allNumbers []string) int {
	winningNumbers := helpers.ConvertStringToIntegerSlice(winning)
	allNumbersInt := helpers.ConvertStringToIntegerSlice(allNumbers)
	helpers.Introsort(winningNumbers, 2)
	helpers.Introsort(allNumbersInt, 2)

	var result int
	for _, v := range winningNumbers {
		// v found in both winning and all numbers
		if helpers.BinarySearch(allNumbersInt, v) != -1 {
			if result == 0 {
				result++
			} else {
				result *= 2
			}
		}
	}

	return result
}

func main() {
	inputLines := helpers.GetInput("./input.txt")

	var sum int
	for _, line := range inputLines {
		winning, allNumbers := extractInfoFromCards(line)
		sum += countWinningMatches(winning, allNumbers)
	}

	fmt.Println("The answer is", sum)
}
