package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const logging bool = false

type instruction struct {
	direction string
	distance  int
}

func readFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(data), "\n")
	return temp[:len(temp)-1], nil
}

func parseInstructions(data []string) []instruction {
	result := []instruction{}
	for _, l := range data {
		tokens := strings.Split(l, " ")
		distance, _ := strconv.Atoi(tokens[1])
		i := instruction{
			direction: tokens[0],
			distance:  distance,
		}
		result = append(result, i)
	}

	return result
}

// Things like this would make it simpler to mentally parse solutions
const x, y = 0, 1

func walk(ins []instruction) map[[2]int]bool {
	set := map[[2]int]bool{}
	var head, tail [2]int
	count := 0

	for _, i := range ins {
		move := [2]int{}
		switch i.direction {
		case "U":
			move[y] = i.distance
		case "D":
			move[y] = -i.distance
		case "R":
			move[x] = i.distance
		case "L":
			move[x] = -i.distance
		}

		if logging {
			fmt.Println("\ndefined move", move)
		}

		for move[x] != 0 || move[y] != 0 {
			switch {
			// Moving left
			case move[x] < 0:
				move[x]++
				head[x]--
				// moving right
			case move[x] > 0:
				move[x]--
				head[x]++
				// moving down
			case move[y] < 0:
				move[y]++
				head[y]--
				// moving up
			case move[y] > 0:
				move[y]--
				head[y]++
			}

			dx, dy := head[x]-tail[x], head[y]-tail[y]
			if logging {
				fmt.Printf("head x=%d; head y=%d\n", head[x], head[y])
				fmt.Printf("tail x=%d; tail y=%d\n", tail[x], tail[y])
				fmt.Printf("dx=%d; dy=%d\n", dx, dy)
			}
			switch {
			case dx > 1:
				tail[x]++
				tail[y] += dy
			case dx < -1:
				tail[x]--
				tail[y] += dy
			case dy > 1:
				tail[y]++
				tail[x] += dx
			case dy < -1:
				tail[y]--
				tail[x] += dx
			}

			if !set[tail] {
				set[tail] = true
				count++
			}
		}
	}

	fmt.Println("count", count)
	return set
}

func main() {
	fmt.Println("starting")
	path := "./input.txt"
	data, err := readFile(path)
	if err != nil {
		fmt.Println(err)
	}

	ins := parseInstructions(data)
	walk(ins)
}
