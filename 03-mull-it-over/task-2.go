package main

import (
	"advent_of_code_2024/utils"
	"regexp"
	"strings"
)

func main() {
	lines := utils.MustReadLinesFromFile()
	input := strings.Join(lines, " ")
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	var result = 0
	split := strings.Split(input, "do()")
	for _, s := range split {
		memory := strings.Split(s, "don't()")[0]
		matches := regex.FindAllStringSubmatch(memory, -1)
		for _, match := range matches {
			result += utils.MustAtoi(match[1]) * utils.MustAtoi(match[2])
		}
	}
	println(result)
}
