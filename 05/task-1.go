package main

import (
	"advent_of_code_2024/utils"
	"strings"
)

func main() {
	lines := utils.MustReadLinesFromFile()

	rules := make(map[string][]string)

	result := 0
	for _, line := range lines {
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			_, exists := rules[pages[0]]
			if !exists {
				rules[pages[0]] = []string{}
			}
			rules[pages[0]] = append(rules[pages[0]], pages[1])
		} else if line != "" {
			pages := strings.Split(line, ",")
			if isUpdateOk(pages, rules) {
				result += utils.MustAtoi(pages[(len(pages)-1)/2])
			}
		}
	}
	println(result)

}

func isUpdateOk(pages []string, rules map[string][]string) bool {
	visited := make(map[string]int)
	for _, page := range pages {
		visited[page] = 1
		if !isPageOk(rules, page, visited) {
			return false
		}
	}
	return true
}

func isPageOk(rules map[string][]string, page string, visited map[string]int) bool {
	for _, rule := range rules[page] {
		if visited[rule] == 1 {
			return false
		}
	}
	return true
}
