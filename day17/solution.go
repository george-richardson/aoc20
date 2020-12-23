package day17

func max(a, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func min(a, b int) (min int) {
	if a < b {
		min = a
	} else {
		min = b
	}
	return
}

type coordinate struct {
	x, y, z int
}

func maxCoordinate(a, b coordinate) (result coordinate) {
	result = coordinate{
		x: max(a.x, b.x+1),
		y: max(a.y, b.y+1),
		z: max(a.z, b.z+1),
	}
	return
}

func minCoordinate(a, b coordinate) (result coordinate) {
	result = coordinate{
		x: min(a.x, b.x-1),
		y: min(a.y, b.y-1),
		z: min(a.z, b.z-1),
	}
	return
}

type cubeSpace struct {
	space                map[coordinate]bool
	maxCorner, minCorner coordinate
	active               int
}

func (space *cubeSpace) addCube(coord coordinate) {
	space.maxCorner = maxCoordinate(space.maxCorner, coord)
	space.minCorner = minCoordinate(space.minCorner, coord)
	space.space[coord] = false
}

func (space *cubeSpace) activate(coord coordinate) {
	if _, exists := space.space[coord]; !exists {
		space.addCube(coord)
	}
	space.space[coord] = true
	space.active++
}

func (space *cubeSpace) deactivate(coord coordinate) {
	space.space[coord] = false
	space.active--
}

func countActiveAdjacentCubes(space *map[coordinate]bool, coord coordinate) (active int) {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				newCoord := coordinate{
					x: coord.x + x,
					y: coord.y + y,
					z: coord.z + z,
				}

				if newCoord != coord && (*space)[newCoord] {
					active++
				}
			}
		}
	}
	return
}

func (space *cubeSpace) runCycle() {
	var oldSpace = make(map[coordinate]bool)
	for key, value := range space.space {
		oldSpace[key] = value
	}

	for x := space.minCorner.x; x <= space.maxCorner.x; x++ {
		for y := space.minCorner.y; y <= space.maxCorner.y; y++ {
			for z := space.minCorner.z; z <= space.maxCorner.z; z++ {
				coord := coordinate{
					x: x,
					y: y,
					z: z,
				}

				nearbyActive := countActiveAdjacentCubes(&oldSpace, coord)

				if oldSpace[coord] && !(nearbyActive == 2 || nearbyActive == 3) {
					space.deactivate(coord)
				} else if !oldSpace[coord] && nearbyActive == 3 {
					space.activate(coord)
				}
			}
		}
	}
}

func (space *cubeSpace) init(startingGrid [][]byte) {
	space.active = 0
	space.space = make(map[coordinate]bool)
	space.maxCorner = coordinate{}
	space.minCorner = coordinate{}

	for y, line := range startingGrid {
		for x, char := range line {
			coord := coordinate{x, y, 0}
			space.addCube(coord)
			if char == '#' {
				space.activate(coord)
			}
		}
	}
}

func bootCubeSpace(input [][]byte, cycles int) (active int) {
	space := cubeSpace{}
	space.init(input)

	for i := 0; i < cycles; i++ {
		space.runCycle()
	}

	return space.active
}
