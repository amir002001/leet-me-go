package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	oldToNew := make(map[*Node]*Node)
	return cloneGraphHelper(node, &oldToNew)
}

func cloneGraphHelper(oldNode *Node, oldToNew *map[*Node]*Node) *Node {
	if copyNode, ok := (*oldToNew)[oldNode]; ok {
		return copyNode
	}
	copyNode := &Node{oldNode.Val, []*Node{}}
	(*oldToNew)[oldNode] = copyNode

	for _, child := range oldNode.Neighbors {
		newChild := cloneGraphHelper(child, oldToNew)
		copyNode.Neighbors = append(copyNode.Neighbors, newChild)
	}

	return copyNode
}
