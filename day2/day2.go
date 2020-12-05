package day2

import (
	"regexp"
	"strconv"
)

func isValid(password string) bool {
	// e.g., 4-7 z: zzzfzlzzz
	re := regexp.MustCompile(`^(\d*)\-(\d*) (.): (.*)$`)
	matches := re.FindStringSubmatch(password)

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])
	char := matches[3]
	pass := matches[4]

	occurs := 0

	for _, c := range pass {
		if string(c) == char {
			occurs++
		}
	}

	return occurs >= min && occurs <= max
}

// Solve solves day2
func Solve(passwords []string) int {
	validCount := 0

	for _, password := range passwords {
		if isValid(password) {
			validCount++
		}
	}

	return validCount
}
