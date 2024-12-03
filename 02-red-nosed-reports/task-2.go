package main

import (
	"advent_of_code_2024/utils"
	"math"
)

func getSign2(a int, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

func isSafe2(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	var sign = getSign2(levels[0], levels[1])
	var prev = levels[0]
	for _, level := range levels[1:] {
		if sign != getSign2(prev, level) {
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

func isSafe3(levels []int) bool {
	if isSafe2(levels[1:]) || isSafe2(append([]int{levels[0]}, levels[2:]...)) {
		return true
	}

	var sign = getSign2(levels[0], levels[1])
	var prev = levels[0]
	var skip = false
	for _, level := range levels[1:] {
		if sign != getSign2(prev, level) {
			if skip {
				return false
			} else {
				skip = true
				continue
			}
		}
		abs := math.Abs(float64(prev - level))
		if abs > 3 || abs < 1 {
			if skip {
				return false
			} else {
				skip = true
				continue
			}
		}
		prev = level
	}
	return true
}

func main() {
	lines := utils.MustReadLinesFromFile()

	var counter = 0
	for _, report := range lines {
		levels := utils.IntFields(report)
		if isSafe3(levels) {
			counter++
		}
	}
	println(counter)
}
