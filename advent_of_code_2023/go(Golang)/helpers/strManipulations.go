package helpers

import (
	"strconv"
)

// Reverse string ("hello" => "olleh")
func ReverseString(line string) string {
	var reversedString string

	for _, char := range line {
		reversedString = string(char) + reversedString
	}

	return reversedString
}

// Convert string that represent an integer to type int => string = "12" -> int = 12
func AddNumberToSumAfterConvertion(strNumber string) int {
	num, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0
		// log.Fatalf("%v is not a number that can be converted...", num)
	}

	return num
}
