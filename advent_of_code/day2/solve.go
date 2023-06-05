package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const rock, scissors, paper uint = 1, 3, 2
	ascii_text, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	var text string = string(ascii_text)

	var text_lines []string = strings.Split(text, "\n")

	var score uint = 0
	// day 2 - part one
	// for _, v := range text_lines {
	// 	// strings.TrimSpace(v[len(v)-2:]) - my turn
	// 	// v[:1] - opponent turn

	// 	me := strings.TrimSpace(v[len(v)-2:])
	// 	opponent := strings.TrimSpace(v[:1])

	// 	switch {
	// 	case me == "X": // Rock
	// 		score += rock
	// 	case me == "Y": // Paper
	// 		score += paper
	// 	default: // Scissors
	// 		score += scissors
	// 	}
	// 	if int(me[0])-23 == int(opponent[0]) {
	// 		// draw
	// 		score += 3
	// 	} else if me == "X" && opponent == "C" || me == "Y" && opponent == "A" || me == "Z" && opponent == "B" {
	// 		// win
	// 		score += 6
	// 	}
	// }

	for _, v := range text_lines {
		// strings.TrimSpace(v[len(v)-2:]) - result
		// v[:1] - opponent

		var result string = strings.TrimSpace(v[len(v)-2:])
		var opponent string = v[:1]
		switch {
		case result == "X": // lose
			switch {
			case opponent == "A": // rock
				score += scissors
			case opponent == "B": // paper
				score += rock
			default: // scissors
				score += paper
			}
		case result == "Y": // draw
			switch {
			case opponent == "A": // rock
				score += rock
			case opponent == "B": // paper
				score += paper
			default: // scissors
				score += scissors
			}
			score += 3
		default: // win
			switch {
			case opponent == "A": // rock
				score += paper
			case opponent == "B": // paper
				score += scissors
			default: // scissors
				score += rock
			}
			score += 6
		}

	}
	fmt.Println(score)

}
