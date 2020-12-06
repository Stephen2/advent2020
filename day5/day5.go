package day5

import (
	"sort"
	"strconv"
	"strings"
)

func binaryToInt(input string, zeroChar string, oneChar string) int64 {
	input = strings.ReplaceAll(input, zeroChar, "0")
	input = strings.ReplaceAll(input, oneChar, "1")
	num, _ := strconv.ParseInt(input, 2, 64)
	return num
}

func calcSeatId(row int64, col int64) int64 {
	return row*8 + col
}

func Solve(rawSeatInfo []string) (highestSeatId int64) {
	for _, line := range rawSeatInfo {
		row := binaryToInt(line[0:7], "F", "B")
		col := binaryToInt(line[7:10], "L", "R")
		seatId := calcSeatId(row, col)

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	return highestSeatId
}

func SolvePart2(rawSeatInfo []string) (seatId int64) {
	seatIds := []int64{}

	for _, line := range rawSeatInfo {
		row := binaryToInt(line[0:7], "F", "B")
		col := binaryToInt(line[7:10], "L", "R")
		seatId := calcSeatId(row, col)
		seatIds = append(seatIds, seatId)
	}

	sort.Slice(seatIds, func(i int, j int) bool { return seatIds[i] < seatIds[j] })

	for i := 0; i < len(seatIds)-2; i++ {
		if seatIds[i] != seatIds[i+1]-1 {
			return seatIds[i] + 1
		}
	}

	panic("no gap found")
}
