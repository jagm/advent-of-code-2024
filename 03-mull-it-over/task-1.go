package main

import (
	"advent_of_code_2024/utils"
	"regexp"
)

func main() {
	lines := utils.MustReadLinesFromFile()
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	var result = 0
	for _, input := range lines {
		matches := regex.FindAllStringSubmatch(input, -1)
		for _, match := range matches {
			result += utils.MustAtoi(match[1]) * utils.MustAtoi(match[2])
		}
	}
	println(result)
}
