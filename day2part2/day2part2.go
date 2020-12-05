package day2part2

import (
	"regexp"
	"strconv"
)

func isValid(password string) bool {
	// e.g., 4-7 z: zzzfzlzzz
	re := regexp.MustCompile("^(\\d*)\\-(\\d*) (.): (.*)$")
	matches := re.FindStringSubmatch(password)

	firstPos, _ := strconv.Atoi(matches[1])
	secondPos, _ := strconv.Atoi(matches[2])
	char := matches[3]
	pass := matches[4]

	numMatches := 0

	if string(pass[firstPos-1]) == char {
		numMatches++
	}

	if string(pass[secondPos-1]) == char {
		numMatches++
	}

	return numMatches == 1
}

// SolvePart2 solves day2part2
func SolvePart2(passwords []string) int {
	validCount := 0

	for _, password := range passwords {
		if isValid(password) {
			validCount++
		}
	}

	return validCount
}
