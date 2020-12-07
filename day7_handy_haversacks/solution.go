package day7_handy_haversacks

import (
	"errors"
	"io"
)

func findBagOptions(reader io.Reader) (r int) {
	parser := NewParser(reader)
	rules := make(map[string]*visitableRule)
	for err := errors.New(""); err != nil; _, err = parser.scanForSequence(&[]Token{WS, EOF}) {
		rule, err := parser.ParseLine()
		if err != nil {
			panic(err)
		}
		rules[rule.name] = &visitableRule{bagRule: *rule}
	}

	for _, rule := range rules {
		if visitRule(rule, &rules) {
			r++
		}
	}

	return
}

type visitableRule struct {
	bagRule
	visited   bool
	contained bool
}

func visitRule(rule *visitableRule, rules *map[string]*visitableRule) bool {
	if rule.visited {
		return rule.contained
	}
	rule.visited = true

	for _, contain := range rule.contains {
		if contain.name == "shiny gold" {
			rule.contained = true
		} else if visitRule((*rules)[contain.name], rules) {
			rule.contained = true
		}
	}
	return rule.contained
}

type visitableRuleCount struct {
	bagRule
	visited bool
	count   int
}

func findBagCount(reader io.Reader) (r int) {
	parser := NewParser(reader)
	rules := make(map[string]*visitableRuleCount)
	for err := errors.New(""); err != nil; _, err = parser.scanForSequence(&[]Token{WS, EOF}) {
		rule, err := parser.ParseLine()
		if err != nil {
			panic(err)
		}
		rules[rule.name] = &visitableRuleCount{bagRule: *rule}
	}

	return visitRuleCount(rules["shiny gold"], &rules) - 1
}

func visitRuleCount(rule *visitableRuleCount, rules *map[string]*visitableRuleCount) int {
	if rule.visited {
		return rule.count
	}
	rule.visited = true
	rule.count = 1

	for _, contain := range rule.contains {
		rule.count += visitRuleCount((*rules)[contain.name], rules) * contain.count
	}

	return rule.count
}
