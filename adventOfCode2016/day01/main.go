// main package
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type dir struct {
	loc string
}

func newDir() *dir {
	return &dir{
		loc: "up",
	}
}

func (d *dir) change(instruction string) {
	if instruction == "R" {
		switch d.loc {
		case "up":
			d.loc = "right"
		case "right":
			d.loc = "down"
		case "down":
			d.loc = "left"
		case "left":
			d.loc = "up"
		}
	} else if instruction == "L" {
		switch d.loc {
		case "up":
			d.loc = "left"
		case "left":
			d.loc = "down"
		case "down":
			d.loc = "right"
		case "right":
			d.loc = "up"
		}
	}
}

func (p *point) up(distance int) {
	p.y += distance
}

func (p *point) down(distance int) {
	p.y -= distance
}

func (p *point) left(distance int) {
	p.x -= distance
}

func (p *point) right(distance int) {
	p.x += distance
}

func readData(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(data), ",")
	for x := 0; x < len(temp); x++ {
		temp[x] = strings.ReplaceAll(temp[x], " ", "")
		temp[x] = strings.ReplaceAll(temp[x], "\n", "")
	}

	return temp, err
}

func splitInstruction(instruction string) []string {
	return strings.Split(instruction, "")
}

// distanceBetweenTwoPoints gets the line drawn between the two points distance
func distanceBetweenTwoPoints(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
}

// distanceBetweenTwoPointsTaxicabGrid gets the distance between two points assuming they can only travel
// on a grid. Alg taken from: https://www.wikiwand.com/en/Taxicab_geometry
// note: the p1 is redundant here because we always start at origin, but this works even if we start at a different point.
func distanceBetweenTwoPointsTaxicabGrid(p1, p2 point) float64 {
	return math.Abs(float64(p1.x)-float64(p2.x)) + math.Abs(float64(p1.y)-float64(p2.y))
}

func main() {
	dataPath := "./input.txt"
	d, err := readData(dataPath)
	if err != nil {
		panic(err)
	}

	dir := newDir()
	walkPoint := point{
		x: 0,
		y: 0,
	}
	origin := point{
		x: 0,
		y: 0,
	}

	// For the second part (the bad section of code below)
	visitedPoints := []point{origin}
	foundFirst := false
	firstDouble := point{}

	for x := 0; x < len(d); x++ {
		instructions := splitInstruction(d[x])

		// first change dir, then walk.
		dir.change(instructions[0])
		// Issue was not parsing the entire remaining split for the whole number (numbers with more than 2 digits...)
		tempNum := strings.Join(instructions[1:], "")
		units, err := strconv.Atoi(tempNum)
		if err != nil {
			panic(err)
		}

		for x := 0; x < units; x++ {
			if dir.loc == "up" {
				walkPoint.up(1)
			}
			if dir.loc == "down" {
				walkPoint.down(1)
			}
			if dir.loc == "left" {
				walkPoint.left(1)
			}
			if dir.loc == "right" {
				walkPoint.right(1)
			}

			// Add the point, and check if we've been there before.
			// This is a very slow solution (well naive - at least)

			if !foundFirst {
				for x := 0; x < len(visitedPoints); x++ {
					if visitedPoints[x].x == walkPoint.x && visitedPoints[x].y == walkPoint.y {
						foundFirst = true
						firstDouble = walkPoint
						fmt.Println("first double point: ", firstDouble)
					}
				}

				visitedPoints = append(visitedPoints, walkPoint)
			}
		}
	}

	// Part 1
	fmt.Println("final location: ", walkPoint)
	// Works, but we can't go diagonally - we have to go across the grids.
	fmt.Println("straight line distance from start to final point: ", distanceBetweenTwoPoints(origin, walkPoint))
	// This makes use of the logic of taxi cab grid movement, i.e. only moving in a grid.
	fmt.Println("taxi cab grid distance between start and final point: ", distanceBetweenTwoPointsTaxicabGrid(origin, walkPoint))

	// Part 2
	fmt.Println("taxi cab grid distance between start and first duplicate (visited twice) point: ", distanceBetweenTwoPointsTaxicabGrid(origin, firstDouble))
}
