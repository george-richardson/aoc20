package day2_password_philosophy

import (
	"aoc20/utils"
	"fmt"
	"testing"
)

func TestDay2_1(t *testing.T)  {
	in, err := utils.ReadListOfStrings(2020, 2)
	if err != nil {
		t.Error("Error getting input")
	}
	out := validatePasswords(in)
	expected := 445
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestValidatePassword(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{ "1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc" }, 2},
		{[]string{ "1-2 q: qqqde", "1-3 z: azzza", "2-9 c: cccccccccccccccc" }, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %v expect %v", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := validatePasswords(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestDay2_2(t *testing.T)  {
	in, err := utils.ReadListOfStrings(2020, 2)
	if err != nil {
		t.Error("Error getting input")
	}
	out := validatePasswordsPart2(in)
	expected := 491
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestValidatePasswordPart2(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{ "1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc" }, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %v expect %v", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := validatePasswordsPart2(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})
	}
}