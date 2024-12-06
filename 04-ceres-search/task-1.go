package main

import (
	"advent_of_code_2024/utils"
	"strings"
)

func main() {
	lines := utils.MustReadLinesFromFile()

	result := 0
	for i, input := range lines {
		result += strings.Count(input, "XMAS")
		result += strings.Count(input, "SAMX")
		for j, char := range input {
			if char == 'X' {
				if i >= 3 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					result++
				}
				if i+3 < len(lines) && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					result++
				}
				if i >= 3 && j >= 3 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					result++
				}
				if i+3 < len(lines) && j >= 3 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					result++
				}
				if i >= 3 && j+3 < len(input) && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					result++
				}
				if i+3 < len(lines) && j+3 < len(input) && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					result++
				}
			}
		}
	}
	println(result)
}

func s(input uint8) string {
	return string(input)
}
