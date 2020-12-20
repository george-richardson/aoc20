package day14

import (
	"math"
	"strconv"
	"strings"
)

func runMasksProgram(program []string) (total uint64) {
	memory := make(map[int]uint64)
	var ones, zeros uint64
	for _, line := range program {
		assignment := strings.Split(line, " = ")
		if assignment[0] == "mask" {
			ones, zeros = parseMask(assignment[1])
		} else {
			memoryLocation, _ := strconv.Atoi(assignment[0][4 : len(assignment[0])-1])
			value, _ := strconv.Atoi(assignment[1])
			valueUint := uint64(value)
			memory[memoryLocation] = (valueUint | ones) & zeros
		}
	}
	for _, v := range memory {
		total += v
	}
	return
}

func parseMask(mask string) (ones uint64, zeros uint64) {
	zeros = math.MaxUint64
	for i, c := range []byte(mask) {
		if c == '1' {
			ones |= 1 << (35 - i)
		} else if c == '0' {
			zeros &^= 1 << (35 - i)
		}
	}
	return
}

func decodeMask(mask string) (zeroes, ones, xes int64) {
	for i, char := range mask {
		switch char {
		case '0':
			zeroes |= 1 << (35 - i)
		case '1':
			ones |= 1 << (35 - i)
		case 'X':
			xes |= 1 << (35 - i)
		}
	}
	return
}

func applyMask(zeroes, ones, xes, address int64) (addresses []int64) {
	if xes == 0 {
		return []int64{address | ones}
	}
	var sigBit int64 = 1 << 35
	for xes&sigBit == 0 {
		sigBit = sigBit >> 1
	}
	xes &^= sigBit
	// Maybe set ones and zeros here instead
	return append(applyMask(zeroes, ones, xes, address|sigBit), applyMask(zeroes, ones, xes, address&^sigBit)...)
}

func runMasksProgramPar2(program []string) (total int64) {
	memory := make(map[int64]int64)
	var zeros, ones, xes int64
	for _, line := range program {
		assignment := strings.Split(line, " = ")
		if assignment[0] == "mask" {
			zeros, ones, xes = decodeMask(assignment[1])
		} else {
			memoryLocation, _ := strconv.ParseInt(assignment[0][4:len(assignment[0])-1], 10, 64)
			value, _ := strconv.ParseInt(assignment[1], 10, 64)
			addresses := applyMask(zeros, ones, xes, memoryLocation)
			for _, address := range addresses {
				memory[address] = value
			}
		}
	}
	for _, v := range memory {
		total += v
	}
	return
}
