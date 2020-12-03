package day3_toboggan_trajectory

func checkSlopePath(slope *[][]byte) (treesHit int) {
	grid := *slope
	height := len(grid)
	width := len(grid[0])
	for x, y := 0, 0; y < height; x, y = (x+3)%width, y+1 {
		if grid[y][x] == '#' {
			treesHit++
		}
	}
	return
}

func checkSlopePaths(slope *[][]byte) (treesHit int) {
	treesHit = 1
	var trajectories = []struct {
		x int
		y int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, t := range trajectories {
		treesHit *= checkSlopePathPart2(slope, t.x, t.y)
	}
	return
}

func checkSlopePathPart2(slope *[][]byte, xDiff, yDiff int) (treesHit int) {
	grid := *slope
	height := len(grid)
	width := len(grid[0])
	for x, y := 0, 0; y < height; x, y = (x+xDiff)%width, y+yDiff {
		if grid[y][x] == '#' {
			treesHit++
		}
	}
	return
}
