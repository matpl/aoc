package main

import (
	"fmt"
	"os"
	"strings"
)

const part = 2

var graph map[string][]string

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	graph = make(map[string][]string)

	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
		splits := strings.Split(lines[index], "-")

		if splits[1] != "start" {
			if graph[splits[0]] == nil {
				graph[splits[0]] = []string{}
			}
			graph[splits[0]] = append(graph[splits[0]], splits[1])
		}

		if splits[0] != "start" {
			if graph[splits[1]] == nil {
				graph[splits[1]] = []string{}
			}
			graph[splits[1]] = append(graph[splits[1]], splits[0])
		}
	}

	visited := make(map[string]int)
	fmt.Println(navigate("start", &visited, true))

}

func navigate(node string, visited *map[string]int, canVisitSmallCaves bool) int {
	if node == "end" {
		return 1
	}

	if (*visited)[node] == 0 {
		(*visited)[node] = 1
	} else {
		(*visited)[node] = (*visited)[node] + 1
	}

	if node[0] >= 97 {
		if part == 1 {
			if (*visited)[node] > 1 {
				(*visited)[node] = (*visited)[node] - 1
				return 0
			}
		} else if (*visited)[node] > 1 {

			if !canVisitSmallCaves {
				(*visited)[node] = (*visited)[node] - 1
				return 0
			} else if (*visited)[node] == 2 {
				canVisitSmallCaves = false
			}
		}
	}

	total := 0
	for _, v := range graph[node] {
		total += navigate(v, visited, canVisitSmallCaves)
	}

	if (*visited)[node] != 0 {
		(*visited)[node] = (*visited)[node] - 1
	}

	return total
}
