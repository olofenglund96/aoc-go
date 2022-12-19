package helpers

import (
	"fmt"
	"strings"
)

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	children     []*TreeNode
	parent       *TreeNode
	name         string
	value        int
	subtreeValue int
}

func (t TreeNode) AddChild(childNode *TreeNode) {
	t.children = append(t.children, childNode)
}

func (t TreeNode) SubtreeValue() int {
	val := t.value
	for _, childNode := range t.children {
		val += childNode.SubtreeValue()
	}

	return val
}

func PrintTree(rootNode *TreeNode, depth int) {
	fmt.Printf("%sNode: %s, Value %d, Depth: %d,  Sum: %d\n", strings.Repeat(" ", depth), rootNode.name, rootNode.value, depth, rootNode.subtreeValue)
	for _, c := range rootNode.children {
		PrintTree(c, depth+1)
	}
}

func SumTree(rootNode *TreeNode) int {
	thisSum := rootNode.value
	for _, c := range rootNode.children {
		thisSum += SumTree(c)
	}

	rootNode.subtreeValue = thisSum

	return thisSum
}
