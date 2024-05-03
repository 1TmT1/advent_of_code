package helpers

import (
	"bufio"
	"log"
	"os"
)

// bufio implementation - read input file
func GetInput(path string) []string {
	text, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer text.Close()

	fileScanner := bufio.NewScanner(text)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
