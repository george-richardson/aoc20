package day16

import (
	"aoc20/utils"
	"bufio"
	"strings"
	"testing"
)

func TestCalculateScanningErrorRate(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12", 71},
	}

	for _, test := range tests {
		actual := calculateScanningErrorRate(bufio.NewScanner(strings.NewReader(test.input)))
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay16_1(t *testing.T) {
	in, err := utils.GetInputScanner(2020, 16)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := calculateScanningErrorRate(in)
	expected := 26053
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestDay16_2(t *testing.T) {
	in, err := utils.GetInputScanner(2020, 16)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := calculateDepartureScore(in)
	expected := 1515506256421
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
