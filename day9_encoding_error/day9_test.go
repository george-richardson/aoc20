package day9_encoding_error

import (
	"aoc20/utils"
	"testing"
)

func TestFindInvalidXmas(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	actual := findInvalidXMAS(input, 5)
	expected := 127
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay9_1(t *testing.T) {
	in, err := utils.ReadListOfInts(2020, 9)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findInvalidXMAS(*in, 25)
	expected := 756008079
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindEncryptionWeakness(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	actual := findEncryptionWeakness(input, 127)
	expected := 62
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay9_2(t *testing.T) {
	in, err := utils.ReadListOfInts(2020, 9)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findEncryptionWeakness(*in, 756008079)
	expected := 93727241
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
