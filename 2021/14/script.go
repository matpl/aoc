package main

import (
	"fmt"
	"os"
	"strings"
)

const part = 2

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	instructions := make(map[string]rune)
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
		if index >= 2 {
			splits := strings.Split(lines[index], " -> ")
			instructions[splits[0]] = []rune(splits[1])[0]
		}
	}

	counters := make(map[string]int)

	lettersCounts := make(map[string]int)
	for i := 0; i < len(lines[0])-1; i++ {
		lettersCounts[string(lines[0][i])] = lettersCounts[string(lines[0][i])] + 1
		counters[string(lines[0][i])+string(lines[0][i+1])] = counters[string(lines[0][i])+string(lines[0][i+1])] + 1
	}
	lettersCounts[string(lines[0][len(lines[0])-1])] = lettersCounts[string(lines[0][len(lines[0])-1])] + 1

	steps := 10
	if part == 2 {
		steps = 40
	}

	counters2 := make(map[string]int)
	for j := 0; j < steps; j++ {

		for k, v := range counters {
			counters2[k] = v
		}

		for k, count := range counters2 {
			v, ok := instructions[k]
			if ok && count != 0 {
				lettersCounts[string(v)] += counters2[k]
				counters[string(k[0])+string(v)] += counters2[k]
				counters[string(v)+string(k[1])] += counters2[k]
				counters[k] -= counters2[k]
			}
		}
	}

	maxCount := 0
	minCount := -1
	for _, v := range lettersCounts {
		if minCount == -1 || v < minCount {
			minCount = v
		}
		if v > maxCount {
			maxCount = v
		}
	}
	fmt.Println(maxCount - minCount)
}
