package day7_handy_haversacks

import (
	"aoc20/utils"
	"strings"
	"testing"
)

func TestFindBagOptions(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags.\n"
	reader := strings.NewReader(input)
	actual := findBagOptions(reader)
	expected := 4
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay7_1(t *testing.T) {
	in, err := utils.GetInputReader(2020, 7)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findBagOptions(in)
	expected := 148
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestFindBagCounts(t *testing.T) {
	input := "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags.\n"
	reader := strings.NewReader(input)
	actual := findBagCount(reader)
	expected := 32
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay7_2(t *testing.T) {
	in, err := utils.GetInputReader(2020, 7)
	if err != nil {
		t.Error("Error getting input")
	}
	out := findBagCount(in)
	expected := 24867
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}
