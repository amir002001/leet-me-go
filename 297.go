package main

import (
	"strconv"
	"strings"
)

type Codec struct{}

// func Constructor() Codec {
// 	return Codec{}
// }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "dudechamp"
	}
	strs := []string{}
	dfs271(root, &strs)

	var builder strings.Builder
	for i := 0; i < len(strs)-1; i++ {
		builder.Write([]byte(strs[i]))
		builder.WriteRune(',')
	}
	builder.Write([]byte(strs[len(strs)-1]))
	return builder.String()
}

func dfs271(root *TreeNode, strs *[]string) {
	if root == nil {
		*strs = append(*strs, "dudechamp")
		return
	}
	*strs = append(*strs, strconv.Itoa(root.Val))
	dfs271(root.Left, strs)
	dfs271(root.Right, strs)
}

func dfs271deser(strs *[]string, counter *int) *TreeNode {
	root := new(TreeNode)
	val := (*strs)[len(*strs)-*counter]
	if val == "dudechamp" {
		*counter--
		return nil

	}
	num, err := strconv.Atoi(val)
	if err != nil {
		panic(1)
	}
	root.Val = num
	*counter--
	if *counter > 0 {
		root.Left = dfs271deser(strs, counter)
		root.Right = dfs271deser(strs, counter)
	}
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	count := len(strs)
	root := dfs271deser(&strs, &count)
	return root
}
