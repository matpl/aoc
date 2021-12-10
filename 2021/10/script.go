package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const part = 2

var lines []string

func main() {
	data, _ := os.ReadFile("input.txt")
	lines = strings.Split(string(data), "\n")

	score := 0
	scores := []int{}
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")

		stack := []rune{}
		var i int
		var c rune
		for i, c = range lines[index] {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				stack = append(stack, c)
			} else if c == ')' {
				if stack[len(stack)-1] != '(' {
					score += 3
					break
				}
				stack = stack[:len(stack)-1]
			} else if c == ']' {
				if stack[len(stack)-1] != '[' {
					score += 57
					break
				}
				stack = stack[:len(stack)-1]
			} else if c == '}' {
				if stack[len(stack)-1] != '{' {
					score += 1197
					break
				}
				stack = stack[:len(stack)-1]
			} else if c == '>' {
				if stack[len(stack)-1] != '<' {
					score += 25137
					break
				}
				stack = stack[:len(stack)-1]
			}
		}

		if part == 2 && i == len(lines[index])-1 {
			score = 0
			for j := len(stack) - 1; j >= 0; j-- {
				switch stack[j] {
				case '(':
					score = score*5 + 1
				case '[':
					score = score*5 + 2
				case '{':
					score = score*5 + 3
				case '<':
					score = score*5 + 4
				}
			}
			scores = append(scores, score)
		}
	}

	if part == 1 {
		fmt.Println(score)
	} else {
		sort.Ints(scores)
		fmt.Println(scores[len(scores)/2])
	}
}
