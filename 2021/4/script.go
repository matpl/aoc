package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	numbers := []string{}
	boards := []*[25]string{}
	arrayIndex := 0
	for index, line := range lines {
		line = strings.Trim(line, "\r\n")
		if index == 0 {
			numbers = strings.Split(line, ",")
		} else {
			if len(line) == 0 {
				arrayIndex = 0
				var array [25]string
				boards = append(boards, &array)
			} else {
				row := strings.Split(line, " ")
				for _, number := range row {
					if len(number) != 0 {
						boards[len(boards)-1][arrayIndex] = strings.Trim(number, " ")
						arrayIndex++
					}
				}
			}
		}
	}

	for _, number := range numbers {
		for i := 0; i < len(boards); i++ {
			markBoard(boards[i], number)
			if bingo(boards[i]) {
				sum := 0
				for _, str := range boards[i] {
					boardNo, _ := strconv.Atoi(str)
					sum += boardNo
				}
				no, _ := strconv.Atoi(number)
				fmt.Println(sum * no)

				boards = append(boards[:i], boards[i+1:]...)
				i--
			}
		}
	}
}

func markBoard(board *[25]string, number string) {
	for i := 0; i < 25; i++ {
		if board[i] == number {
			board[i] = "x"
			return
		}
	}
}

func bingo(board *[25]string) bool {
	for r := 0; r < 5; r++ {
		bingoR := true
		bingoC := true
		for c := 0; c < 5; c++ {
			bingoR = bingoR && board[r*5+c] == "x"
			bingoC = bingoC && board[r+c*5] == "x"

			if !bingoR && !bingoC {
				break
			}
		}
		if bingoR || bingoC {
			return true
		}
	}
	return false
}
