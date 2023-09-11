package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	nodeType string
	size     int
	children []*Node
	previous *Node
}

func (n *Node) AddChild(node Node) {
	n.children = append(n.children, &node)
}

func (n *Node) FindChild(name string) *Node {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	// Fails when you try to cd to a dir that doesn't exist
	panic("not implemented")
}

func (n *Node) GetNodeSize() int {
	result := 0
	for _, c := range n.children {
		result += c.size
		result += c.GetNodeSize()
	}

	return result
}

func (n Node) PrintNode() {
	fmt.Println(n.name, "->")
	for _, node := range n.children {
		node.PrintNode()
	}
}

type tokenType int

const (
	Instruction tokenType = iota
	File
	Directory
)

type sizeCheck struct {
	dir  string
	size int
}

// parseToken takes in the tokens, and returns the type.
// note, should this be replaced with a regex parser if something more complex
// came up?
func parseToken(tokens []string) tokenType {
	if len(tokens) == 2 {
		f := tokens[0]
		if f == "dir" {
			return Directory
		} else if f == "$" {
			return Instruction
		} else {
			return File
		}
	} else {
		return Instruction
	}
}

// readFile takes a path and returns a string slice.
func readFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(data), "\n")
	return temp[:len(temp)-1], nil
}

// linearWalk takes a slice of strings as instructions and
// iterates over them to build up a file tree structure.
func linearWalk(instructions []string) *Node {
	root := Node{
		name:     "/",
		nodeType: "dir",
		size:     0,
		children: []*Node{},
		previous: nil,
	}
	current := &root

	for _, instruction := range instructions {
		tokens := strings.Split(instruction, " ")
		token := parseToken(tokens)

		switch token {
		case Instruction:
			instruct := tokens[1]
			if instruct == "cd" {
				path := tokens[2]
				if path == ".." {
					if current.previous == nil {
						continue
					}
					current = current.previous
				} else if path == "/" {
					continue
				} else {
					previous := current
					// will break if cd to non-existing dir
					current = current.FindChild(path)
					current.previous = previous
				}
			} else if instruct == "ls" {
				// skipping ls for now
				continue
			}
		case File:
			size, _ := strconv.Atoi(tokens[0])
			fileNode := Node{
				name:     tokens[1],
				nodeType: "file",
				size:     size,
				children: []*Node{},
				previous: current.previous,
			}
			current.AddChild(fileNode)
		case Directory:
			path := tokens[1]
			dirNode := Node{
				name:     path,
				nodeType: "dir",
				size:     0,
				children: []*Node{},
				previous: current.previous,
			}
			current.AddChild(dirNode)
		}
	}

	return &root
}

func getDirSize(n *Node) sizeCheck {
	// Assuming dirs are passed here
	c := sizeCheck{
		dir:  n.name,
		size: n.GetNodeSize(),
	}

	return c
}

func allDirs(n *Node, result []sizeCheck) []sizeCheck {
	if n.nodeType == "dir" {
		for _, c := range n.children {
			if c.nodeType == "dir" {
				// fmt.Println(c.name, c.GetNodeSize())

				check := sizeCheck{
					dir:  c.name,
					size: c.GetNodeSize(),
				}

				// Don't understand this
				result = append(result, check)
				result = allDirs(c, result)
			}
		}
	}

	return result
}

func filterLower(d []sizeCheck, filterAmount int) []int {
	result := []int{}

	for _, e := range d {
		if e.size < filterAmount {
			result = append(result, e.size)
		}
	}

	return result
}

func sum(d []int) int {
	result := 0
	for _, e := range d {
		result += e
	}

	return result
}

func main() {
	fmt.Println("starting")
	inputPath := "./input.txt"

	instructions, err := readFile(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	root := linearWalk(instructions)
	// root.PrintNode()
	// fmt.Println(root)
	r := allDirs(root, []sizeCheck{})
	// fmt.Println(r)

	checkAmount := 100000
	filteredDirSizes := filterLower(r, checkAmount)
    fmt.Println("Sum:", sum(filteredDirSizes))
}
