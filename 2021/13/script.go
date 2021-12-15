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
	coordinates := [][]int{}
	folds := [][]string{}
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
		if strings.Contains(lines[index], "fold") {
			splits := strings.Split(lines[index][11:], "=")
			folds = append(folds, []string{splits[0], splits[1]})
		} else if lines[index] != "" {
			splits := strings.Split(lines[index], ",")
			n0, _ := strconv.Atoi(splits[0])
			n1, _ := strconv.Atoi(splits[1])
			coordinates = append(coordinates, []int{n0, n1})
		}
	}

	for _, v := range folds {
		for i := 0; i < len(coordinates); i++ {
			n, _ := strconv.Atoi(v[1])
			switch v[0] {
			case "x":
				if coordinates[i][0] > n {
					coordinates[i][0] = n - (coordinates[i][0] - n)
				}
			case "y":
				if coordinates[i][1] > n {
					coordinates[i][1] = n - (coordinates[i][1] - n)
				}
			}
		}
		if part == 1 {
			break
		}
	}

	set := make(map[string]struct{})

	maxR := 0
	maxC := 0
	for _, v := range coordinates {
		if v[0] > maxR {
			maxR = v[0]
		}
		if v[1] > maxC {
			maxC = v[1]
		}
		val := fmt.Sprint(v[0]) + "-" + fmt.Sprint(v[1])
		set[val] = struct{}{}
	}

	if part == 1 {
		fmt.Println(len(set))
	} else {
		for c := 0; c <= maxC; c++ {
			for r := 0; r <= maxR; r++ {
				_, ok := set[fmt.Sprint(r)+"-"+fmt.Sprint(c)]
				if ok {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
	}
}
