package day4_passport_processing

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func validatePassports(input *bufio.Scanner) (r int) {
	var passports []map[string]string
	passport := map[string]string{}
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}
		for _, segment := range strings.Split(line, " ") {
			sides := strings.Split(segment, ":")
			passport[sides[0]] = sides[1]
		}
	}
	passports = append(passports, passport)

OUTER:
	for _, passport := range passports {
		for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if _, present := passport[key]; !present {
				continue OUTER
			}
		}
		r++
	}

	return
}

func validatePassportsPart2(input *bufio.Scanner) (r int) {
	var passports []map[string]string
	passport := map[string]string{}
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}
		for _, segment := range strings.Split(line, " ") {
			sides := strings.Split(segment, ":")
			passport[sides[0]] = sides[1]
		}
	}
	passports = append(passports, passport)

	validators := map[string]func(string) bool{
		"byr": validateYearStringFunc(1920, 2002),
		"iyr": validateYearStringFunc(2010, 2020),
		"eyr": validateYearStringFunc(2020, 2030),
		"hgt": validateHeight,
		"hcl": validateHairColor,
		"ecl": validateEyeColor,
		"pid": validatePassportID,
	}
OUTER:
	for _, passport := range passports {
		for key, validator := range validators {
			if value, present := passport[key]; !present || !validator(value) {
				continue OUTER
			}
		}
		r++
	}

	return
}

func validateYearStringFunc(lower, upper int) func(s string) bool {
	return func(s string) bool {
		return validateYearString(s, lower, upper)
	}
}

func validateYearString(s string, lower, upper int) bool {
	number, err := strconv.Atoi(s)
	return err == nil && boundsCheck(number, lower, upper)
}

func boundsCheck(input, lower, upper int) bool {
	return input >= lower && input <= upper
}

func validateHeight(s string) bool {
	l := len(s)
	if l < 3 {
		return false
	}
	value, err := strconv.Atoi(s[:l-2])
	unit := s[l-2:]
	return err == nil && ((unit == "cm" && boundsCheck(value, 150, 193)) || (unit == "in" && boundsCheck(value, 59, 76)))
}

func validateHairColor(s string) bool {
	matched, err := regexp.MatchString("^#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]$", s)
	return err == nil && matched
}

func validateEyeColor(s string) bool {
	switch s {
	case
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func validatePassportID(s string) bool {
	matched, err := regexp.MatchString("^[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]$", s)
	return err == nil && matched
}
