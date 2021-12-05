package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	m := make(map[string]int)

	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")

		points := strings.Split(lines[index], " -> ")
		p1 := strings.Split(points[0], ",")
		p2 := strings.Split(points[1], ",")

		x1, _ := strconv.Atoi(p1[0])
		y1, _ := strconv.Atoi(p1[1])
		x2, _ := strconv.Atoi(p2[0])
		y2, _ := strconv.Atoi(p2[1])

		xSign := 1
		ySign := 1

		if y1 > y2 {
			ySign = -1
		} else if y1 == y2 {
			ySign = 0
		}

		if x1 > x2 {
			xSign = -1
		} else if x1 == x2 {
			xSign = 0
		}

		for i := 0; x1+i*xSign != x2+1*xSign || y1+i*ySign != y2+1*ySign; i++ {
			key := strconv.Itoa(x1+i*xSign) + "," + strconv.Itoa(y1+i*ySign)
			if _, ok := m[key]; ok {
				m[key] = m[key] + 1
			} else {
				m[key] = 1
			}
		}
	}

	total := 0
	for _, v := range m {
		if v >= 2 {
			total++
		}
	}
	println(total)
}
