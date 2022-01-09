package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const part = 2

var enhancement []int

var p1Wins int64
var p2Wins int64

var possibleRollsCounts map[int]int64

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	pos1, _ := strconv.Atoi(strings.Split(strings.Trim(lines[0], "\r\n"), ": ")[1])
	pos2, _ := strconv.Atoi(strings.Split(strings.Trim(lines[1], "\r\n"), ": ")[1])
	pos1--
	pos2--

	p1Score := 0
	p2Score := 0

	if part == 1 {
		for i := 0; ; i++ {
			increment := 3 * ((i * 3) + 2)
			if i%2 == 0 {
				pos1 = (pos1 + increment) % 10
				p1Score += pos1 + 1
				if p1Score >= 1000 {
					fmt.Println((i + 1) * 3 * p2Score)
					break
				}
			} else {
				pos2 = (pos2 + increment) % 10
				p2Score += pos2 + 1
				if p2Score >= 1000 {
					fmt.Println((i + 1) * 3 * p1Score)
					break
				}
			}
		}
	} else {
		p1Wins = 0
		p2Wins = 0
		possibleRollsCounts = make(map[int]int64)
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					roll := i + j + k
					v, ok := possibleRollsCounts[roll]
					if !ok {
						possibleRollsCounts[roll] = 1
					} else {
						possibleRollsCounts[roll] = v + 1
					}
				}
			}
		}
		for i := 3; i <= 9; i++ {
			roll(i, 0, p1Score, p2Score, pos1, pos2, possibleRollsCounts[i])
		}
		if p1Wins > p2Wins {
			fmt.Println(p1Wins)
		} else {
			fmt.Println(p2Wins)
		}
	}
}

func roll(number int, player int, p1Score int, p2Score int, p1Pos int, p2Pos int, factor int64) {

	if player == 0 {
		p1Pos = (p1Pos + number) % 10
		p1Score += p1Pos + 1

		if p1Score >= 21 {
			p1Wins = p1Wins + factor
			return
		}
	} else {
		p2Pos = (p2Pos + number) % 10
		p2Score += p2Pos + 1

		if p2Score >= 21 {
			p2Wins = p2Wins + factor
			return
		}
	}
	player = (player - 1) * -1

	for i := 3; i <= 9; i++ {
		roll(i, player, p1Score, p2Score, p1Pos, p2Pos, factor*possibleRollsCounts[i])
	}
}
