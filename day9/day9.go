package day9

import (
	"math"
	"strconv"
)

func getNumbers(rawNumbers []string) []int {
	nums := make([]int, len(rawNumbers))

	for i, n := range rawNumbers {
		nums[i], _ = strconv.Atoi(n)
	}

	return nums
}

func isNumberValid(number int, toConsider []int) bool {
	for i, n1 := range toConsider {
		for j, n2 := range toConsider {
			if i == j {
				continue
			}

			if n1+n2 == number {
				return true
			}
		}
	}

	return false
}

func Solve(rawNumbers []string, preambleLength int, lookBehind int) int {
	numbers := getNumbers(rawNumbers)

	for i := preambleLength; i < len(numbers); i++ {
		toConsider := numbers[i-lookBehind : i]

		if !isNumberValid(numbers[i], toConsider) {
			return numbers[i]
		}
	}

	panic("all look legit!")
}

func SolvePart2(rawNumbers []string, numToFind int) int {
	numbers := getNumbers(rawNumbers)

	for i := 0; i < len(numbers); i++ {
		for howManyToAdd := 1; true; howManyToAdd++ {
			sum := 0
			smallest := math.MaxInt32
			largest := 0

			for n := i; n < i+howManyToAdd; n++ {
				num := numbers[n]

				if num < smallest {
					smallest = num
				}

				if num > largest {
					largest = num
				}

				sum += num
			}

			if sum == numToFind {
				return smallest + largest
			}

			if sum > numToFind {
				break
			}
		}
	}

	panic("no valid answer found")
}
