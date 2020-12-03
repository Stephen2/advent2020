package main

import (
	"advent2020/day1"
	"advent2020/day1part2"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	return strings.Split(string(content), "\n")
}

func main() {
	// Day1
	rawInput := readFile("day1/input.txt")
	input := make([]int, len(rawInput))

	for i, n := range rawInput {
		num, _ := strconv.Atoi(n)
		input[i] = num
	}

	day1Solution, _ := day1.SolveProblem1(input)
	fmt.Println("day1 solution: ", day1Solution)

	// Day1Part2
	day1part2Solution, _ := day1part2.SolveProblem2(input)
	fmt.Println("day1part2 solution: ", day1part2Solution)
}
