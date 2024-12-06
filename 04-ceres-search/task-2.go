package main

import (
	"advent_of_code_2024/utils"
)

func main() {
	lines := utils.MustReadLinesFromFile()

	result := 0
	for i, input := range lines {
		for j, char := range input {
			if char == 'A' {
				if i > 0 && j > 0 && i < len(lines)-1 && j < len(input)-1 {
					if ((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) &&
						((lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') || (lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M')) {
						result++
					}
				}
			}
		}
	}
	println(result)
}
