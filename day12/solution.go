package day12

import (
	"math"
	"strconv"
)

func findDestinationManhattanDistance(directions []string) int {
	bearings := map[int]byte{
		0:   'N',
		90:  'E',
		180: 'S',
		270: 'W',
	}
	bearing := 90
	var NS, EW int
	for _, direction := range directions {
		instruction := direction[0]
		value, _ := strconv.Atoi(direction[1:])
		if instruction == 'F' {
			instruction = bearings[bearing]
		}
		switch instruction {
		case 'N':
			NS += value
		case 'S':
			NS -= value
		case 'E':
			EW += value
		case 'W':
			EW -= value
		case 'L':
			bearing = (bearing + (360 - value)) % 360
		case 'R':
			bearing = (bearing + value) % 360
		}
	}
	return int(math.Abs(float64(NS)) + math.Abs(float64(EW)))
}

func findDestinationManhattanDistanceWaypoint(directions []string) int {
	var NS, EW int
	WNS, WEW := 1, 10
	turn := func(degreesRight int) {
		for ; degreesRight > 0; degreesRight -= 90 {
			WNS, WEW = WEW*-1, WNS
		}
	}
	for _, direction := range directions {
		instruction := direction[0]
		value, _ := strconv.Atoi(direction[1:])
		switch instruction {
		case 'N':
			WNS += value
		case 'S':
			WNS -= value
		case 'E':
			WEW += value
		case 'W':
			WEW -= value
		case 'L':
			turn(360 - value)
		case 'R':
			turn(value)
		case 'F':
			NS += WNS * value
			EW += WEW * value
		}
	}
	return int(math.Abs(float64(NS)) + math.Abs(float64(EW)))
}
