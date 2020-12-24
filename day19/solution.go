package day19

import (
	"bufio"
	"strconv"
	"strings"
)

var pointDepth = 0

type match struct {
	matches   bool
	remainder string
}

type rule interface {
	matches(string) match
	multiMatch(string) *[]match
}

type runeRule rune

func (r runeRule) multiMatch(s string) *[]match {
	m := r.matches(s)
	if m.matches {
		return &[]match{m}
	} else {
		return &[]match{}
	}
}

func (r runeRule) matches(s string) (m match) {
	if len(s) != 0 && rune(s[0]) == rune(r) {
		m.matches = true
		if len(s) > 1 {
			m.remainder = s[1:]
		}
	}
	return
}

type compoundRule struct {
	rules []rule
}

func (c compoundRule) multiMatch(s string) (matches *[]match) {
	if len(s) == 0 {
		matches = &[]match{}
		return
	}

	matches = c.rules[0].multiMatch(s)
	for _, rule := range c.rules[1:] {
		tmpMatches := make([]match, 0)
		for _, match := range *matches {
			tmpMatches = append(tmpMatches, *rule.multiMatch(match.remainder)...)
		}
		matches = &tmpMatches
	}
	return
}

func (c compoundRule) matches(s string) (m match) {
	if len(s) == 0 {
		return
	}
	m.remainder = s
	for _, rule := range c.rules {
		m = rule.matches(m.remainder)
		if !m.matches {
			return
		}
	}
	return
}

type orRule struct {
	left, right rule
}

func (o orRule) multiMatch(s string) *[]match {
	matches := *o.left.multiMatch(s)
	matches = append(matches, *o.right.multiMatch(s)...)
	return &matches
}

func (o orRule) matches(s string) (m match) {
	if len(s) == 0 {
		return
	}
	m = o.left.matches(s)
	if m.matches {
		return
	}
	m = o.right.matches(s)
	return
}

type pointerRule struct {
	pointer   int
	rulesList *map[int]rule
}

func (p pointerRule) multiMatch(s string) *[]match {
	if len(s) == 0 || pointDepth > 2000 {
		return &[]match{}
	}
	result := (*p.rulesList)[p.pointer].multiMatch(s)
	return result
}

func (p pointerRule) matches(s string) (m match) {
	if len(s) == 0 {
		return
	}
	return (*p.rulesList)[p.pointer].matches(s)
}

func parseTopLevelRule(s string, rulesList *map[int]rule) {
	tmp := strings.Split(s, ": ")
	num, _ := strconv.Atoi(tmp[0])
	rule := parseRule(tmp[1], rulesList)
	(*rulesList)[num] = rule
}

func parseRule(s string, rulesList *map[int]rule) rule {
	if strings.Contains(s, "\"") {
		return parseRuneRule(s)
	}
	if strings.Contains(s, "|") {
		return parseOrRule(s, rulesList)
	}
	return parseCompoundRule(s, rulesList)
}

func parseRuneRule(s string) rule {
	return runeRule(rune(s[1]))
}

func parseCompoundRule(s string, rulesList *map[int]rule) rule {
	elements := strings.Split(s, " ")
	rules := make([]rule, len(elements))
	for i, element := range elements {
		pointer, _ := strconv.Atoi(element)
		rules[i] = pointerRule{
			pointer:   pointer,
			rulesList: rulesList,
		}
	}
	return compoundRule{rules: rules}
}

func parseOrRule(s string, rulesList *map[int]rule) rule {
	sides := strings.Split(s, " | ")
	return orRule{
		left:  parseRule(sides[0], rulesList),
		right: parseRule(sides[1], rulesList),
	}
}

func validateInputs(scanner *bufio.Scanner) (valid int) {
	rulesList := make(map[int]rule)
	for scanner.Scan() && scanner.Text() != "" {
		parseTopLevelRule(scanner.Text(), &rulesList)
	}

	for scanner.Scan() {
		m := rulesList[0].matches(scanner.Text())
		if m.matches && len(m.remainder) == 0 {
			valid++
		}
	}
	return
}

func validateInputsPart2(scanner *bufio.Scanner) (valid int) {
	rulesList := make(map[int]rule)
	for scanner.Scan() && scanner.Text() != "" {
		parseTopLevelRule(scanner.Text(), &rulesList)
	}

	for scanner.Scan() {
		matches := *rulesList[0].multiMatch(scanner.Text())
		for _, m := range matches {
			if m.matches && len(m.remainder) == 0 {
				valid++
				break
			}
		}
	}
	return
}
