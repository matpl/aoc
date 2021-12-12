package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const part = 2

var octopi [][]int

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	octopi = [][]int{}
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
		octopi = append(octopi, []int{})
		for _, c := range lines[index] {
			n, _ := strconv.Atoi(string(c))
			octopi[index] = append(octopi[index], n)
		}
	}

	flashes := 0

	steps := 100
	if part == 2 {
		steps = -1
	}

	for i := 0; i != steps; i++ {
		for r := 0; r < len(octopi); r++ {
			for c := 0; c < len(octopi[r]); c++ {
				flashes += increment(r, c)
			}
		}

		allFlashed := true
		for r := 0; r < len(octopi); r++ {
			for c := 0; c < len(octopi[r]); c++ {
				if octopi[r][c] > 9 {
					octopi[r][c] = 0
				} else {
					allFlashed = false
				}
			}
		}

		if part == 2 && allFlashed {
			fmt.Println(i + 1)
			break
		}
	}

	if part == 1 {
		fmt.Println(flashes)
	}
}

func increment(r int, c int) int {
	if r < 0 || r >= len(octopi) || c < 0 || c >= len(octopi[r]) {
		return 0
	}
	flashes := 0
	octopi[r][c] = octopi[r][c] + 1
	if octopi[r][c] == 10 {
		flashes = 1
		flashes += increment(r-1, c-1)
		flashes += increment(r-1, c)
		flashes += increment(r-1, c+1)
		flashes += increment(r, c-1)
		flashes += increment(r, c)
		flashes += increment(r, c+1)
		flashes += increment(r+1, c-1)
		flashes += increment(r+1, c)
		flashes += increment(r+1, c+1)
	}

	return flashes
}
