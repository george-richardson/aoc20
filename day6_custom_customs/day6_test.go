package day6_custom_customs

import (
	"aoc20/utils"
	"bufio"
	"strings"
	"testing"
)

func TestDay6_1(t *testing.T) {
	in, err := utils.GetInputScanner(2020, 6)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findCustomsSum(in)
	expected := 6947
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindCustomsSum(t *testing.T) {
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	scanner := bufio.NewScanner(strings.NewReader(input))
	actual := findCustomsSum(scanner)
	expected := 11
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay6_2(t *testing.T) {
	in, err := utils.GetInputScanner(2020, 6)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findCustomsSumPart2(in)
	expected := 3398
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindCustomsSumPart2(t *testing.T) {
	input := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	scanner := bufio.NewScanner(strings.NewReader(input))
	actual := findCustomsSumPart2(scanner)
	expected := 6
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}
