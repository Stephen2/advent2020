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

func countChainsFrom(joltages []int, lastJoltage int, startIndex int) int {
	count := 0

	if cached, ok := cache[lastJoltage]; ok {
		return cached
	}

	for i := startIndex; i < len(joltages); i++ {
		joltage := joltages[i]
		joltageDiff := joltage - lastJoltage

		if joltageDiff >= 1 && joltageDiff <= 3 {
			count += countChainsFrom(joltages, joltage, i)
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

	return countChainsFrom(joltages, 0, 0)
}

type JoltageChain struct {
	startAtIndex int
	prevJoltages []int
}

func SolvePart2Alt(rawJoltages []string) (possibleChains int) {
	joltages := getSortedJoltages(rawJoltages)

	stack := []JoltageChain{}
	stack = append(stack, JoltageChain{
		startAtIndex: 0,
		prevJoltages: []int{},
	})

	for len(stack) > 0 {
		joltageChain := stack[0]
		prevJoltages := []int{}

		for _, prevJoltage := range joltageChain.prevJoltages {
			print(prevJoltage, ", ")
		}
		print("| ")

		for i := joltageChain.startAtIndex; i < len(joltages); i++ {
			joltage := joltages[i]
			prevJoltages = append(prevJoltages, joltage)
			maxNextJoltageAdapter := joltage + 3

			// We may be in a position to create a new adapter chain to check
			// if i+2 or i+3 is a valid choice (AKA we could have skipped this current one)
			for j := 2; j <= 3; j++ {
				if i+j < len(joltages) && joltages[i+j] <= maxNextJoltageAdapter {
					fullChainSoFar := []int{}
					fullChainSoFar = append(fullChainSoFar, joltageChain.prevJoltages...)
					fullChainSoFar = append(fullChainSoFar, prevJoltages...)
					fullChainSoFar = append(fullChainSoFar, joltages[i+j])

					stack = append(stack, JoltageChain{
						startAtIndex: i + j + 1,
						prevJoltages: fullChainSoFar,
					})
				}
			}
		}

		possibleChains++

		// Let's write out the whole chain!
		for _, joltage := range prevJoltages {
			print(joltage, ", ")
		}

		println("---")

		stack = stack[1:]
	}

	return possibleChains
}
