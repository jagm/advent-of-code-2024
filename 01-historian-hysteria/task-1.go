package main

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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

	var left, right []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := splitter.Split(line, -1)
		number1, _ := strconv.Atoi(numbers[0])
		left = append(left, number1)
		number2, _ := strconv.Atoi(numbers[1])
		right = append(right, number2)
	}
	sort.Ints(left)
	sort.Ints(right)

	var result float64
	for index, element := range left {
		result += math.Abs(float64(element - right[index]))
	}

	println(int(result))
}
