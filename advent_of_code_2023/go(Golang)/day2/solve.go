package main

import (
	"adventOfCode/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Get the game ID => Game **14**:
func getGameID(line string) int {
	startIndex := 5
	var numString string

	for line[startIndex] != ':' {
		numString += string(line[startIndex])
		startIndex++
	}

	number, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatalf("This game ID is not formatted correctly: ! %v !", line)
	}

	return number
}

// Get starting index point of rounds in the game
func getStartGameIndex(line string) int {
	for index, value := range line {
		if value == ':' {
			return index + 1
		}
	}

	log.Fatalf("Game not formatted correctly: ! %v !", line)
	return 0
}

// Prepare the rounds within a given game as a slice of strings
func getGameInfo(line string) []string {
	startingIndex := getStartGameIndex(line)
	gameInfo := line[startingIndex+1:]

	gameInfoParts := strings.Split(gameInfo, "; ")

	return gameInfoParts
}

// Solve the first part of the problem using all the above functions given a full game string
func solveFirstPart(line string) int {
	gameID := getGameID(line)
	gameInfo := getGameInfo(line)
	for _, game := range gameInfo {
		gameSplit := strings.Split(game, ", ")
		for _, round := range gameSplit {
			roundInfo := strings.Split(round, " ")
			switch roundInfo[1] {
			case "red":
				if val, err := strconv.Atoi(roundInfo[0]); val > 12 && err == nil {
					return 0
				}
			case "green":
				if val, err := strconv.Atoi(roundInfo[0]); val > 13 && err == nil {
					return 0
				}
			case "blue":
				if val, err := strconv.Atoi(roundInfo[0]); val > 14 && err == nil {
					return 0
				}
			default:
				return 0
			}
		}
	}

	return gameID
}

func solveSecondPart(line string) int {

	gameInfo := getGameInfo(line)
	var maxes [3]int // [1 - red, 2 - green, 3 - blue]
	for _, game := range gameInfo {
		gameSplit := strings.Split(game, ", ")
		for _, round := range gameSplit {
			roundInfo := strings.Split(round, " ")
			switch roundInfo[1] {
			case "red":
				if val, err := strconv.Atoi(roundInfo[0]); val > maxes[0] && err == nil {
					maxes[0] = val
				}
			case "green":
				if val, err := strconv.Atoi(roundInfo[0]); val > maxes[1] && err == nil {
					maxes[1] = val
				}
			case "blue":
				if val, err := strconv.Atoi(roundInfo[0]); val > maxes[2] && err == nil {
					maxes[2] = val
				}
			}
		}
	}

	powerResult := maxes[0] * maxes[1] * maxes[2]

	return powerResult
}

// Main function
func main() {
	inputLines := helpers.GetInput("./input.txt")

	sumFirstPart := 0
	sumSecondPart := 0
	for _, line := range inputLines {
		sumFirstPart += solveFirstPart(line)
		sumSecondPart += solveSecondPart(line)
	}

	fmt.Println("First part answer: ", sumFirstPart)
	fmt.Println("Second part answer: ", sumSecondPart)
}
