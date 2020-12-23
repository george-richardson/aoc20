package day16

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type ticket []int

type ticketInput struct {
	rules         []*validityRule
	yourTicket    ticket
	nearbyTickets []ticket
}

type validityRange struct {
	lower, upper int
}

type validityRule struct {
	name           string
	validityRanges []validityRange
}

func (rule validityRule) validate(value int) (valid bool) {
	for _, validityRange := range rule.validityRanges {
		if value <= validityRange.upper && value >= validityRange.lower {
			return true
		}
	}
	return false
}

func (ticket ticket) validate(rules []*validityRule) (valid bool) {
	for _, value := range ticket {
		valid := false
		for _, rule := range rules {
			if rule.validate(value) {
				valid = true
				break
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

func calculateScanningErrorRate(scanner *bufio.Scanner) (errorRate int) {
	input := parseInput(scanner)
	invalidValues := make([]int, 0)
	for _, ticket := range input.nearbyTickets {
	VALUELOOP:
		for _, value := range ticket {
			for _, rule := range input.rules {
				if rule.validate(value) {
					continue VALUELOOP
				}
			}
			invalidValues = append(invalidValues, value)
		}
	}
	for _, value := range invalidValues {
		errorRate += value
	}
	return
}

func parseInput(input *bufio.Scanner) ticketInput {
	rules := *parseRules(input)
	input.Scan()
	input.Scan()
	yourTicket := parseTicket(input.Text())
	input.Scan()
	input.Scan()
	nearbyTickets := make([]ticket, 0)
	for input.Scan() {
		nearbyTickets = append(nearbyTickets, parseTicket(input.Text()))
	}
	return ticketInput{
		rules:         rules,
		yourTicket:    yourTicket,
		nearbyTickets: nearbyTickets,
	}
}

func parseRules(input *bufio.Scanner) *[]*validityRule {
	rules := make([]*validityRule, 0)
	for input.Scan() && input.Text() != "" {
		segments := strings.Split(input.Text(), ": ")
		ruleRanges := make([]validityRange, 0)
		for _, rangeString := range strings.Split(segments[1], " or ") {
			values := strings.Split(rangeString, "-")
			ruleRange := validityRange{}
			ruleRange.lower, _ = strconv.Atoi(values[0])
			ruleRange.upper, _ = strconv.Atoi(values[1])
			ruleRanges = append(ruleRanges, ruleRange)
		}
		rules = append(rules, &validityRule{
			name:           segments[0],
			validityRanges: ruleRanges,
		})
	}
	return &rules
}

func parseTicket(line string) (ticket []int) {
	valStrings := strings.Split(line, ",")
	for _, valString := range valStrings {
		val, _ := strconv.Atoi(valString)
		ticket = append(ticket, val)
	}
	return
}

func calculateDepartureScore(scanner *bufio.Scanner) (score int) {
	input := parseInput(scanner)
	validTickets := make([]ticket, 0)
	for _, ticket := range input.nearbyTickets {
		if ticket.validate(input.rules) {
			validTickets = append(validTickets, ticket)
		}
	}

	type ruleOption struct {
		rule    *validityRule
		columns *[]int
	}

	var options = make([]ruleOption, len(input.rules))
	for i, rule := range input.rules {
		columns := make([]int, 0)
		option := ruleOption{rule, &columns}
	COLLOOP:
		for i := 0; i < len(validTickets[0]); i++ {
			for _, ticket := range validTickets {
				if !rule.validate(ticket[i]) {
					continue COLLOOP
				}
			}
			columns = append(columns, i)
		}
		options[i] = option
	}

	sort.Slice(options, func(i, j int) bool {
		return len(*options[i].columns) < len(*options[j].columns)
	})

	ruleMappings := make(map[int]*validityRule)

OPTION:
	for _, option := range options {
		for _, col := range *option.columns {
			if ruleMappings[col] == nil {
				ruleMappings[col] = option.rule
				continue OPTION
			}
		}
	}

	score = 1
	for col, rule := range ruleMappings {
		if len((*rule).name) > 8 && (*rule).name[0:9] == "departure" {
			score *= input.yourTicket[col]
		}
	}

	return
}
