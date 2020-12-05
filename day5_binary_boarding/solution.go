package day5_binary_boarding

import "sort"

func findMaxSeatLocation(input *[]string) (r int) {
	for _, line := range *input {
		_, _, id := findSeatLocation(line)
		if id > r {
			r = id
		}
	}
	return
}

func findSeatLocation(input string) (row, column, id int) {
	l := len(input)
	row = binarySearch(input[:l-3], 128, 'F', 'B')
	column = binarySearch(input[l-3:], 8, 'L', 'R')
	id = (row * 8) + column
	return
}

func binarySearch(input string, size int, lowerIndicator, upperIndicator byte) (r int) {
	lowerBound, upperBound := 0, size-1
	for _, indicator := range []byte(input) {
		r = ((upperBound - lowerBound) / 2) + lowerBound
		if indicator == lowerIndicator {
			upperBound = r
		} else if indicator == upperIndicator {
			r++
			lowerBound = r
		}
	}
	return
}

func findMissingSeat(input *[]string) int {

	var seats []struct {
		row, col, id int
	}

	for _, line := range *input {
		row, col, id := findSeatLocation(line)
		seats = append(seats, struct{ row, col, id int }{row: row, col: col, id: id})
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].id < seats[j].id
	})

	for i, seat := range seats {
		nextSeat := seats[i+1]
		if nextSeat.id == seat.id+2 {
			return seat.id + 1
		}
	}

	return -1
}
