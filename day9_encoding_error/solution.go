package day9_encoding_error

import "sort"

func findInvalidXMAS(in []int, preambleSize int) (r int) {
	for i := preambleSize; i < len(in); i++ {
		if !validateLineXMAS(in[i], in[i-preambleSize:i]) {
			return in[i]
		}
	}
	return 0
}

func validateLineXMAS(target int, options []int) bool {
	for i, x := range options {
		for _, y := range options[i+1:] {
			if x+y == target {
				return true
			}
		}
	}
	return false
}

func findEncryptionWeakness(in []int, target int) int {
	lower := 0
	upper := 1
	total := in[0] + in[1]
	for true {
		if total == target {
			break
		}
		if total < target {
			upper++
			total += in[upper]
		} else {
			total -= in[lower]
			lower++
		}
	}
	sorted := in[lower:upper]
	sort.Ints(sorted)
	return sorted[0] + sorted[len(sorted)-1]
}
