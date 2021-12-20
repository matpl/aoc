package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const part = 2

type Node struct {
	value    int
	children []*Node
	parent   *Node
	isPair   bool
}

var flatValueNodes []*Node

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	if part == 1 {
		fmt.Println(addition(lines))
	} else {
		maxMagnitude := 0
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines); j++ {
				if i != j {
					magnitude := addition([]string{lines[i], lines[j]})
					if magnitude > maxMagnitude {
						maxMagnitude = magnitude
					}
				}
			}
		}
		fmt.Println(maxMagnitude)
	}
}

func addition(lines []string) int {
	var root *Node
	var current *Node = nil
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
		if index != 1 {
			if root != nil {
				root = &Node{
					children: []*Node{root},
					value:    -1,
				}
				root.children[0].parent = root
			} else {
				root = &Node{
					children: []*Node{},
					value:    -1,
				}
			}
		}
		current = root

		for _, c := range lines[index] {
			if c == '[' {
				node := &Node{
					children: []*Node{},
					parent:   current,
					value:    -1,
				}

				if current != nil {
					current.children = append(current.children, node)
				}
				current = node
			} else if c == ']' {
				current = current.parent
			} else if c != ',' {
				node := &Node{
					children: nil,
					parent:   current,
					value:    int(c) - 48,
				}

				current.children = append(current.children, node)

				setIsPair(current)
			}
		}

		if index > 0 {
			for {
				flatValueNodes = []*Node{}
				flatten(root)
				if !explode(root, 1) && !split(root) {
					break
				}
			}
		}
	}

	flatValueNodes = []*Node{}
	flatten(root)
	return getMagnitude(root)
}

func getMagnitude(node *Node) int {
	if len(node.children) == 0 {
		return node.value
	} else {
		if len(node.children) == 1 {
			return getMagnitude(node.children[0])
		} else {
			return getMagnitude(node.children[0])*3 + getMagnitude(node.children[1])*2
		}
	}
}

func explode(node *Node, level int) bool {
	if level > 4 && node.isPair {

		for i, flatNode := range flatValueNodes {
			if node.children[0] == flatNode {
				if i != 0 {
					flatValueNodes[i-1].value += node.children[0].value
				}
			} else if node.children[1] == flatNode {
				if i != len(flatValueNodes)-1 {
					flatValueNodes[i+1].value += node.children[1].value
				}
			}
		}
		node.children = nil
		node.value = 0
		setIsPair(node.parent)
		setIsPair(node)
		return true
	} else {
		for _, child := range node.children {
			if explode(child, level+1) {
				return true
			}
		}
	}
	return false
}

func split(node *Node) bool {
	if node.value >= 10 {
		node.children = []*Node{}
		node.children = append(node.children, &Node{
			children: nil,
			parent:   node,
			value:    int(math.Floor(float64(node.value) / 2.0)),
		})
		node.children = append(node.children, &Node{
			children: nil,
			parent:   node,
			value:    int(math.Ceil(float64(node.value) / 2.0)),
		})
		node.value = -1
		setIsPair(node)
		setIsPair(node.parent)
		return true
	} else {
		for _, child := range node.children {
			if split(child) {
				return true
			}
		}
	}
	return false
}
func setIsPair(node *Node) {
	if len(node.children) == 2 && node.children[0].value != -1 && node.children[1].value != -1 {
		node.isPair = true
	} else {
		node.isPair = false
	}
}

func flatten(node *Node) {
	if node.value != -1 {
		flatValueNodes = append(flatValueNodes, node)
	} else {
		for _, child := range node.children {
			flatten(child)
		}
	}
}
