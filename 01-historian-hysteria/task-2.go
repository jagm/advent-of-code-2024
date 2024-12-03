package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the file name as an argument.")
	}
	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	splitter := regexp.MustCompile(`\s+`)

	left := make(map[int]int)
	right := make(map[int]int)

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := splitter.Split(line, -1)
		number1, _ := strconv.Atoi(numbers[0])
		left[number1]++
		number2, _ := strconv.Atoi(numbers[1])
		right[number2]++
	}

	var result int
	for index, element := range left {
		result += element * index * right[index]
	}

	println(result)
}
