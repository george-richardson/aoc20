package day17

import (
	"aoc20/utils"
	"testing"
)

func TestDay17_1(t *testing.T) {
	in, err := utils.ReadGridOfBytes(2020, 17)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := bootCubeSpace(*in, 6)
	//expected := 304
	//if out != expected {
	//	t.Errorf("%d does not match expected %d", out, expected)
	//}
	t.Log(out)
}

func TestCalculateScanningErrorRate(t *testing.T) {
	var tests = []struct {
		input    [][]byte
		expected int
	}{
		{
			[][]byte{
				{'.', '#', '.'},
				{'.', '.', '#'},
				{'#', '#', '#'},
			}, 112},
	}

	for _, test := range tests {
		actual := bootCubeSpace(test.input, 6)
		if actual != test.expected {
			t.Errorf("actual %v, expected %v", actual, test.expected)
		}
	}
}

func TestDay17_2(t *testing.T) {
	in, err := utils.ReadGridOfBytes(2020, 17)
	if err != nil {
		t.Fatal("Error getting input")
	}
	out := bootHyperCubeSpace(*in, 6)
	expected := 1868
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
