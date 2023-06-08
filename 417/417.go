/*
we need to find a common set between cells that reach the pacific ocean and cells
that reach the atlantic ocean. To do this successfully, we need to do the following:

cells on left and top are pacific
cells on right and bottom are atlantic

we want to dfs for each cell bordering pacific:
if through dfs we find a cell that is bigger or equal to caller cell we will
add it to the set.

If not we will ignore and return.
likewise we do the same thing for the other set

we get union of both sets

return that union turned into an array.
*/

package main

import "math"

type Cell struct {
	row int
	col int
}

func pacificAtlantic(heights [][]int) [][]int {
	reachesPacific := make(map[Cell]bool)
	reachesAtlantic := make(map[Cell]bool)

	for row := 0; row < len(heights); row++ {
		cellDfs(heights, &reachesPacific, Cell{row, 0})
		cellDfs(heights, &reachesAtlantic, Cell{row, len(heights[0]) - 1})
	}
	for col := 0; col < len(heights[0]); col++ {
		cellDfs(heights, &reachesPacific, Cell{0, col})
		cellDfs(heights, &reachesAtlantic, Cell{len(heights) - 1, col})
	}

	intersection := [][]int{}

	for k := range reachesAtlantic {
		if reachesPacific[k] {
			intersection = append(intersection, []int{k.row, k.col})
		}
	}

	return intersection
}

func cellDfs(heights [][]int, reaches *map[Cell]bool, startingCell Cell) {
	visited := make(map[Cell]bool)
	var dfsHelper func(Cell, int)
	dfsHelper = func(cell Cell, prevValue int) {
		if !boundsOk(heights, cell) || visited[cell] {
			return
		}
		if heights[cell.row][cell.col] < prevValue {
			return
		}
		visited[cell] = true
		(*reaches)[cell] = true

		nextCells := []Cell{
			{row: cell.row + 1, col: cell.col},
			{row: cell.row - 1, col: cell.col},
			{row: cell.row, col: cell.col + 1},
			{row: cell.row, col: cell.col - 1},
		}

		for _, nextCell := range nextCells {
			dfsHelper(nextCell, heights[cell.row][cell.col])
		}
	}
	dfsHelper(startingCell, math.MinInt)
}

func boundsOk(heights [][]int, cell Cell) bool {
	return cell.row != -1 && cell.col != -1 &&
		cell.row != len(heights) && cell.col != len(heights[0])
}
