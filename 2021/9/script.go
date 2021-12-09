package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const part = 2

var lines []string

func main() {
	data, _ := os.ReadFile("input.txt")
	lines = strings.Split(string(data), "\n")

	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
	}

	basinSizes := []int{}

	total := 0
	for index, line := range lines {
		for i := 0; i < len(line); i++ {
			isLow := true

			if i != 0 {
				if line[i-1] <= line[i] {
					isLow = false
				}
			}
			if isLow && i != len(line)-1 {
				if line[i+1] <= line[i] {
					isLow = false
				}
			}
			if isLow && index != 0 {
				if lines[index-1][i] <= line[i] {
					isLow = false
				}
			}
			if isLow && index != len(lines)-1 {
				if lines[index+1][i] <= line[i] {
					isLow = false
				}
			}

			if isLow {
				currentNo, _ := strconv.Atoi(string(line[i]))
				total += currentNo + 1
				basinSizes = append(basinSizes, len(getBasinSize(index, i, [][]int{}))+1)
			}
		}
	}

	sort.Ints(basinSizes)
	basinSizes = basinSizes[len(basinSizes)-3:]

	if part == 1 {
		fmt.Println(total)
	} else {
		fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
	}
}

func getBasinSize(r int, c int, basin [][]int) [][]int {
	if r != 0 && lines[r-1][c] != '9' && lines[r-1][c] > lines[r][c] && !visited(r-1, c, basin) {
		basin = append(basin, []int{r - 1, c})
		basin = getBasinSize(r-1, c, basin)
	}
	if r != len(lines)-1 && lines[r+1][c] != '9' && lines[r+1][c] > lines[r][c] && !visited(r+1, c, basin) {
		basin = append(basin, []int{r + 1, c})
		basin = getBasinSize(r+1, c, basin)
	}
	if c != 0 && lines[r][c-1] != '9' && lines[r][c-1] > lines[r][c] && !visited(r, c-1, basin) {
		basin = append(basin, []int{r, c - 1})
		basin = getBasinSize(r, c-1, basin)
	}
	if c != len(lines[r])-1 && lines[r][c+1] != '9' && lines[r][c+1] > lines[r][c] && !visited(r, c+1, basin) {
		basin = append(basin, []int{r, c + 1})
		basin = getBasinSize(r, c+1, basin)
	}

	return basin
}

func visited(r int, c int, basin [][]int) bool {
	for _, v := range basin {
		if v[0] == r && v[1] == c {
			return true
		}
	}
	return false
}
