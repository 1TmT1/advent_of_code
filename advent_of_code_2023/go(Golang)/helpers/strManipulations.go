package helpers

// Reverse string ("hello" => "olleh")
func ReverseString(line string) string {
	var reversedString string

	for _, char := range line {
		reversedString = string(char) + reversedString
	}

	return reversedString
}
