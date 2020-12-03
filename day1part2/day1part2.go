package day1part2

import "fmt"

// SolveProblem2 solves problem 2 of day 1
func SolveProblem2(numbers []int) (int, error) {
	for firstIndex, firstNum := range numbers {
		for secondIndex, secondNum := range numbers {
			// Can't reuse the same number multiple times (e.g., 1010 in this casegit)
			if firstIndex == secondIndex {
				continue
			}

			for thirdIndex, thirdNum := range numbers {
				if thirdIndex == secondIndex {
					continue
				}

				if firstNum+secondNum+thirdNum == 2020 {
					return firstNum * secondNum * thirdNum, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("No match found")
}
