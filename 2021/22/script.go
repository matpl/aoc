package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const part = 2

type Cuboid struct {
	on bool
	p1 [3]int
	p2 [3]int
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	rebootSteps := []Cuboid{}
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
		splits := strings.Split(lines[index], " ")
		cuboid := Cuboid{
			on: splits[0] == "on",
		}

		splits = strings.Split(splits[1], ",")
		x := strings.Split(strings.Split(splits[0], "=")[1], "..")
		y := strings.Split(strings.Split(splits[1], "=")[1], "..")
		z := strings.Split(strings.Split(splits[2], "=")[1], "..")
		x1, _ := strconv.Atoi(x[0])
		x2, _ := strconv.Atoi(x[1])
		y1, _ := strconv.Atoi(y[0])
		y2, _ := strconv.Atoi(y[1])
		z1, _ := strconv.Atoi(z[0])
		z2, _ := strconv.Atoi(z[1])
		cuboid.p1 = [3]int{x1, y1, z1}
		cuboid.p2 = [3]int{x2 + 1, y2 + 1, z2 + 1}

		rebootSteps = append(rebootSteps, cuboid)
	}

	cubes := []Cuboid{}
	for i, step := range rebootSteps {
		if part == 1 && (step.p1[0] < -50 || step.p2[0] > 51 || step.p1[1] < -50 || step.p2[1] > 51 || step.p1[2] < -50 || step.p2[2] > 51) {
			continue
		}
		if i == 0 {
			cubes = append(cubes, step)
		} else {
			newCubes := []Cuboid{}
			for _, cube := range cubes {
				if intersects(cube, step) {
					intersection := getIntersection(cube, step)
					newCubes = append(newCubes, remove(cube, intersection)...)
				} else {
					newCubes = append(newCubes, cube)
				}
			}
			if step.on {
				newCubes = append(newCubes, step)
			}
			cubes = newCubes
		}
	}

	fmt.Println(getVolume(cubes))
}

func intersects(c1 Cuboid, c2 Cuboid) bool {

	intersection := getIntersection(c1, c2)

	return intersection.p2[0] > intersection.p1[0] &&
		intersection.p2[1] > intersection.p1[1] &&
		intersection.p2[2] > intersection.p1[2]
}

func remove(cuboid Cuboid, toRemove Cuboid) []Cuboid {

	cuboids := []Cuboid{cuboid}
	for i := 0; i < 3; i++ {
		newCuboids := []Cuboid{}
		for _, c := range cuboids {
			var newCuboid Cuboid
			if toRemove.p1[i] > c.p1[i] {
				if i == 0 {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], c.p1[1], c.p1[2]},
						p2: [3]int{toRemove.p1[0], c.p2[1], c.p2[2]},
					}

				} else if i == 1 {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], c.p1[1], c.p1[2]},
						p2: [3]int{c.p2[0], toRemove.p1[1], c.p2[2]},
					}
				} else {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], c.p1[1], c.p1[2]},
						p2: [3]int{c.p2[0], c.p2[1], toRemove.p1[2]},
					}
				}

				if !same(newCuboid, toRemove) {
					newCuboids = append(newCuboids, newCuboid)
				}
			}

			if toRemove.p1[i] >= c.p1[i] && toRemove.p2[i] <= c.p2[i] {
				if i == 0 {
					newCuboid = Cuboid{
						p1: [3]int{toRemove.p1[0], c.p1[1], c.p1[2]},
						p2: [3]int{toRemove.p2[0], c.p2[1], c.p2[2]},
					}
				} else if i == 1 {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], toRemove.p1[1], c.p1[2]},
						p2: [3]int{c.p2[0], toRemove.p2[1], c.p2[2]},
					}
				} else {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], c.p1[1], toRemove.p1[2]},
						p2: [3]int{c.p2[0], c.p2[1], toRemove.p2[2]},
					}
				}

				if !same(newCuboid, toRemove) {
					newCuboids = append(newCuboids, newCuboid)
				}
			}

			if toRemove.p2[i] < c.p2[i] {
				if i == 0 {
					newCuboid = Cuboid{
						p1: [3]int{toRemove.p2[0], c.p1[1], c.p1[2]},
						p2: [3]int{c.p2[0], c.p2[1], c.p2[2]},
					}
				} else if i == 1 {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], toRemove.p2[1], c.p1[2]},
						p2: [3]int{c.p2[0], c.p2[1], c.p2[2]},
					}
				} else {
					newCuboid = Cuboid{
						p1: [3]int{c.p1[0], c.p1[1], toRemove.p2[2]},
						p2: [3]int{c.p2[0], c.p2[1], c.p2[2]},
					}
				}

				if !same(newCuboid, toRemove) {
					newCuboids = append(newCuboids, newCuboid)
				}
			}
		}
		cuboids = newCuboids
	}
	return cuboids
}

func getIntersection(c1 Cuboid, c2 Cuboid) Cuboid {
	p1 := [3]int{}
	p2 := [3]int{}
	for i := 0; i < 3; i++ {
		if c1.p1[i] > c2.p1[i] {
			p1[i] = c1.p1[i]
		} else {
			p1[i] = c2.p1[i]
		}
		if c1.p2[i] < c2.p2[i] {
			p2[i] = c1.p2[i]
		} else {
			p2[i] = c2.p2[i]
		}
	}
	return Cuboid{
		p1: p1,
		p2: p2,
	}
}

func same(c1 Cuboid, c2 Cuboid) bool {
	for i := 0; i < 3; i++ {
		if !(c1.p1[i] == c2.p1[i] && c1.p2[i] == c2.p2[i]) {
			return false
		}
	}
	return true
}

func getVolume(cuboids []Cuboid) int64 {
	var volume int64
	volume = 0
	for _, c := range cuboids {
		volume += int64((c.p2[0] - c.p1[0]) * (c.p2[1] - c.p1[1]) * (c.p2[2] - c.p1[2]))
	}

	return volume
}
