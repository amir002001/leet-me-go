package graphvalidtree

func ShutUp() {
	validTree(1, [][]int{})
}

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}
	// make the graph
	graph := make(map[int][]int)

	for i := 0; i < n; i++ {
		graph[n] = []int{}
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	// dfs every elem
	for i := 0; i < n; i++ {
		visited := make(map[int]struct{})
		if dfs(i, -1, &visited, &graph) && len(visited) == n {
			return true
		}
	}

	return false
}

func dfs(node, prevNode int, visited *map[int]struct{}, graph *map[int][]int) bool {
	if _, ok := (*visited)[node]; ok {
		return false
	}
	(*visited)[node] = struct{}{}

	for _, neighbour := range (*graph)[node] {
		if neighbour == prevNode {
			continue
		}
		if !dfs(neighbour, node, visited, graph) {
			return false
		}
	}

	return true
}
