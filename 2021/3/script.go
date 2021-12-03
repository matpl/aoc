package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	lineLength := 0
	for index, line := range lines {
		line = strings.Trim(line, "\r\n")
		lines[index] = line
		if lineLength == 0 {
			lineLength = len(line)
		}
	}

	line := ""
	for i := 0; i < lineLength; i++ {
		line += string([]byte{getMostCommonBit(lines, i)})
	}
	gamma := getDecimal(line)

	fmt.Println(gamma * (^gamma & (1<<lineLength - 1)))

	fmt.Println(getDecimal(getBitCriteria(lines, true)) * getDecimal(getBitCriteria(lines, false)))
}

func getDecimal(line string) int {
	lineLength := len(line)
	decimal := 0
	for i := 0; i < lineLength; i++ {
		if line[i] == '1' {
			decimal = decimal | 1<<(lineLength-i-1)
		}
	}
	return decimal
}

func getBitCriteria(lines []string, wantsMostCommon bool) string {
	lineLength := len(lines[0])
	for i := 0; i < lineLength; i++ {
		mostCommon := getMostCommonBit(lines, i)

		if wantsMostCommon {
			if mostCommon == ' ' {
				mostCommon = '1'
			}
			lines = getRows(lines, i, mostCommon)
		} else {
			if mostCommon == ' ' || mostCommon == '1' {
				mostCommon = '0'
			} else {
				mostCommon = '1'
			}
			lines = getRows(lines, i, mostCommon)
		}

		if len(lines) == 1 {
			return lines[0]
		}
	}
	return ""
}

func getMostCommonBit(lines []string, column int) byte {
	oneCount := 0
	for _, line := range lines {
		if line[column] == '1' {
			oneCount++
		}
	}

	threshold := float32(len(lines)) / 2.0

	if float32(oneCount) > threshold {
		return '1'
	} else if float32(oneCount) < threshold {
		return '0'
	} else {
		return ' '
	}
}

func getRows(lines []string, column int, digit byte) []string {
	rows := []string{}
	for i := 0; i < len(lines); i++ {
		if lines[i][column] == digit {
			rows = append(rows, lines[i])
		}
	}
	return rows
}
