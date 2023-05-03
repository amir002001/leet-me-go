package main

import "container/list"

type QueueItem struct {
	depth int
	node  *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := list.New()
	queue.PushFront(QueueItem{0, root})

	levelMap := make(map[int][]int)
	for queue.Len() != 0 {
		curr := queue.Remove(queue.Back()).(QueueItem)
		if curr.node == nil {
			continue
		}
		levelMap[curr.depth] = append(levelMap[curr.depth], curr.node.Val)
		queue.PushFront(QueueItem{curr.depth + 1, curr.node.Left})
		queue.PushFront(QueueItem{curr.depth + 1, curr.node.Right})
	}
	out := make([][]int, len(levelMap))
	for i := 0; i < len(out); i++ {
		out[i] = levelMap[i]
	}
	return out
}
