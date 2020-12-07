package day7

import (
	"regexp"
	"strconv"
	"strings"
)

type BagRule struct {
	colour     string
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
		colour:     colour,
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

func getRules(rawRules []string) map[string]BagRule {
	allRules := make(map[string]BagRule)

	for _, rawRule := range rawRules {
		colour, rule := rawRuleToRule(rawRule)
		allRules[colour] = rule
	}

	return allRules
}

func Solve(rawRules []string) int {
	allRules := getRules(rawRules)
	bagsThatCanContainShinyGold := make([]BagRule, 0)

	for _, rule := range allRules {
		if bagCanContain("shiny gold", rule, allRules) {
			bagsThatCanContainShinyGold = append(bagsThatCanContainShinyGold, rule)
		}
	}

	return len(bagsThatCanContainShinyGold)
}

type ContainedBag struct {
	colour  string
	numBags int
}

func SolvePart2(rawRules []string) int {
	allRules := getRules(rawRules)
	toCheck := []ContainedBag{}
	toCheck = append(toCheck, ContainedBag{colour: "shiny gold", numBags: 1})
	requiredBags := 0

	for len(toCheck) > 0 {
		rule := allRules[toCheck[0].colour]

		for containedColour, containedNumBags := range rule.canContain {
			bagsToAdd := containedNumBags * toCheck[0].numBags
			requiredBags += bagsToAdd
			toCheck = append(toCheck, ContainedBag{colour: containedColour, numBags: bagsToAdd})
		}

		toCheck = toCheck[1:]
	}

	return requiredBags
}
