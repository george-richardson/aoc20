package day13_shuttle_search

import (
	"math"
	"strconv"
	"strings"
)

func findNextShuttle(input []string) int {
	remainder, smallestRemainder, smallestValue := 0, math.MaxInt32, math.MaxInt32
	bestDeparture, _ := strconv.Atoi(input[0])
	for _, s := range strings.Split(input[1], ",") {
		value, err := strconv.Atoi(s)
		if err == nil {
			nextDeparture := (bestDeparture/value)*value + value
			remainder = nextDeparture % bestDeparture
			if remainder < smallestRemainder {
				smallestRemainder, smallestValue = remainder, value
			}
		}
	}
	return smallestRemainder * smallestValue
}

func findShuttleSequence(input []string) int64 {
	valueStrings := strings.Split(input[1], ",")
	values := []int64{}
	var largestValue int64
	largestIndex := 0
	for i, s := range valueStrings {
		v, err := strconv.Atoi(s)
		if err != nil {
			values = append(values, -1)
		} else {
			if int64(v) > largestValue {
				largestValue = int64(v)
				largestIndex = i
			}
			values = append(values, int64(v))
		}
	}
	var start int64 = largestValue
OUT:
	for true {
		start += largestValue
		for i, v := range values {
			if v == -1 || i == largestIndex {
				continue
			}
			if (start+int64(i-largestIndex))%v != 0 {
				continue OUT
			}
		}
		return start - int64(largestIndex)
	}
	return -1
}

func findShuttleSequenceOptimized(input []string) (x int64) {
	schedule := make(map[int64]int64)
	valueStrings := strings.Split(input[1], ",")
	for i, s := range valueStrings {
		if s != "x" {
			v, _ := strconv.Atoi(s)
			schedule[int64(i)] = int64(v)
		}
	}
	var inc int64 = schedule[0]
	for i, bus := range schedule {
		if i == 0 {
			continue
		}
		for (x+i)%bus != 0 {
			x += inc
		}
		inc *= bus
	}
	return
}
