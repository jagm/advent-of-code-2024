package utils

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert '%s' to int: %v", s, err)
	}
	return num
}

func MustReadLinesFromFile() []string {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the file name as an argument.")
	}
	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	} else {
		return lines
	}
}

func Fields(s string) []string {
	splitter := regexp.MustCompile(`\s+`)
	return splitter.Split(s, -1)
}

func IntFields(s string) []int {
	splitter := regexp.MustCompile(`\s+`)
	var result []int
	for _, v := range splitter.Split(s, -1) {
		result = append(result, MustAtoi(v))
	}
	return result
}
