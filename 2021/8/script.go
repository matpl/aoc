package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const part = 2

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	number := 0
	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")

		splits := strings.Split(lines[index], " | ")
		if part == 1 {
			splits := strings.Split(splits[1], " ")

			for _, v := range splits {
				len := len(v)
				if len == 2 || len == 4 || len == 3 || len == 7 {
					number++
				}
			}
		} else {
			mappings := make(map[string]string)

			digits := strings.Split(splits[0], " ")

			var one string
			var seven string
			var four string
			for _, v := range digits {
				switch len(v) {
				case 2:
					one = v
				case 3:
					seven = v
				case 4:
					four = v
				}
			}
			for i := 0; i < len(one); i++ {
				seven = strings.ReplaceAll(seven, string(one[i]), "")
			}

			mappings["a"] = seven

			charCounts := make(map[int]int)
			for i := 97; i <= 103; i++ {
				count := strings.Count(splits[0], string(i))
				switch count {
				case 4:
					mappings["e"] = string(i)
				case 6:
					mappings["b"] = string(i)
				case 7: // either d or g, d is the one in digit 4
					if strings.Count(four, string(i)) == 1 {
						mappings["d"] = string(i)
					} else {
						mappings["g"] = string(i)
					}
				case 8:
					if mappings["a"] != string(i) {
						mappings["c"] = string(i)
					}
				case 9:
					mappings["f"] = string(i)
				}
				charCounts[i] = strings.Count(splits[0], string(i))
			}

			digits = strings.Split(splits[1], " ")
			for i, digit := range digits {
				switch len(digit) {
				case 2:
					number += int(1 * math.Pow10(len(digits)-i-1))
				case 3:
					number += int(7 * math.Pow10(len(digits)-i-1))
				case 4:
					number += int(4 * math.Pow10(len(digits)-i-1))
				case 5:
					if strings.Count(digit, mappings["e"]) == 1 {
						number += int(2 * math.Pow10(len(digits)-i-1))
					} else if strings.Count(digit, mappings["c"]) == 1 {
						number += int(3 * math.Pow10(len(digits)-i-1))
					} else {
						number += int(5 * math.Pow10(len(digits)-i-1))
					}
				case 6:
					if strings.Count(digit, mappings["d"]) != 0 {
						if strings.Count(digit, mappings["e"]) == 1 {
							number += int(6 * math.Pow10(len(digits)-i-1))
						} else {
							number += int(9 * math.Pow10(len(digits)-i-1))
						}
					}
				case 7:
					number += int(8 * math.Pow10(len(digits)-i-1))
				}
			}
		}
	}
	fmt.Println(number)
}
