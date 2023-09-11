package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var rps = map[string]int{
	"r": 1, // Rock
	"p": 2, // Paper
	"s": 3, // Scissor
}


var outcomeValues = map[string]int{
	"X": 0, // Lose
	"Y": 3, // draw
	"Z": 6, // win
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

// pick decides on the move you will execute based on their move
// and the required outcome
func pick(them, outcome string) int {
    // lose
	if outcome == "X" {
		if them == "A" {
			return rps["s"]
		}
		if them == "B" {
			return rps["r"]
		}
		if them == "C" {
			return rps["p"]
		}
	}

    // draw
	if outcome == "Y" {
		if them == "A" {
			return rps["r"]
		}
		if them == "B" {
			return rps["p"]
		}
		if them == "C" {
			return rps["s"]
		}
	}

    // win
	if outcome == "Z" {
		if them == "A" {
			return rps["p"]
		}
		if them == "B" {
			return rps["s"]
		}
		if them == "C" {
			return rps["r"]
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

	theirChoice := g[0]
	requiredOutcome := g[1]

    score := outcomeValues[requiredOutcome]
    score += pick(theirChoice, requiredOutcome)

    return score, nil
}

func main() {
	fmt.Println("a main func")
	tempGame := "A Y"
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
