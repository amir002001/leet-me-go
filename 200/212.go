package main

type Point struct {
	row    int
	column int
}

const (
	land  byte = '1'
	water byte = '0'
)

func numIslands(grid [][]byte) int {
	globalVisited := make(map[Point]bool)
	islandCount := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			point := Point{row, col}

			if _, ok := globalVisited[point]; grid[point.row][point.column] == land && !ok {
				dfsIsland(grid, &globalVisited, point)
				islandCount++
			}
		}
	}
	return islandCount
}

func dfsIsland(grid [][]byte, globalVisited *map[Point]bool, point Point) {
	_, pointVisited := (*globalVisited)[point]
	if !boundsOk(point, grid) || grid[point.row][point.column] == water || pointVisited {
		return
	}

	(*globalVisited)[point] = true

	nextPoints := []Point{
		{point.row, point.column - 1},
		{point.row, point.column + 1},
		{point.row - 1, point.column},
		{point.row + 1, point.column},
	}

	for _, nextPoint := range nextPoints {
		dfsIsland(grid, globalVisited, nextPoint)
	}
}

func boundsOk(point Point, grid [][]byte) bool {
	return !(point.row == -1 || point.row == len(grid) ||
		point.column == -1 || point.column == len(grid[0]))
}
