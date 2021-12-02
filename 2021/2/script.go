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

	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		splits := strings.Split(strings.Trim(line, "\r\n"), " ")
		number, _ := strconv.Atoi(splits[1])
		if part == 1 {
			switch splits[0] {
			case "forward":
				horizontal += number
			case "down":
				depth += number
			case "up":
				depth -= number
			}
		} else {
			switch splits[0] {
			case "forward":
				horizontal += number
				depth += aim * number
			case "down":
				aim += number
			case "up":
				aim -= number
			}
		}
	}

	fmt.Println(horizontal * depth)
}

func sum(numbers []int, index int, count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += numbers[index-i]
	}
	return sum
}
