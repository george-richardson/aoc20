package day13_shuttle_search

import (
	"aoc20/utils"
	"testing"
)

func TestDay13_1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 13)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := findNextShuttle(*in)
	expected := 4722
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindNextShuttle(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			295,
		},
	}

	for _, test := range tests {
		actual := findNextShuttle(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestFindShuttleSequence(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int64
	}{
		{
			[]string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			1068781,
		},
	}

	for _, test := range tests {
		actual := findShuttleSequenceOptimized(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay13_2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 13)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := findShuttleSequenceOptimized(*in)
	expected := int64(825305207525452)
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
