package main

import (
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func readInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data)[:len(string(data))-1], nil
}

func Move(p Point, dir string) Point {
	if dir == "<" {
		p.x--
	} else if dir == ">" {
		p.x++
	} else if dir == "^" {
		p.y++
	} else if dir == "v" {
		p.y--
	}

	return p
}

func MakeKey(p Point) string {
	return fmt.Sprintf("%d|%d", p.x, p.y)
}

func Walk(instructions string) map[string]int {
	p := Point{0, 0}
	result := make(map[string]int)
	key := MakeKey(p)
	result[key] = 1

	for _, instruction := range instructions {
		i := string(instruction)
		p = Move(p, i)

		key = MakeKey(p)
		_, exists := result[key]
		if exists {
			result[key]++
		} else {
			result[key] = 1
		}
	}

	return result
}

func WalkTwice(instructions string) map[string]int {
    santaPoint := Point{0, 0}
    roboPoint := Point{0, 0}

	key := MakeKey(santaPoint)
	result := make(map[string]int)
	result[key] = 1

	for idx, el := range instructions {
		i := string(el)
		if idx%2 == 0 {
			santaPoint = Move(santaPoint, i)
			key = MakeKey(santaPoint)
		} else {
			roboPoint = Move(roboPoint, i)
			key = MakeKey(roboPoint)
		}

		_, exists := result[key]
		if exists {
			result[key]++
		} else {
			result[key] = 1
		}
	}

	return result
}

func PartOne(instructions string) int {
	mappedResult := Walk(instructions)
	return len(mappedResult)
}

func PartTwo(instructions string) int {
	mappedResult := WalkTwice(instructions)
	return len(mappedResult)
}

func main() {
	path := "./input.txt"
	instructions, err := readInput(path)
	if err != nil {
		panic(err)
	}

	// Working!
	houses := PartOne(instructions)
	fmt.Println("Part 1")
	fmt.Println("Houses:", houses)

	// Working!
	houses = PartTwo(instructions)
	fmt.Println("Part 2")
	fmt.Println("Houses:", houses)
}
