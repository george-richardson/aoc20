package day11_seating_system

import (
	"aoc20/utils"
	"testing"
)

func TestFindOccupiedConways(t *testing.T) {
	var tests = []struct {
		input    [][]byte
		expected int
	}{
		{
			[][]byte{
				{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
				{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
				{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			},
			37,
		},
	}

	for _, test := range tests {
		actual := findOccupiedConways(&test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay10_1(t *testing.T) {
	in, err := utils.ReadGridOfBytes(2020, 11)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findOccupiedConways(in)
	expected := 2319
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindOccupiedConwaysAxis(t *testing.T) {
	var tests = []struct {
		input    [][]byte
		expected int
	}{
		{
			[][]byte{
				{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
				{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
				{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
				{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
				{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
			},
			26,
		},
	}

	for _, test := range tests {
		actual := findOccupiedConwaysAxis(&test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay10_2(t *testing.T) {
	in, err := utils.ReadGridOfBytes(2020, 11)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findOccupiedConwaysAxis(in)
	expected := 2117
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
