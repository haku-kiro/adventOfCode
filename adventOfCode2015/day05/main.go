package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines[:len(lines)-1], nil
}

func ContainsThreeVowels(in string) bool {
	vowels := "aeiou"
	vowelCount := 0
	for _, el := range in {
		if strings.Contains(vowels, string(el)) {
			vowelCount += 1
		}
	}

	if vowelCount >= 3 {
		return true
	}
	return false
}

func ContainsDoubleLetters(in string) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			return true
		}
	}

	return false
}

func IllegalCharFilter(in string) bool {
	r, _ := regexp.Compile("ab|cd|pq|xy")
	check := r.FindString(in)
	if check == "" {
		return true
	}

	return false
}

func HasSkipOneDuplicate(in string) bool {
	for i := 0; i < len(in)-2; i++ {
		if in[i] == in[i+2] {
			return true
		}
	}

	return false
}

func HasDoubleOverlap(in string) bool {
	for i := 0; i < len(in)-1; i++ {
		test := in[i : i+2]
		for ii := i + 2; ii < len(in)-1; ii++ {
			next := in[ii : ii+2]
			if test == next {
				return true
			}
		}
	}

	return false
}

func CollectionChecker(collection []string, checkers []func(string) bool) int {
	result := 0
	checks := len(checkers)
	for _, element := range collection {
		checkCount := 0
		for _, check := range checkers {
			if check(element) {
				checkCount++
			}
		}
		if checkCount == checks {
			result++
		}
	}

	return result
}

func PartOne(collection []string) int {
	checkers := []func(string) bool{
		IllegalCharFilter,
		ContainsDoubleLetters,
		ContainsThreeVowels,
	}

	result := CollectionChecker(collection, checkers)
	return result
}

func PartTwo(collection []string) int {
	checkers := []func(string) bool{
		HasSkipOneDuplicate,
		HasDoubleOverlap,
	}

	result := CollectionChecker(collection, checkers)
	return result
}

func main() {
	path := "./input.txt"
	data, err := readLines(path)
	if err != nil {
		panic(err)
	}

	partOneCount := PartOne(data)
	fmt.Println("Part 1")
	fmt.Println("Good strings:", partOneCount)

	partTwoCount := PartTwo(data)
	fmt.Println("Part 2")
	fmt.Println("Good strings:", partTwoCount)
}
