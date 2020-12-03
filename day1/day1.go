package day1

import "fmt"

// SolveProblem1 solves problem 1 of day 1
func SolveProblem1(numbers []int) (int, error) {
	for firstIndex, firstNum := range numbers {
		for secondIndex, secondNum := range numbers {
			// Can't reuse the same number multiple times (e.g., 1010 in this case)
			if firstIndex == secondIndex {
				continue
			}

			if firstNum+secondNum == 2020 {
				return firstNum * secondNum, nil
			}
		}
	}

	return 0, fmt.Errorf("No match found")
}
