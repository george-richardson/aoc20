package day2_password_philosophy

import (
	"bytes"
	"strconv"
	"strings"
)

func validatePasswords(passwords *[]string) (r int) {
	for _, password := range *passwords {
		if validatePassword(password) {
			r++
		}
	}
	return
}

func validatePassword(passwordRule string) bool {
	sections := strings.Split(passwordRule, " ")
	countRange := sections[0]
	char := sections[1][0]
	password := sections[2]

	bounds := strings.Split(countRange, "-")
	lower, _ := strconv.Atoi(bounds[0])
	upper, _ := strconv.Atoi(bounds[1])
	count := bytes.Count([]byte(password), []byte{char})

	return count >= lower && count <= upper
}

func validatePasswordsPart2(passwords *[]string) (r int) {
	for _, password := range *passwords {
		if validatePasswordPart2(password) {
			r++
		}
	}
	return
}

func validatePasswordPart2(passwordRule string) bool {
	sections := strings.Split(passwordRule, " ")
	countRange := sections[0]
	char := sections[1][0]
	password := sections[2]

	positions := strings.Split(countRange, "-")
	first, _ := strconv.Atoi(positions[0])
	second, _ := strconv.Atoi(positions[1])

	return (password[first - 1] == char) != (password[second - 1] == char)
}