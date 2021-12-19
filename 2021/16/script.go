package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const part = 2

var bitCount int
var bits []int
var totalVersion int = 0

func main() {
	data, _ := os.ReadFile("input.txt")
	line := strings.Trim(strings.Split(string(data), "\n")[0], "\r\n")
	bitCount = len(line) * 4
	bits = []int{}

	for _, char := range line {
		v, _ := strconv.ParseInt(string(char), 16, 64)
		for i := 0; i < 4; i++ {
			bits = append(bits, int(v)>>(4-i-1)&1)
		}
	}
	currentBit := 0
	total := readPacket(&currentBit)

	if part == 1 {
		fmt.Println(totalVersion)
	} else {
		fmt.Println(total)
	}
}

func readPacketCount(currentBit *int, numberToRead int, typeId int) int {
	total := -1
	for i := 0; i < numberToRead; i++ {
		total = updateTotal(readPacket(currentBit), typeId, total)
	}
	return total
}
func readBitCount(currentBit *int, numberToRead int, typeId int) int {
	total := -1
	for i := *currentBit; i+numberToRead > *currentBit; {
		total = updateTotal(readPacket(currentBit), typeId, total)
	}
	return total
}

func updateTotal(newValue int, typeId int, total int) int {
	switch typeId {
	case 0:
		if total == -1 {
			total = newValue
		} else {
			total += newValue
		}
	case 1:
		if total == -1 {
			total = 1
		}
		total *= newValue
	case 2:
		if total == -1 {
			total = newValue
		} else if newValue < total {
			total = newValue
		}
	case 3:
		if newValue > total {
			total = newValue
		}
	case 5:
		if total == -1 {
			total = newValue
		} else if total > newValue {
			total = 1
		} else {
			total = 0
		}
	case 6:
		if total == -1 {
			total = newValue
		} else if total < newValue {
			total = 1
		} else {
			total = 0
		}
	case 7:
		if total == -1 {
			total = newValue
		} else if total == newValue {
			total = 1
		} else {
			total = 0
		}
	}

	return total
}

func readPacket(currentBit *int) int {
	version := nextBits(currentBit, 3)
	totalVersion += version
	typeId := nextBits(currentBit, 3)

	if typeId == 4 {
		total := 0
		for i := 0; ; i++ {
			if i != 0 {
				total = total << 4
			}
			val := nextBits(currentBit, 5)
			total |= val & 15
			if val < 16 {
				return total
			}
		}
	} else {
		lengthTypeId := nextBits(currentBit, 1)

		if lengthTypeId == 0 {
			bitCount := nextBits(currentBit, 15)
			return readBitCount(currentBit, bitCount, typeId)
		} else {
			packetCount := nextBits(currentBit, 11)
			return readPacketCount(currentBit, packetCount, typeId)
		}
	}
}

func nextBits(currentBit *int, bitCount int) int {
	value := 0
	for i := 0; i < bitCount; i++ {
		value |= bits[i+*currentBit] << (bitCount - 1 - i)
	}
	*currentBit += bitCount
	return value
}
