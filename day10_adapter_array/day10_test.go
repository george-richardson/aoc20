package day10_adapter_array

import (
	"aoc20/utils"
	"testing"
)

func TestFindAdapterDifferential(t *testing.T) {
	input := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	actual := findAdapterDifferential(input)
	expected := 220
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay10_1(t *testing.T) {
	in, err := utils.ReadListOfInts(2020, 10)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findAdapterDifferential(*in)
	expected := 1848
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestCountAdapterPermutations1(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}, 19208},
		{[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}, 8},
	}

	for _, test := range tests {
		actual := countAdapterPermutations(test.input)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay10_2(t *testing.T) {
	in, err := utils.ReadListOfInts(2020, 10)
	if err != nil {
		t.Error("Error getting input")
	}
	out := countAdapterPermutations(*in)
	expected := 8099130339328
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
