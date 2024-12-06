package main

import (
	"advent_of_code_2024/utils"
	"math"
	"strings"
)

func main() {
	lines := utils.MustReadLinesFromFile()

	rules := make(map[string][]string)

	result := 0
	for _, line := range lines {

		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			_, exists := rules[pages[1]]
			if !exists {
				rules[pages[1]] = []string{}
			}
			rules[pages[1]] = append(rules[pages[1]], pages[0])
		} else if line != "" {
			pages := strings.Split(line, ",")
			if !isUpdateOk2(pages, rules) {
				new := correctOrder(pages, rules)
				println(strings.Join(new, ","))
				result += utils.MustAtoi(new[(len(pages)-1)/2])
			}
		}
	}
	println(result)

}

func correctOrder(pages []string, rules map[string][]string) []string {
	visited := make(map[string]int)
	for index, page := range pages {
		visited[page] = index
	}
	for index, page := range pages {
		maxIndex := isPageOk2(rules, page, index, visited)
		if maxIndex > -1 {
			final := moveElement(pages, index, maxIndex)
			return correctOrder(final, rules)
		}
	}
	return pages
}

func isUpdateOk2(pages []string, rules map[string][]string) bool {
	visited := make(map[string]int)
	for index, page := range pages {
		visited[page] = index
	}
	//fmt.Println(visited)
	for index, page := range pages {
		if isPageOk2(rules, page, index, visited) > -1 {
			//fmt.Println("not ok:", pages)
			return false
		}
	}
	//fmt.Println("ok:", pages)
	return true
}

func isPageOk2(rules map[string][]string, page string, pageIndex int, visited map[string]int) int {
	maxIndex := -1
	for _, rule := range rules[page] {
		if visited[rule] > pageIndex {
			maxIndex = int(math.Max(float64(visited[rule]), float64(maxIndex)))
		}
	}
	return maxIndex
}

func moveElement(list []string, srcIndex, destIndex int) []string {
	element := list[srcIndex]
	list = append(list[:srcIndex], list[srcIndex+1:]...)
	list = append(list[:destIndex], append([]string{element}, list[destIndex:]...)...)
	return list
}
