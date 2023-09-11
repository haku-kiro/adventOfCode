package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// readFile reads the contents of a file given the path into a slice of strings
func readFile(path string) ([]string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    return strings.Split(string(data), "\n"), nil
}

// transformToSums takes in a slice and sums numbers grouped by spaces
func transformToSums(data []string) ([]int, error) {
    counter := 0
    result := []int{}
    for _, elem := range data {
        num, err := strconv.Atoi(elem)
        if err == nil {
            counter += num
        } else {
            result = append(result, counter)
            counter = 0
        }
    }

    return result, nil
}

// getTopThreeSum given an input, sorts and returns the sum of the highest 3
// should probably have just made the getMax method more generic?
func getTopThreeSum(data []int) int {
    sort.Ints(data)
    lastThree := data[len(data)-3:]

    sum := 0
    for _, elem := range lastThree {
        sum += elem
    }

    return sum
}

// getMax returns the largest number in the slice
func getMax(data []int) int {
    temp := 0
    for _, elem := range data {
        if elem > temp {
            temp = elem
        }
    }

    return temp
}

func main() {
    fmt.Println("Day 1! woo!")
    path := "./input.txt"
    data, err := readFile(path)
    if err != nil {
        panic(err)
    }

    sums, err := transformToSums(data)
    if err != nil {
        panic(err)
    }

    answer := getTopThreeSum(sums)
    fmt.Printf("the answer is: %d", answer)
}

