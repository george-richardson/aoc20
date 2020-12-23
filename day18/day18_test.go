package day18

import (
	"aoc20/utils"
	"testing"
)

func TestRunCalc(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, test := range tests {
		actual := runCalc(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay18_1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 18)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := runHomework(in)
	expected := 4696493914530
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestRunCalcPart2(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, test := range tests {
		actual := runCalcPart2(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay18_2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 18)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := runHomeworkPart2(in)
	expected := 362880372308125
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
