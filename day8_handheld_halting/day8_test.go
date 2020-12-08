package day8_handheld_halting

import (
	"aoc20/utils"
	"strings"
	"testing"
)

func TestRunProgramFindCycle(t *testing.T) {
	input := strings.Split("nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6", "\n")
	actual := runProgramFindCycle(&input)
	expected := 5
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay8_1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 8)
	if err != nil {
		t.Error("Error getting input")
	}
	out := runProgramFindCycle(in)
	expected := 1179
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestPatchProgram(t *testing.T) {
	input := strings.Split("nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6", "\n")
	actual := patchProgram(&input)
	expected := 8
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay8_2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2020, 8)
	if err != nil {
		t.Error("Error getting input")
	}
	out := patchProgram(in)
	expected := 1089
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
