package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var moveValues = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissor
}

// readData reads the data from the input file, and returns a slice with each line
// being an individual game.
func readData(path string) ([]string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    return strings.Split(string(data), "\n"), nil
}

// rps checks if I win = 6, draw = 3, or lose = 0
// x = rock, y = paper, z = scissors
func rps(me, you string) int {
	if me == "X" {
		if you == "A" {
			return 3
		}
		if you == "C" {
			return 6
		}
	}

	if me == "Y" {

		if you == "B" {
			return 3
		}
		if you == "A" {
			return 6
		}
	}

	if me == "Z" {
		if you == "C" {
			return 3
		}
		if you == "B" {
			return 6
		}
	}

	return 0
}

// scoreGame takes in a game '<player input> <opponent input>'
// and calculates the score for the player.
func scoreGame(game string) (int, error) {

	g := strings.Split(game, " ")
	if len(g) != 2 {
		return 0, errors.New("game length was incorrect")
	}

	o := g[0]
	p := g[1]

    score := moveValues[p]
    score += rps(p, o)

    return score, nil
}

func main() {
	fmt.Println("a main func")
    tempGame := "C Z"
    path := "./input.txt"

    result, err := scoreGame(tempGame)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Result:", result)

    lines, err := readData(path)
    if err != nil {
        fmt.Println(err)
    }

    total := 0
    for _, game := range lines {
        score, err := scoreGame(game)
        if err != nil {
            fmt.Println(err)
        }

        total += score
    }

    fmt.Println("the result:", total)
}
