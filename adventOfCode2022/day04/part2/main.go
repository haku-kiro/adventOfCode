package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	temp := strings.Split(string(data), "\n")
	return temp[:len(temp)-1], nil
}

func getRanges(row string) (string, string) {
	ranges := strings.Split(row, ",")
	return ranges[0], ranges[1]
}

func getLowHigh(r string) (int, int) {
	t := strings.Split(r, "-")
	ls, _ := strconv.Atoi(t[0])
	rs, _ := strconv.Atoi(t[1])
	return ls, rs
}

func checkOutside(rangeA, rangeB string) bool {
	la, ha := getLowHigh(rangeA) // 2 4
	lb, hb := getLowHigh(rangeB) // 6 8

	if la < lb || ha > hb {
		return true
	}

	return false
}

func overlappingRanges(rangeA, rangeB string) bool {
	la, ha := getLowHigh(rangeA) // 2 6
	lb, hb := getLowHigh(rangeB) // 4 8

    if (la >= lb && la <= hb) || (ha >= lb && ha <= hb) {
        return true
    }

    return false
}

func main() {
	path := "./input.txt"
	data, err := readFile(path)
	if err != nil {
		fmt.Println(err)
	}

	result := 0
	for _, elem := range data {
		l, r := getRanges(elem)

        check := overlappingRanges(l, r)
        check2 := overlappingRanges(r, l)

        if check || check2 {
            result += 1
        }
	}

	fmt.Println("amount contained:", result)
}
