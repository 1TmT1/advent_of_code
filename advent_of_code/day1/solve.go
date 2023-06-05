package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "testing"
	"strconv"
	"strings"
)

func Test() {
	// bufio implementation - read file
	address_text, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(address_text)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	address_text.Close()

	// for _, line := range fileLines{
	// 	fmt.Println(line)
	// }

	// os(ioutil deprecated) implementation - read file

	// ascii_text, err := os.ReadFile("./input.txt")
	// if err != nil{
	// 	log.Fatalln(err)
	// }

	// var text string = string(ascii_text)

	// var text_lines []string = strings.Split(text, "\n")

	// var max int = 0
	// var total int = 0
	// var counter_total int = 0
	// for _, line := range text_lines {
	// 	if line == "…"{
	// 		break
	// 	}
	// 	if line == "\r"{
	// 		if max < total {
	// 			max = total
	// 		}
	// 		counter_total += total
	// 		total = 0
	// 	}else{
	// 		num, err := strconv.Atoi(strings.TrimSuffix(line, "\r"))
	// 		if err != nil{
	// 			log.Fatalln(err)
	// 		}
	// 		total += num
	// 	}
	// }
	// fmt.Println(max)
	// fmt.Println(counter_total)

	var total []int
	var elf_total int = 0
	for _, line := range fileLines {
		if line == "…" {
			break
		}
		if line == "" {
			total = append(total, elf_total)
			elf_total = 0
		} else {
			num, err := strconv.Atoi(strings.TrimSuffix(line, "\r"))
			if err != nil {
				log.Fatalln(err)
			}
			elf_total += num
		}
	}

	var maxes [3]int // 0 - first, 1 - second, 2 - third

	for _, max := range total {
		if maxes[0] < max {
			maxes[2] = maxes[1]
			maxes[1] = maxes[0]
			maxes[0] = max
		} else if maxes[1] < max && max != maxes[0] {
			maxes[2] = maxes[1]
			maxes[1] = max
		} else if maxes[2] < max && max != maxes[1] {
			maxes[2] = max
		}
	}
	fmt.Println(maxes)

}

func main() {
	fmt.Println("hello!")
}
