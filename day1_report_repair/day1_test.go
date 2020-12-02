package main

import (
	"aoc20/utils"
	"fmt"
	"testing"
)

func TestDay1_1(t *testing.T)  {
	in, err := utils.ReadListOfInts(2020, 1)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findSum2020ThenMultiply(in)
	expected := 1014171
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindSum2020ThenMultiply(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2019, 1, 0}, 2019},
		{[]int{1721, 979, 366, 299, 675, 1456}, 514579},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %d expect %d", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := findSum2020ThenMultiply(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %d, expected %d", actual, tt.expected)
			}
		})
	}
}

func TestDay1_2(t *testing.T)  {
	in, err := utils.ReadListOfInts(2020, 1)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findSum2020ThenMultiplyPart2(in)
	expected := 46584630
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindSum2020ThenMultiplyPart2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{2018, 1, 1}, 2018},
		{[]int{1721, 979, 366, 299, 675, 1456}, 241861950},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %d expect %d", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := findSum2020ThenMultiplyPart2(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %d, expected %d", actual, tt.expected)
			}
		})
	}
}
