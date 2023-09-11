package main

// Read in data
// split each line in half
// find item(s) that exists in both halfs
// Assign priority to those values
// sum those values
// return sum

import (
	"fmt"
	"os"
	"strings"
)

// readData takes in a path and reads the input file at that location
func readData(path string) ([]string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    return strings.Split(string(data), "\n"), nil
}

// splitLineInTwo takes a string and returns two equal parts 
// by splitting it into two
func splitLineInTwo(line string) (string, string) {
    return line[0:len(line)/2], line[len(line)/2:]
}

// intersection is similar to the set operation of an intersection which
// returns the elements that are in both set a, and b.
func intersection(a, b string) []string {
    result := []string{}
    for _, left := range a {
        for _, right := range b {
            if left == right {
                result = append(result, string(left))
                break
            }
        }
    }

    return result
}

// distinct returns a set where each element on exists once in a slice
func distinct(set []string) []string {
    result := []string{}
    for _, elem := range set {
        add := true
        for _, in := range result {
            if in == elem {
                add = false
            }
        }
        if add {
            result = append(result, elem)
        }
    }

    return result
}

// countPriority assigns a priority and returns the value
func countPriority(data []string) int {
    result := 0
    standing := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for _, elem := range data {
        for i, stand := range standing {
            if elem == string(stand) {
                // Zero based indexing
                result += i+1
                break
            }
        }
    }

    return result
}

func main() {
    path := "./input.txt"
    lines, err := readData(path)
    if err != nil {
        panic(err)
    }

    result := 0

    for _, line := range lines {
        a, b := splitLineInTwo(line)
        same := intersection(a, b)
        distinctSet := distinct(same)

        priority := countPriority(distinctSet)
        result += priority
    }

    fmt.Println("result:", result)
}

