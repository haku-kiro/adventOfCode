package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	index    int
	contents []string
}

type ship struct {
	stacks []stack
}

// This is to read the moves/ instructions
func readFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(data), "\n")
	return temp[:len(temp)-1], nil
}

// readBoard takes a path and reads the contents of the file into a
// ship object
func readBoard(path string) (ship, error) {
	data, err := readFile(path)
	s := ship{}
	if err != nil {
		return s, err
	}

	for _, e := range data {
		index := 0
		check := 0
		cell := ""
		skipSpace := false
		for _, c := range e {
			if skipSpace {
				skipSpace = false
				continue
			}
			cell += string(c)
			check += 1

			if check == 3 {
				index += 1
				// maybe parse cell properly?
				shouldAdd := strings.Contains(cell, "[")
				if shouldAdd {
					updateShipStack(index, cell, &s)
				}

				check = 0
				cell = ""
				skipSpace = true
			}
		}
		index = 0
	}

	// Make sure each stack is in correct order by reversing
	for _, stack := range s.stacks {
		reverseSlice(&stack.contents)
	}

	return s, nil
}

func reverseSlice(data *[]string) {
	i := 0
	j := len(*data) - 1
	for i < j {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
		i++
		j--
	}
}

func updateShipStack(index int, element string, s *ship) {
	for i, stack := range s.stacks {
		if stack.index == index {
			// Can't modify indirectly
			// stack.contents = append(stack.contents, element)
			s.stacks[i].contents = append(s.stacks[i].contents, element)
			return
		}
	}

	stack := stack{index: index, contents: []string{element}}
	s.stacks = append(s.stacks, stack)
}

// parseInstruction takes a string and parses into 3 integer values
// that are used for running an operation against a ship
// 1 - Amount to move
// 2 - from index (non-zero based)
// 3 - to index (non-zero based)
func parseInstruction(move string) (int, int, int) {
	// Assuming data is always correct
	data := strings.Split(move, " ")
	amount, _ := strconv.Atoi(data[1])
	from, _ := strconv.Atoi(data[3])
	to, _ := strconv.Atoi(data[5])

	return amount, from, to
}

// executeInstruction takes in commands and runs them against the ship.
func executeInstruction(count, from, to int, s *ship) {
	for i := 0; i < count; i++ {
		newStack := []string{}
		elem := ""
		for idx, stack := range s.stacks {
			if stack.index == from {
				if len(s.stacks[idx].contents) == 0 {
					continue
				}
				newStack, elem = pop(s.stacks[idx].contents)
				s.stacks[idx].contents = newStack
			}
		}

		for idx, stack := range s.stacks {
			if stack.index == to {
				if elem == "" {
					continue
				}
				result := push(s.stacks[idx].contents, elem)
				s.stacks[idx].contents = result
			}
		}
	}
}

// push new element on stack, creates a new stack with the change,
// todo: change to in memory? (create actual adt?)
func push(stack []string, element string) []string {
	stack = append(stack, element)
	return stack
}

// pop removes the top element of the stack and returns a new stack
// after the operation, and the element that was popped.
func pop(stack []string) ([]string, string) {
	lastElem := stack[len(stack)-1:]
	return stack[:len(stack)-1], lastElem[0]
}

// take Specifies how many elements you want to take, and returns
// 1 - the new stack after the operation
// 2 - the elements taken
func take(stack []string, count int) ([]string, []string) {
    panic("not implemented")
}

// place puts the elems slice on the end of the stack (maintain order)
// and returns the new stack
func place(stack, elems []string) []string {
    panic("not implemented")
}

// getNiceAnswer is just a helper to print out the result, because why not
// And because my board parser made the answer not in order
func getNiceAnswer(s ship) {
    result := ""
	for i := 0; i < len(s.stacks); i++ {
        for ii := 0; ii < len(s.stacks); ii++ {
            if s.stacks[ii].index == i+1 {

                _, el := pop(s.stacks[ii].contents)
                result += el
            }
        }
	}

    result = strings.ReplaceAll(result, "[", "")
    result = strings.ReplaceAll(result, "]", "")
    fmt.Println(result)
}

func main() {
	boardPath := "./board.txt"
	movesPath := "./input.txt"

	ship, err := readBoard(boardPath)
	if err != nil {
		fmt.Println(err)
	}
	moves, err := readFile(movesPath)
	if err != nil {
		fmt.Println(err)
	}


	for _, move := range moves {
		a, f, t := parseInstruction(move)
		executeInstruction(a, f, t, &ship)
	}

    getNiceAnswer(ship)
}
