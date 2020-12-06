package day6_custom_customs

import (
	"bufio"
)

func findCustomsSum(input *bufio.Scanner) (r int) {
	answers := make(map[byte]int)
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			r += len(answers)
			answers = map[byte]int{}
			continue
		}
		for _, char := range []byte(line) {
			answers[char]++
		}
	}
	r += len(answers)
	return
}

func findCustomsSumPart2(input *bufio.Scanner) (r int) {
	answers := make(map[byte]int)
	groupSize := 0
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			r += findGroupSumEveryone(&answers, groupSize)
			answers = map[byte]int{}
			groupSize = 0
			continue
		}
		groupSize++
		for _, char := range []byte(line) {
			answers[char]++
		}
	}
	r += findGroupSumEveryone(&answers, groupSize)
	return
}

func findGroupSumEveryone(answerCounts *map[byte]int, groupSize int) (r int) {
	for _, v := range *answerCounts {
		if v == groupSize {
			r++
		}
	}
	return
}
