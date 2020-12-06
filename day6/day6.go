package day6

type AnswerGroup struct {
	numPeople int
	answers   string
}

func SolvePart2(rawAnswers []string) int {
	groups := getAnswerGroups(rawAnswers)
	totalYes := 0

	for _, group := range groups {
		uniq := map[rune]int{}

		for _, c := range group.answers {
			val := uniq[c]
			uniq[c] = val + 1
		}

		for _, num := range uniq {
			// If everyone answered YES
			if num == group.numPeople {
				totalYes++
			}
		}
	}

	return totalYes
}

func getAnswerGroups(rawAnswers []string) []*AnswerGroup {
	groups := []*AnswerGroup{}
	group := &AnswerGroup{}

	for _, line := range rawAnswers {
		if line == "" {
			groups = append(groups, group)
			group = &AnswerGroup{}
		} else {
			group.answers += line
			group.numPeople++
		}
	}

	groups = append(groups, group)

	return groups
}

func Solve(rawAnswers []string) int {
	groups := getAnswerGroups(rawAnswers)
	totalYes := 0

	for _, group := range groups {
		uniq := map[rune]struct{}{}

		for _, c := range group.answers {
			_, ok := uniq[c]

			if !ok {
				uniq[c] = struct{}{}
			}
		}

		totalYes += len(uniq)
	}

	return totalYes
}
