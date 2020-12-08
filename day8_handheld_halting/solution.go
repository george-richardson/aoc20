package day8_handheld_halting

import (
	"fmt"
	"strconv"
)

func runProgramFindCycle(programInput *[]string) (acc int) {
	program := *programInput
	l := len(program)
	visited := make([]bool, l)
	for i := 0; !visited[i]; {
		visited[i] = true
		op, sign, valString := program[i][:3], program[i][4], program[i][5:]
		val, _ := strconv.Atoi(valString)
		if sign == '-' {
			val *= -1
		}
		switch op {
		case "acc":
			acc += val
		case "jmp":
			i += val
			continue
		}
		i++
	}
	return
}

func patchProgram(programInput *[]string) (acc int) {
	program := *programInput
	l := len(program)
	for i := 0; i < l; i++ {
		op := program[i][:3]
		if op == "nop" || op == "jmp" {
			tmp := make([]string, l)
			copy(tmp, program)
			if op == "nop" {
				tmp[i] = "jmp" + tmp[i][3:]
			} else {
				tmp[i] = "nop" + tmp[i][3:]
			}
			acc, err := runProgramWithPatch(&tmp)
			if err != nil {
				continue
			}
			return acc
		}
	}
	return -1
}

func runProgramWithPatch(programInput *[]string) (acc int, err error) {
	program := *programInput
	l := len(program)
	visited := make([]bool, l)
	for i := 0; i < l; {
		if visited[i] {
			return -1, fmt.Errorf("found cycle at %v", i)
		}
		visited[i] = true
		op, sign, valString := program[i][:3], program[i][4], program[i][5:]
		val, _ := strconv.Atoi(valString)
		if sign == '-' {
			val *= -1
		}
		switch op {
		case "acc":
			acc += val
		case "jmp":
			i += val
			continue
		}
		i++
	}
	return
}
