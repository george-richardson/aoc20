package day17

type hyperCoordinate struct {
	x, y, z, t int
}

func maxHyperCoordinate(a, b hyperCoordinate) (result hyperCoordinate) {
	result = hyperCoordinate{
		x: max(a.x, b.x+1),
		y: max(a.y, b.y+1),
		z: max(a.z, b.z+1),
		t: max(a.t, b.t+1),
	}
	return
}

func minHyperCoordinate(a, b hyperCoordinate) (result hyperCoordinate) {
	result = hyperCoordinate{
		x: min(a.x, b.x-1),
		y: min(a.y, b.y-1),
		z: min(a.z, b.z-1),
		t: min(a.t, b.t-1),
	}
	return
}

type hyperCubeSpace struct {
	space                map[hyperCoordinate]bool
	maxCorner, minCorner hyperCoordinate
	active               int
}

func (space *hyperCubeSpace) addCube(coord hyperCoordinate) {
	space.maxCorner = maxHyperCoordinate(space.maxCorner, coord)
	space.minCorner = minHyperCoordinate(space.minCorner, coord)
	space.space[coord] = false
}

func (space *hyperCubeSpace) activate(coord hyperCoordinate) {
	if _, exists := space.space[coord]; !exists {
		space.addCube(coord)
	}
	space.space[coord] = true
	space.active++
}

func (space *hyperCubeSpace) deactivate(coord hyperCoordinate) {
	space.space[coord] = false
	space.active--
}

func countActiveAdjacentHyperCubes(space *map[hyperCoordinate]bool, coord hyperCoordinate) (active int) {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				for t := -1; t < 2; t++ {
					newCoord := hyperCoordinate{
						x: coord.x + x,
						y: coord.y + y,
						z: coord.z + z,
						t: coord.t + t,
					}

					if newCoord != coord && (*space)[newCoord] {
						active++
					}
				}
			}
		}
	}
	return
}

func (space *hyperCubeSpace) runCycle() {
	var oldSpace = make(map[hyperCoordinate]bool)
	for key, value := range space.space {
		oldSpace[key] = value
	}

	for x := space.minCorner.x; x <= space.maxCorner.x; x++ {
		for y := space.minCorner.y; y <= space.maxCorner.y; y++ {
			for z := space.minCorner.z; z <= space.maxCorner.z; z++ {
				for t := space.minCorner.t; t <= space.maxCorner.t; t++ {
					coord := hyperCoordinate{
						x: x,
						y: y,
						z: z,
						t: t,
					}

					nearbyActive := countActiveAdjacentHyperCubes(&oldSpace, coord)

					if oldSpace[coord] && !(nearbyActive == 2 || nearbyActive == 3) {
						space.deactivate(coord)
					} else if !oldSpace[coord] && nearbyActive == 3 {
						space.activate(coord)
					}
				}
			}
		}
	}
}

func (space *hyperCubeSpace) init(startingGrid [][]byte) {
	space.active = 0
	space.space = make(map[hyperCoordinate]bool)
	space.maxCorner = hyperCoordinate{}
	space.minCorner = hyperCoordinate{}

	for y, line := range startingGrid {
		for x, char := range line {
			coord := hyperCoordinate{x, y, 0, 0}
			space.addCube(coord)
			if char == '#' {
				space.activate(coord)
			}
		}
	}
}

func bootHyperCubeSpace(input [][]byte, cycles int) (active int) {
	space := hyperCubeSpace{}
	space.init(input)

	for i := 0; i < cycles; i++ {
		space.runCycle()
	}

	return space.active
}
