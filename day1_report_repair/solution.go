package main

func findSum2020ThenMultiply(in *[]int) int {
	input := *in
	l := len(input)
	for i := 0; i < l; i ++ {
		for j := i + 1; j < l; j ++ {
			if input[i] + input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}
	return 0
}

func findSum2020ThenMultiplyPart2(in *[]int) int {
	input := *in
	l := len(input)
	for i := 0; i < l; i ++ {
		for j := i + 1; j < l; j ++ {
			for k := j + 1; k < l; k++ {
				if input[i] + input[j] + input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return 0
}