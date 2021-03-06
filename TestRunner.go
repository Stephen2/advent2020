package main

import (
	"advent2020/day1"
	"advent2020/day10"
	"advent2020/day11"
	"advent2020/day12"
	"advent2020/day13"
	"advent2020/day1part2"
	"advent2020/day2"
	"advent2020/day2part2"
	"advent2020/day3"
	"advent2020/day3part2"
	"advent2020/day4"
	"advent2020/day5"
	"advent2020/day6"
	"advent2020/day7"
	"advent2020/day8"
	"advent2020/day9"
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
	fmt.Println("day1 solution:", day1Solution)

	// Day1Part2
	day1part2Solution, _ := day1part2.SolveProblem2(input)
	fmt.Println("day1part2 solution:", day1part2Solution)

	// Day2
	passwords := readFile("day2/input.txt")
	fmt.Println("day2:", day2.Solve(passwords))
	fmt.Println("day2part2:", day2part2.SolvePart2(passwords))

	// Day3
	mapLines := readFile("day3/input.txt")
	println("day3:", day3.Solve(mapLines, 3))
	a := day3part2.Solve(mapLines, 1, 1)
	b := day3part2.Solve(mapLines, 3, 1)
	c := day3part2.Solve(mapLines, 5, 1)
	d := day3part2.Solve(mapLines, 7, 1)
	e := day3part2.Solve(mapLines, 1, 2)
	println(a, b, c, d, e)
	println("day3part2:", a*b*c*d*e)

	// Day4
	rawPassportData := readFile("day4/input.txt")
	println("day4:", day4.Solve(rawPassportData, day4.IsValid))
	println("day4part2:", day4.Solve(rawPassportData, day4.IsValidComplex))

	// Day5
	rawSeatInfo := readFile("day5/input.txt")
	println("day5:", day5.Solve(rawSeatInfo))
	println("day5part2:", day5.SolvePart2(rawSeatInfo))

	// Day6
	rawAnswers := readFile("day6/input.txt")
	println("day6:", day6.Solve(rawAnswers))
	println("day6part2:", day6.SolvePart2(rawAnswers))

	// Day7
	rawRules := readFile("day7/input.txt")
	println("day7:", day7.Solve(rawRules))
	println("day7part2:", day7.SolvePart2(rawRules))

	// Day8
	rawInstructions := readFile("day8/input.txt")
	println("day8:", day8.Solve(rawInstructions))
	println("day8part2:", day8.SolvePart2(rawInstructions))

	// Day9
	rawNumbers := readFile("day9/input.txt")
	day9Solution := day9.Solve(rawNumbers, 25, 25)
	println("day9:", day9Solution)
	println("day9part2:", day9.SolvePart2(rawNumbers, day9Solution))

	// Day10
	rawJoltages := readFile("day10/input.txt")
	println("day10:", day10.Solve(rawJoltages))
	println("day10part2:", day10.SolvePart2(rawJoltages))
	// println("day10part2:", day10.SolvePart2Alt(rawJoltages))

	// Day11
	rawSeats := readFile("day11/input.txt")
	println("day11:", day11.Solve(rawSeats))
	println("day11part2:", day11.SolvePart2(rawSeats))

	// Day12
	rawDirections := readFile("day12/input.txt")
	println("day12:", day12.Solve(rawDirections))
	println("day12part2:", day12.SolvePart2(rawDirections))

	// Day13
	rawBusInput := readFile("day13/input.txt")
	println("day13:", day13.Solve(rawBusInput))
	println("day13part2:", day13.SolvePart2(rawBusInput))
}
