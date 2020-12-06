package day4

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func IsValid(passport *Passport) bool {
	return passport.Byr != "" &&
		passport.Iyr != "" &&
		passport.Eyr != "" &&
		passport.Hgt != "" &&
		passport.Hcl != "" &&
		passport.Ecl != "" &&
		passport.Pid != ""

	// HACK: allow blank country codes because I am Hacker
	// passport.Cid != ""
}

func isNumBetween(value string, min int, max int) bool {
	val, err := strconv.Atoi(value)

	return err == nil && val >= min && val <= max
}

func isHeightValid(height string) bool {
	re := regexp.MustCompile(`^(\d{2,3})(cm|in)$`)
	matches := re.FindStringSubmatch(height)

	if len(matches) != 3 {
		return false
	}

	num, unit := matches[1], matches[2]

	switch unit {
	case "cm":
		return isNumBetween(num, 150, 193)
	case "in":
		return isNumBetween(num, 59, 76)
	default:
		panic("wtf")
	}
}

func isHclValid(hcl string) bool {
	re := regexp.MustCompile(`^#[\da-f]{6}$`)
	return re.MatchString(hcl)
}

func isEclValid(ecl string) bool {
	switch ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func IsValidComplex(passport *Passport) bool {
	isValid := isNumBetween(passport.Byr, 1920, 2002)
	isValid = isValid && isNumBetween(passport.Iyr, 2010, 2020)
	isValid = isValid && isNumBetween(passport.Eyr, 2020, 2030)
	isValid = isValid && isHeightValid(passport.Hgt)
	isValid = isValid && isHclValid(passport.Hcl)
	isValid = isValid && isEclValid(passport.Ecl)
	isValid = isValid && regexp.MustCompile(`^\d{9}$`).MatchString(passport.Pid)

	return isValid
}

func Solve(rawPassportData []string, validate func(*Passport) bool) int {
	passportData := []*Passport{}
	passport := &Passport{}

	for _, line := range rawPassportData {
		if line == "" {
			passportData = append(passportData, passport)
			passport = &Passport{}
		} else {
			fields := strings.Fields(line)

			for _, field := range fields {
				kv := strings.Split(field, ":")
				key, value := kv[0], kv[1]

				switch key {
				case "byr":
					passport.Byr = value
				case "iyr":
					passport.Iyr = value
				case "eyr":
					passport.Eyr = value
				case "hgt":
					passport.Hgt = value
				case "hcl":
					passport.Hcl = value
				case "ecl":
					passport.Ecl = value
				case "pid":
					passport.Pid = value
				case "cid":
					passport.Cid = value
				}
			}
		}
	}

	passportData = append(passportData, passport)

	// Check validity
	validCount := 0

	for _, passport := range passportData {
		if validate(passport) {
			validCount++
		}
	}

	return validCount
}
