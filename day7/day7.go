package day7

import (
	"regexp"
	"strconv"
	"strings"
)

type BagRule struct {
	canContain map[string]int
}

var ruleRegex = regexp.MustCompile(`^(.*) bags? contain (.*)$`)
var canContainRegex = regexp.MustCompile(`(\d) (.*) (?:bags?\.?)`)

func rawRuleToRule(rawRule string) (string, BagRule) {
	matches := ruleRegex.FindStringSubmatch(rawRule)
	colour := matches[1]
	canContain := make(map[string]int)

	if matches[2] != "no other bags." {
		canContainDetails := strings.Split(matches[2], ", ")

		for _, canContainDetail := range canContainDetails {
			canContainMatches := canContainRegex.FindStringSubmatch(canContainDetail)
			canContainNum, _ := strconv.Atoi(canContainMatches[1])
			canContainColour := canContainMatches[2]

			// Finally, add to the map for the Rule
			canContain[canContainColour] = canContainNum
		}
	}

	return colour, BagRule{
		canContain: canContain,
	}
}

func bagCanContain(colour string, rule BagRule, allRules map[string]BagRule) bool {
	// I had this as a recursive function and got unpredictable results O.o
	// I'd love to talk it through with someone

	var toCheck []BagRule
	toCheck = append(toCheck, rule)

	for len(toCheck) > 0 {
		if _, ok := toCheck[0].canContain[colour]; ok {
			return true
		}

		for colour := range toCheck[0].canContain {
			rule := allRules[colour]
			toCheck = append(toCheck, rule)
		}

		toCheck = toCheck[1:]
	}

	return false
}

func Solve(rawRules []string) int {
	allRules := make(map[string]BagRule)

	for _, rawRule := range rawRules {
		colour, rule := rawRuleToRule(rawRule)
		allRules[colour] = rule
	}

	bagsThatCanContainShinyGold := make([]BagRule, 0)

	for _, rule := range allRules {
		if bagCanContain("shiny gold", rule, allRules) {
			bagsThatCanContainShinyGold = append(bagsThatCanContainShinyGold, rule)
		}
	}

	return len(bagsThatCanContainShinyGold)
}
