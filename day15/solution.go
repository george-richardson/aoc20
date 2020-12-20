package day15

import (
	"strconv"
	"strings"
)

func elfGame(seed string, targetNumber int) (lastSpoken int) {
	valStrings := strings.Split(seed, ",")
	vals := make([]int, len(valStrings))
	for i, valString := range valStrings {
		vals[i], _ = strconv.Atoi(valString)
	}
	var tracker = make(map[int]int)
	lastSpoken = vals[0]
	for i, val := range vals {
		tracker[lastSpoken] = i + 1
		lastSpoken = val
	}

	for i := len(valStrings) + 1; i <= targetNumber; i++ {
		if tracker[lastSpoken] == 0 {
			lastSpoken, tracker[lastSpoken] = 0, i
		} else {
			lastSpoken, tracker[lastSpoken] = i-tracker[lastSpoken], i
		}
	}

	return
}
