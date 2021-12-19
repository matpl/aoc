package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const part = 2

var grid [][]int
var risks [][]int
var unvisited map[string]struct{} = make(map[string]struct{})

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	grid = [][]int{}
	risks = [][]int{}

	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")

		risks = append(risks, []int{})
		grid = append(grid, []int{})
		for i := 0; i < len(lines[index]); i++ {
			risks[len(risks)-1] = append(risks[len(risks)-1], -1)
			grid[len(grid)-1] = append(grid[len(grid)-1], int(lines[index][i])-48)
		}
	}

	if part == 2 {
		length := len(grid[0])
		for r := 0; r < len(grid); r++ {
			for steps := 0; steps < 4; steps++ {
				for c := 0; c < length; c++ {
					grid[r] = append(grid[r], (grid[r][c]+1+steps-1)%9+1)
					risks[r] = append(risks[r], -1)
				}
			}
		}
		length = len(grid)
		for steps := 0; steps < 4; steps++ {
			for r := 0; r < length; r++ {
				grid = append(grid, []int{})
				risks = append(risks, []int{})
				for c := 0; c < len(grid[0]); c++ {
					grid[len(grid)-1] = append(grid[len(grid)-1], (grid[r][c]+1+steps-1)%9+1)
					risks[len(risks)-1] = append(risks[len(risks)-1], -1)
				}
			}
		}
	}

	navigate(0, 0, 0)
	fmt.Println(risks[len(risks)-1][len(risks[len(risks)-1])-1])
}

func navigate(r int, c int, totalRisk int) {
	delete(unvisited, fmt.Sprint(r)+"-"+fmt.Sprint(c))
	risks[r][c] = totalRisk

	if r > 0 && risks[r-1][c] == -1 {
		unvisited[fmt.Sprint(r-1)+"-"+fmt.Sprint(c)] = struct{}{}
	}
	if r < len(risks)-1 && risks[r+1][c] == -1 {
		unvisited[fmt.Sprint(r+1)+"-"+fmt.Sprint(c)] = struct{}{}
	}
	if c > 0 && risks[r][c-1] == -1 {
		unvisited[fmt.Sprint(r)+"-"+fmt.Sprint(c-1)] = struct{}{}
	}
	if c < len(risks[r])-1 && risks[r][c+1] == -1 {
		unvisited[fmt.Sprint(r)+"-"+fmt.Sprint(c+1)] = struct{}{}
	}

	totalRisk = -1
	lowestR := -1
	lowestC := -1
	for k := range unvisited {
		splits := strings.Split(k, "-")
		r, _ = strconv.Atoi(splits[0])
		c, _ = strconv.Atoi(splits[1])

		if r > 0 && risks[r-1][c] != -1 && (totalRisk == -1 || risks[r-1][c]+grid[r][c] < totalRisk) {
			lowestR = r
			lowestC = c
			totalRisk = risks[r-1][c] + grid[r][c]
		}
		if r < len(risks)-1 && risks[r+1][c] != -1 && (totalRisk == -1 || risks[r+1][c]+grid[r][c] < totalRisk) {
			lowestR = r
			lowestC = c
			totalRisk = risks[r+1][c] + grid[r][c]
		}
		if c > 0 && risks[r][c-1] != -1 && (totalRisk == -1 || risks[r][c-1]+grid[r][c] < totalRisk) {
			lowestR = r
			lowestC = c
			totalRisk = risks[r][c-1] + grid[r][c]
		}
		if c < len(risks[r])-1 && risks[r][c+1] != -1 && (totalRisk == -1 || risks[r][c+1]+grid[r][c] < totalRisk) {
			lowestR = r
			lowestC = c
			totalRisk = risks[r][c+1] + grid[r][c]
		}
	}
	if totalRisk != -1 {
		if risks[lowestR][lowestC] == -1 {
			navigate(lowestR, lowestC, totalRisk)
		}
	}
}
