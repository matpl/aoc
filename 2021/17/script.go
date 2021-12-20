package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const part = 2

func main() {
	data, _ := os.ReadFile("input.txt")
	splits := strings.Split(strings.Trim(strings.Split(string(data), ": ")[1], "\r\n"), ", ")
	splits[0] = strings.TrimLeft(splits[0], "x=")
	splits[1] = strings.TrimLeft(splits[1], "y=")
	xs := strings.Split(splits[0], "..")
	ys := strings.Split(splits[1], "..")

	minX, _ := strconv.Atoi(xs[0])
	maxX, _ := strconv.Atoi(xs[1])
	minY, _ := strconv.Atoi(ys[0])
	maxY, _ := strconv.Atoi(ys[1])

	minVx := 0
	maxVx := maxX
	minT := 0

	for v := 0; minVx == 0; v++ {
		for t := 0; ; t++ {
			x := int((float32(v) - (float32(t)-1.0)/2.0) * float32(t))
			if x >= minX && x <= maxX {
				minVx = v
				minT = t
				break
			} else if x > maxX || x < 0 {
				break
			}
		}
	}

	minVy := minY
	maxVy := 0
	maxHeight := minY
	bestMaxHeight := 0
	for v := 0; v < int(math.Abs(float64(minY))); v++ {
		y := 0
		maxHeight = 0
		for t := minT; ; t++ {
			y = int((float32(v) - (float32(t)-1.0)/2.0) * float32(t))
			if y > maxHeight {
				maxHeight = y
			}
			if y >= minY && y <= maxY {
				if maxHeight > bestMaxHeight {
					bestMaxHeight = maxHeight
				}
				maxVy = v
				break
			} else if y < minY {
				break
			}
		}
	}

	total := 0
	for vx := minVx; vx <= maxVx; vx++ {
		for vy := minVy; vy <= maxVy; vy++ {
			highestX := 0
			for t := 0; ; t++ {
				x := int((float32(vx) - (float32(t)-1.0)/2.0) * float32(t))
				y := int((float32(vy) - (float32(t)-1.0)/2.0) * float32(t))
				if x > highestX {
					highestX = x
				} else {
					x = highestX
				}
				if x >= minX && x <= maxX && y >= minY && y <= maxY {
					total++
					break
				} else if y < minY {
					break
				}
			}
		}
	}

	if part == 1 {
		fmt.Println(bestMaxHeight)
	} else {
		fmt.Println(total)
	}
}
