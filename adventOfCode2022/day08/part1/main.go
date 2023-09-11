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
	sum := 0
	for x := range m {
		for y := range m[x] {
			// boundries
			if x == 0 || x == len(m)-1 || y == 0 || y == len(m[x])-1 {
				continue
			}
			// Doing this block because we want to stop processing
			// if we see that it's visable from any side
			visable := checkUp(m, x, y)
			if visable {
				sum += 1
			} else {
				visable = checkRight(m, x, y)
				if visable {
					sum += 1
				} else {
					visable = checkDown(m, x, y)
					if visable {
						sum += 1
					} else {
						visable = checkLeft(m, x, y)
						if visable {
							sum += 1
						}
					}
				}
			}
		}
	}

	fmt.Println("sum", sum)
    return sum
}


func checkUp(m [][]int, x, y int) bool {
    checkVal := m[x][y]
    for i := y-1; i >= 0; i-- {
        if checkVal <= m[x][i] {
            return false
        }
    }

    return true
}

func checkDown(m [][]int, x, y int) bool {
    checkVal := m[x][y]
    for i := y+1; i < len(m); i++ {
        if checkVal <= m[x][i] {
            return false
        }
    }

    return true
}

func checkRight(m [][]int, x, y int) bool {
    checkVal := m[x][y]
    for i := x+1; i < len(m[x]); i++ {
        if checkVal <= m[i][y] {
            return false
        }
    }

    return true
}

func checkLeft(m [][]int, x, y int) bool {
    checkVal := m[x][y]
    for i := x-1; i >= 0; i-- {
        if checkVal <= m[i][y] {
            return false
        }
    }

    return true
}

func calcPerimiter(in int) int {
    if in <= 2 && in > 0 {
        return in * in
    } else if in > 2 {
        a := in * in
        b := (in-2) * (in-2)
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
    // Assuming square
    outsideArea := calcPerimiter(len(mapArea))
    fmt.Println(outsideArea)

    inner := iterateOverMap(mapArea)
    fmt.Println("final", inner + outsideArea)
}
