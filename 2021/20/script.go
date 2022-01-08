package main

import (
	"fmt"
	"os"
	"strings"
)

const part = 2

var enhancement []int

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
	}

	enhancement = []int{}
	for _, char := range lines[0] {
		if char == '.' {
			enhancement = append(enhancement, 0)
		} else {
			enhancement = append(enhancement, 1)
		}
	}

	rawImage := lines[2:]
	image := [][]int{}
	for _, line := range rawImage {
		image = append(image, []int{})
		for _, char := range line {
			if char == '.' {
				image[len(image)-1] = append(image[len(image)-1], 0)
			} else {
				image[len(image)-1] = append(image[len(image)-1], 1)
			}
		}
	}

	iterations := 2
	if part == 2 {
		iterations = 50
	}

	var count int
	for i := 0; i < iterations; i++ {
		fallback := ((i % 2) - 1) * -1
		expandedImage := expand(image, fallback)
		transformedImage := transform(image, expandedImage, fallback)
		count = litCount(transformedImage)
		image = transformedImage
	}
	fmt.Println(count)
}

func expand(image [][]int, fallback int) [][]int {
	expandedImage := [][]int{}
	expandedImage = append(expandedImage, []int{})
	for i := 0; i < len(image[0])+2; i++ {
		expandedImage[len(expandedImage)-1] = append(expandedImage[len(expandedImage)-1], fallback)
	}

	for _, row := range image {
		expandedImage = append(expandedImage, []int{})
		expandedImage[len(expandedImage)-1] = append(expandedImage[len(expandedImage)-1], fallback)
		for _, char := range row {
			expandedImage[len(expandedImage)-1] = append(expandedImage[len(expandedImage)-1], char)
		}
		expandedImage[len(expandedImage)-1] = append(expandedImage[len(expandedImage)-1], fallback)
	}

	expandedImage = append(expandedImage, []int{})
	for i := 0; i < len(image[0])+2; i++ {
		expandedImage[len(expandedImage)-1] = append(expandedImage[len(expandedImage)-1], fallback)
	}

	return expandedImage
}

func transform(image [][]int, expandedImage [][]int, fallback int) [][]int {
	for row := 0; row < len(expandedImage); row++ {
		for col := 0; col < len(expandedImage[row]); col++ {
			number := 0
			for i := -2; i <= 0; i++ {
				for j := -2; j <= 0; j++ {
					outOfBounds := (row+i < 0 || row+i >= len(image) || col+j < 0 || col+j >= len(image[0]))
					if outOfBounds && fallback == 0 || !outOfBounds && image[row+i][col+j] == 1 {
						number |= 1 << (8 - ((i+2)*3 + (j + 2)))
					}
				}
			}
			expandedImage[row][col] = enhancement[number]
		}
	}

	return expandedImage
}

func litCount(image [][]int) int {
	count := 0
	for _, row := range image {
		for _, pixel := range row {
			if pixel == 1 {
				count++
			}
		}
	}
	return count
}
