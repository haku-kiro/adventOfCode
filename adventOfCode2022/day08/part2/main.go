package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readData(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(data), "\n")
	return temp[:len(temp)-1], nil
}

func buildMap(data []string) [][]int {
	result := [][]int{}
	for _, row := range data {
		temp := strings.Split(row, "")
		r := []int{}
		for _, el := range temp {
			intEl, _ := strconv.Atoi(el)
			r = append(r, intEl)
		}
		result = append(result, r)
	}

	return result
}

func iterateOverMap(m [][]int) int {
	max := 0
	for x := range m {
		for y := range m[x] {
			// boundries
			if x == 0 || x == len(m)-1 || y == 0 || y == len(m[x])-1 {
				continue
			}

			// fmt.Printf("Checking for x=%d, y=%d -> %d\n", x, y, m[x][y])
			upScore := checkUp(m, x, y)
			downScore := checkDown(m, x, y)
			leftScore := checkLeft(m, x, y)
			rightScore := checkRight(m, x, y)

			test := upScore * downScore * leftScore * rightScore
			if test > max {
				max = test
			}
		}
	}

	return max
}

func checkLeft(m [][]int, x, y int) int {
	checkVal := m[x][y]
	count := 0
	for i := y - 1; i >= 0; i-- {
		if checkVal <= m[x][i] {
			count += 1
			return count
		} else {
			count += 1
		}
	}

	if count == 0 {
		return 1
	}
	return count
}

func checkRight(m [][]int, x, y int) int {
	checkVal := m[x][y]
	count := 0
	for i := y + 1; i < len(m); i++ {
		if checkVal <= m[x][i] {
			count += 1
			return count
		} else {
			count += 1
		}
	}

	if count == 0 {
		return 1
	}
	return count
}

func checkDown(m [][]int, x, y int) int {
	checkVal := m[x][y]
	count := 0
	for i := x + 1; i < len(m[x]); i++ {
		if checkVal <= m[i][y] {
			count += 1
			return count
		} else {
			count += 1
		}
	}

	if count == 0 {
		return 1
	}
	return count
}

func checkUp(m [][]int, x, y int) int {
	checkVal := m[x][y]
	count := 0
	for i := x - 1; i >= 0; i-- {
		if checkVal <= m[i][y] {
			count += 1
			return count
		} else {
			count += 1
		}
	}

	if count == 0 {
		return 1
	}
	return count
}

func calcPerimiter(in int) int {
	if in <= 2 && in > 0 {
		return in * in
	} else if in > 2 {
		a := in * in
		b := (in - 2) * (in - 2)
		return a - b
	}
	return -1
}

func main() {
	fmt.Println("starting")
	path := "./input.txt"
	data, err := readData(path)
	if err != nil {
		fmt.Println(err)
	}

	mapArea := buildMap(data)

	// too high: 1199520
	bestScenicScore := iterateOverMap(mapArea)
	fmt.Println("best scenic score:", bestScenicScore)
}
