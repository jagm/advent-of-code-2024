package main

import (
	"advent_of_code_2024/utils"
	"math"
)

func getSign(a int, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	var sign = getSign(levels[0], levels[1])
	var prev = levels[0]
	for _, level := range levels[1:] {
		if sign != getSign(prev, level) {
			return false
		}
		abs := math.Abs(float64(prev - level))
		if abs > 3 || abs < 1 {
			return false
		}
		prev = level
	}
	return true
}

func main() {
	lines := utils.MustReadLinesFromFile()

	var counter = 0
	for _, report := range lines {
		if isSafe(utils.IntFields(report)) {
			counter++
		}
	}
	println(counter)
}
