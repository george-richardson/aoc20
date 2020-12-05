package day5_binary_boarding

import (
	"aoc20/utils"
	"testing"
)

func TestDay5_1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 5)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findMaxSeatLocation(in)
	expected := 963
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindSeatLocation(t *testing.T) {
	var tests = []struct {
		input           string
		row, column, id int
	}{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for _, test := range tests {
		row, column, id := findSeatLocation(test.input)
		if row != test.row {
			t.Errorf("ROW %v: actual %v, expected %v", test.input, row, test.row)
		}
		if column != test.column {
			t.Errorf("COL %v: actual %v, expected %v", test.input, column, test.column)
		}
		if id != test.id {
			t.Errorf("SID %v: actual %v, expected %v", test.input, id, test.id)
		}
	}
}

func TestDay5_2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 5)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findMissingSeat(in)
	expected := 592
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
