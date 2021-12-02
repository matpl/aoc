package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const part = 2

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	numbers := []int{}

	increment := 1

	if part == 2 {
		increment = 3
	}

	increased := 0
	for i := 0; i < len(lines); i += 1 {
		number, _ := strconv.Atoi(strings.Trim(lines[i], "\r\n"))
		numbers = append(numbers, number)

		if i > increment-1 {
			if sum(numbers, i, increment) > sum(numbers, i-1, increment) {
				increased++
			}
		}
	}
	fmt.Println(increased)
}

func sum(numbers []int, index int, count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += numbers[index-i]
	}
	return sum
}
