package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	length int
	width  int
	height int
}

func (p Present) Slack() int {
	low := p.length * p.width
	next := p.width * p.height
	if next < low {
		low = next
	}
	next = p.height * p.length
	if next < low {
		low = next
	}

	return low
}

func (p Present) WrappingPaper(withSlack bool) int {
	a := 2 * p.length * p.width
	b := 2 * p.width * p.height
	c := 2 * p.height * p.length

	result := a + b + c
	if withSlack {
		result += p.Slack()
	}

	return result
}

func (p Present) Volume() int {
	return p.height * p.width * p.length
}

func (p Present) SmallestPerimeter() int {
	low := 2 * (p.length + p.width)
	next := 2 * (p.width + p.height)
	if next < low {
		low = next
	}
	next = 2 * (p.height + p.length)
	if next < low {
		low = next
	}
	return low
}

func (p Present) RibbonSize() int {
    // Smallest size + volume for bow
    return p.SmallestPerimeter() + p.Volume()
}

func parseLineToPresent(d string) (Present, error) {
	parts := strings.Split(d, "x")
	l, err := strconv.Atoi(parts[0])
	if err != nil {
		return Present{}, err
	}
	w, err := strconv.Atoi(parts[1])
	if err != nil {
		return Present{}, err
	}
	h, err := strconv.Atoi(parts[2])
	if err != nil {
		return Present{}, err
	}

	result := Present{
		length: l,
		width:  w,
		height: h,
	}

	return result, nil
}

func readInputIntoPresents(path string) ([]Present, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stringLines := strings.Split(string(data), "\n")
	// Getting rid of last new line
	stringLines = stringLines[:len(stringLines)-1]

	result := []Present{}
	for _, el := range stringLines {
		p, err := parseLineToPresent(el)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

// PartOne given a collection of presents, calculate the total required feet of wrapping paper required
// with slack.
func PartOne(presents []Present) int {
	result := 0
	for _, present := range presents {
		result += present.WrappingPaper(true)
	}

	return result
}

func PartTwo(presents []Present) int {
    result := 0
    for _, present := range presents {
        result += present.RibbonSize()
    }

    return result
}

func main() {
	path := "./input.txt"
	presents, err := readInputIntoPresents(path)
	if err != nil {
		panic(err)
	}

    // Working!
	feetOfWrappingPaper := PartOne(presents)
	fmt.Println("Part 1")
	fmt.Println("Feet:", feetOfWrappingPaper)

    // Working!
	feetOfRibbon := PartTwo(presents)
	fmt.Println("Part 2")
	fmt.Println("Feet:", feetOfRibbon)
}
