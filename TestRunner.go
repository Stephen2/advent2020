package main

import (
	"advent2020/day1"
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
	rawInput := readFile("day1/input.txt")
	input := make([]int, len(rawInput))

	for i, n := range rawInput {
		num, _ := strconv.Atoi(n)
		input[i] = num
	}

	day1Solution, _ := day1.SolveProblem1(input)
	fmt.Println("day1 solution: ", day1Solution)
}
