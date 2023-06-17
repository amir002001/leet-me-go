package threetwothree

func countComponents(n int, edges [][]int) int {
	compGraph := make([][]int, n)
	for i := 0; i < n; i++ {
		compGraph[i] = []int{}
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		compGraph[a] = append(compGraph[a], b)
		compGraph[b] = append(compGraph[b], a)
	}

	compCount := 0

	visited := make(map[int]bool)
	var dfs func(node int)
	dfs = func(node int) {
		if _, ok := visited[node]; ok {
			return
		}
		visited[node] = true

		for _, connection := range compGraph[node] {
			dfs(connection)
		}
	}
	for node := range compGraph {
		if _, ok := visited[node]; ok {
			continue
		}
		compCount++
		dfs(node)
	}

	return compCount
}
