package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	ascii_text, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var text []string = strings.Split(string(ascii_text), "\r")
	// [0 : len(text[0])/2] // first half
	// [len(text[0])/2 : len(text[0])] // second half

	var total int = 0

	// day 3 - part one

	// for _, v := range text {
	// 	var properV string = strings.TrimPrefix(v, "\n")
	// 	var first string = properV[0 : len(properV)/2]
	// 	var second string = properV[len(properV)/2:]
	// mainLoop:
	// 	for _, letter1 := range first {
	// 		for _, letter2 := range second {
	// 			if letter1 == letter2 {
	// 				if unicode.IsUpper(letter1) { // uppercase
	// 					total += int(letter1) - 38
	// 				} else { // lowercase
	// 					total += int(letter1) - 96
	// 				}
	// 				break mainLoop
	// 			}
	// 		}
	// 	}
	// }

	var trio []string
	for i, v := range text {
		trio = append(trio, strings.TrimPrefix(v, "\n"))
		if (i-2)%3 == 0 {
		mainLoop:
			for _, letter1 := range trio[0] {
				for _, letter2 := range trio[1] {
					if letter1 == letter2 {
						for _, letter3 := range trio[2] {
							if letter1 == letter3 {
								if unicode.IsUpper(letter1) { // uppercase
									total += int(letter1) - 38
								} else { // lowercase
									total += int(letter1) - 96
								}
								break mainLoop
							}
						}
					}
				}
			}
			trio = nil
		}
	}
	fmt.Println(total)
}
