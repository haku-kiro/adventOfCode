package main

import (
	"fmt"
	"os"
)

func readDataStream(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func streamData(data string, startPoint int) int {
    startPoint = startPoint -1
	for i := startPoint; i < len(data); i++ {
		segment := data[i-startPoint : i+1]
		if distinct(segment) {
			return i + 1
		}
	}

	return -1
}

func distinct(segment string) bool {
	for i, el1 := range segment {
		for ii, el2 := range segment {
			if i == ii {
				continue
			}

			if el1 == el2 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("starting...")
	path := "./input.txt"

	data, err := readDataStream(path)
	if err != nil {
		fmt.Println(err)
	}

    part1 := 4
    part2 := 14
    fmt.Printf("Part 1 find first distinct at %d length\n", part1)
    fmt.Printf("Part 2 find first distinct at %d length\n", part2)

	resultPart1 := streamData(data, part1)
	resultPart2 := streamData(data, part2)

    fmt.Println("Part 1", resultPart1)
    fmt.Println("Part 2", resultPart2)
}
