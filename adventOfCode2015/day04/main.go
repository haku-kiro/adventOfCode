package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoadSecret(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data)[:len(string(data))-1], nil
}

func PassHash(str string) string {
	md5 := md5.New()
	io.WriteString(md5, str)

	return fmt.Sprintf("%x", md5.Sum(nil))
}

func FindKey(input, checkText string) int {
	i := 0
	for {
		in := fmt.Sprintf("%s%d", input, i)
		hash := PassHash(in)
		check := strings.Index(hash, checkText) == 0
		if check {
			return i
		}

		i++
	}
}

func PartOne(secret string) int {
	return FindKey(secret, "00000")
}

func PartTwo(secret string) int {
	return FindKey(secret, "000000")
}

func main() {
	secretPath := "./input.txt"
	secret, err := LoadSecret(secretPath)
	if err != nil {
		panic(err)
	}

    // Working, but at what cost??
	keyOne := PartOne(secret)
	keyTwo := PartTwo(secret)

	fmt.Println("Key1:", keyOne)
	fmt.Println("Key2:", keyTwo)
}
