package main

import (
	"fmt"
	"io/ioutil"
)

func readInput(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Last char is a newline?
	return string(data[:len(data)-1]), nil
}

// PartOne parses the instructions, and goes up for '(' and down for ')'
// the result being the floor you end up on.
func PartOne(input string) int {
	currentFloor := 0
	for _, el := range input {
		goUp := string(el) == "("
		// Debug print
		// fmt.Println("char", string(el), "check", goUp)
		if goUp {
			currentFloor += 1
		} else {
			currentFloor -= 1
		}
	}

	return currentFloor
}

// PartTwo is to find out which instruction (1 based index) is the first instruction
// to make the floor negative
func PartTwo(input string) int {
	currentFloor := 0
	for i, el := range input {
		goUp := string(el) == "("
		if goUp {
			currentFloor += 1
		} else {
			currentFloor -= 1
		}

		if currentFloor < 0 {
			return i + 1
		}
	}

	return -1
}

func main() {
	inputPath := "./input.txt"
	data, err := readInput(inputPath)
	if err != nil {
		panic(err)
	}

    // Working
	floor := PartOne(data)

	fmt.Println("Part 1")
	fmt.Println("Floor:", floor)

    // Working
	instruction := PartTwo(data)

	fmt.Println("Part 2")
	fmt.Println("Instruction", instruction)
}
