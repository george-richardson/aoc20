package day12

import (
	"aoc20/utils"
	"testing"
)

func TestFindDestinationManhattanDistance(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			25,
		},
	}

	for _, test := range tests {
		actual := findDestinationManhattanDistance(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay12_1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 12)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := findDestinationManhattanDistance(*in)
	expected := 1294
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindDestinationManhattanDistanceWaypoint(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			286,
		},
	}

	for _, test := range tests {
		actual := findDestinationManhattanDistanceWaypoint(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay12_2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 12)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := findDestinationManhattanDistanceWaypoint(*in)
	expected := 20592
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
