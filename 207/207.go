package main

type Graph struct{}

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseGraph := constructGraph(prerequisites)

	var dfs func(course int) bool

	visited := make(map[int]bool)

	dfs = func(course int) bool {
		prereqs := courseGraph[course]
		if _, ok := visited[course]; ok {
			return false
		}
		if len(prereqs) == 0 {
			return true
		}
		visited[course] = true
		for _, prereq := range prereqs {
			if !dfs(prereq) {
				return false
			}
		}
		delete(visited, course)
		courseGraph[course] = []int{}
		return true
	}

	for course := range courseGraph {
		if !dfs(course) {
			return false
		}
	}
	return true
}

func constructGraph(prerequisites [][]int) map[int][]int {
	courseGraph := make(map[int][]int)
	for _, edge := range prerequisites {
		if _, ok := courseGraph[edge[0]]; ok {
			courseGraph[edge[0]] = append(courseGraph[edge[0]], edge[1])
		} else {
			courseGraph[edge[0]] = []int{edge[1]}
		}
	}
	return courseGraph
}
