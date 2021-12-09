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

	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
	}
	min := -1
	max := 0
	splits := strings.Split(lines[0], ",")
	numbers := []int{}
	for _, s := range splits {
		number, _ := strconv.Atoi(s)
		numbers = append(numbers, number)

		if min == -1 || number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	variation := (max - min) / 2
	min = -1

	fuelCount := 0
	for i := 0; i < variation; i++ {
		fuelCount = 0
		for j := 0; j < len(numbers); j++ {
			if numbers[j] < i {
				if part == 1 {
					fuelCount += i - numbers[j]
				} else {
					fuelCount += (i - numbers[j] + 1) * (i - numbers[j]) / 2.0
				}
			} else if numbers[j] > i {
				if part == 1 {
					fuelCount += numbers[j] - i
				} else {
					fuelCount += (numbers[j] - i + 1) * (numbers[j] - i) / 2.0
				}
			}
		}

		if min == -1 || fuelCount < min {
			min = fuelCount
		}
	}

	fmt.Println(min)
}
