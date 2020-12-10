package day10

import (
	"sort"
	"strconv"
)

func getSortedJoltages(rawJoltages []string) []int {
	joltages := make([]int, len(rawJoltages))

	for i, rawJoltage := range rawJoltages {
		joltages[i], _ = strconv.Atoi(rawJoltage)
	}

	sort.Ints(joltages)

	return joltages
}

func Solve(rawJoltages []string) int {
	joltDiff1 := 0
	joltDiff3 := 1 // Start at 1 for the device

	joltages := getSortedJoltages(rawJoltages)
	lastJoltage := 0

	for _, joltage := range joltages {
		if joltage-lastJoltage == 1 {
			joltDiff1++
		}

		if joltage-lastJoltage == 3 {
			joltDiff3++
		}

		lastJoltage = joltage
	}

	return joltDiff1 * joltDiff3
}

func countChainsFrom(joltages []int, lastJoltage int) int {
	count := 0

	if cached, ok := cache[lastJoltage]; ok {
		return cached
	}

	for _, joltage := range joltages {
		joltageDiff := joltage - lastJoltage

		if joltageDiff >= 1 && joltageDiff <= 3 {
			count += countChainsFrom(joltages, joltage)
		}
	}

	if lastJoltage == joltages[len(joltages)-1] {
		count++
	}

	cache[lastJoltage] = count

	return count
}

var cache = make(map[int]int)

func SolvePart2(rawJoltages []string) int {
	joltages := getSortedJoltages(rawJoltages)

	return countChainsFrom(joltages, 0)
}
