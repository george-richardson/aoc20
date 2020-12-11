package day11_seating_system

func findOccupiedConways(input *[][]byte) (occupied int) {
	for runConwayFrame(input) {
	}
	for _, row := range *input {
		for _, cell := range row {
			if cell == '#' {
				occupied++
			}
		}
	}
	return
}

func runConwayFrame(input *[][]byte) (madeChange bool) {
	original := *input
	ySize, xSize := len(original), len(original[0])
	facsimile := make([][]byte, ySize)
	for i := range original {
		facsimile[i] = make([]byte, xSize)
		copy(facsimile[i], original[i])
	}
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			if original[y][x] == '.' {
				continue
			}
			occupied, empty, floor := 0, 0, 0
			for innerX := x - 1; innerX < x+2; innerX++ {
				for innerY := y - 1; innerY < y+2; innerY++ {
					if (innerX == x && innerY == y) || innerX < 0 || innerX >= xSize || innerY < 0 || innerY >= ySize {
						continue
					}
					switch facsimile[innerY][innerX] {
					case 'L':
						empty++
					case '#':
						occupied++
					case '.':
						floor++
					}
				}
			}
			if original[y][x] == 'L' && occupied == 0 {
				original[y][x] = '#'
				madeChange = true
			} else if original[y][x] == '#' && occupied >= 4 {
				original[y][x] = 'L'
				madeChange = true
			}
		}
	}
	return
}

func findOccupiedConwaysAxis(input *[][]byte) (occupied int) {
	for runConwayFrameOnAxis(input) {
	}
	for _, row := range *input {
		for _, cell := range row {
			if cell == '#' {
				occupied++
			}
		}
	}
	return
}

func runConwayFrameOnAxis(input *[][]byte) (madeChange bool) {
	original := *input
	ySize, xSize := len(original), len(original[0])
	facsimile := make([][]byte, ySize)
	for i := range original {
		facsimile[i] = make([]byte, xSize)
		copy(facsimile[i], original[i])
	}
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			if original[y][x] == '.' {
				continue
			}
			occupied, _, _ := findSeatingCountsInAllDirections(&facsimile, x, y, xSize, ySize)

			if original[y][x] == 'L' && occupied == 0 {
				original[y][x] = '#'
				madeChange = true
			} else if original[y][x] == '#' && occupied >= 5 {
				original[y][x] = 'L'
				madeChange = true
			}
		}
	}
	return
}

func findSeatingCountsInAllDirections(inputGrid *[][]byte, originX, originY, xSize, ySize int) (occupied, empty, floor int) {
	for deltaX := -1; deltaX <= 1; deltaX++ {
		for deltaY := -1; deltaY <= 1; deltaY++ {
			if deltaX == 0 && deltaY == 0 {
				continue
			}
			switch findNonFloorInDirection(inputGrid, originX, originY, deltaX, deltaY, xSize, ySize) {
			case '#':
				occupied++
			case 'L':
				empty++
			case '.':
				floor++
			}
		}
	}
	return
}

func findNonFloorInDirection(inputGrid *[][]byte, originX, originY, deltaX, deltaY, xSize, ySize int) byte {
	grid := *inputGrid
	for x, y := originX+deltaX, originY+deltaY; x >= 0 && x < xSize && y >= 0 && y < ySize; x, y = x+deltaX, y+deltaY {
		if grid[y][x] != '.' {
			return grid[y][x]
		}
	}
	return '.'
}
