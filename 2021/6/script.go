package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const totalDays = 256

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	for index, line := range lines {
		lines[index] = strings.Trim(line, "\r\n")
	}

	perTimer := [9]int{}
	newDay := [9]int{}

	for i := 0; i < 9; i++ {
		perTimer[i] = 0
		newDay[i] = 0
	}

	splitted := strings.Split(lines[0], ",")
	for _, str := range splitted {
		no, _ := strconv.Atoi(str)
		perTimer[no] = perTimer[no] + 1
	}

	for days := 0; days < totalDays; days++ {
		for i := len(perTimer) - 1; i >= 0; i-- {
			if i != 0 {
				newDay[i-1] = perTimer[i]
			} else {
				newDay[8] = perTimer[i]
				newDay[6] = newDay[6] + perTimer[i]
			}
		}
		perTimer = newDay
	}

	total := 0
	for i := 0; i < len(newDay); i++ {
		total += newDay[i]
	}

	fmt.Println(total)
}
