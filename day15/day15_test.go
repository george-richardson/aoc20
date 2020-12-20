package day15

import (
	"aoc20/utils"
	"testing"
)

func TestElfGame(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"0,3,6", 436},
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for _, test := range tests {
		actual := elfGame(test.input, 2020)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay15_1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 15)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := elfGame((*in)[0], 2020)
	expected := 517
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestElfGameLarger(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"0,3,6", 175594},
		{"1,3,2", 2578},
		{"2,1,3", 3544142},
		{"1,2,3", 261214},
		{"2,3,1", 6895259},
		{"3,2,1", 18},
		{"3,1,2", 362},
	}

	for _, test := range tests {
		actual := elfGame(test.input, 30000000)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay15_2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 15)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := elfGame((*in)[0], 30000000)
	expected := 1047739
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
