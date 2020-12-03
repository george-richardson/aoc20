package day3_toboggan_trajectory

import (
	"aoc20/utils"
	"fmt"
	"testing"
)

func TestDay3_1(t *testing.T) {
	in, err := utils.ReadGridOfBytes(2020, 3)
	if err != nil {
		t.Error("Error getting input")
	}
	out := checkSlopePath(in)
	expected := 156
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestCheckSlope(t *testing.T) {
	var tests = []struct {
		input    [][]byte
		expected int
	}{
		{[][]byte{
			{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'},
			{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.'},
			{'.', '#', '.', '.', '.', '.', '#', '.', '.', '#', '.'},
			{'.', '.', '#', '.', '#', '.', '.', '.', '#', '.', '#'},
			{'.', '#', '.', '.', '.', '#', '#', '.', '.', '#', '.'},
			{'.', '.', '#', '.', '#', '#', '.', '.', '.', '.', '.'},
			{'.', '#', '.', '#', '.', '#', '.', '.', '.', '.', '#'},
			{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
			{'#', '.', '#', '#', '.', '.', '.', '#', '.', '.', '.'},
			{'#', '.', '.', '.', '#', '#', '.', '.', '.', '.', '#'},
			{'.', '#', '.', '.', '#', '.', '.', '.', '#', '.', '#'},
		}, 7},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %v expect %v", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := checkSlopePath(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestDay3_2(t *testing.T) {
	in, err := utils.ReadGridOfBytes(2020, 3)
	if err != nil {
		t.Error("Error getting input")
	}
	out := checkSlopePaths(in)
	expected := 3521829480
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestCheckSlopes(t *testing.T) {
	var tests = []struct {
		input    [][]byte
		expected int
	}{
		{[][]byte{
			{'.', '.', '#', '#', '.', '.', '.', '.', '.', '.', '.'},
			{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.'},
			{'.', '#', '.', '.', '.', '.', '#', '.', '.', '#', '.'},
			{'.', '.', '#', '.', '#', '.', '.', '.', '#', '.', '#'},
			{'.', '#', '.', '.', '.', '#', '#', '.', '.', '#', '.'},
			{'.', '.', '#', '.', '#', '#', '.', '.', '.', '.', '.'},
			{'.', '#', '.', '#', '.', '#', '.', '.', '.', '.', '#'},
			{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
			{'#', '.', '#', '#', '.', '.', '.', '#', '.', '.', '.'},
			{'#', '.', '.', '.', '#', '#', '.', '.', '.', '.', '#'},
			{'.', '#', '.', '.', '#', '.', '.', '.', '#', '.', '#'},
		}, 336},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %v expect %v", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := checkSlopePaths(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})
	}
}
