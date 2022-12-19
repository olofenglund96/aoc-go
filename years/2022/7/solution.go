package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olofenglund96/aoc-go/helpers"
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

func FindObjectiveSum(rootNode *TreeNode) int {
	if len(rootNode.children) == 0 {
		return 0
	}
	thisSum := rootNode.subtreeValue
	if thisSum > 100000 {
		thisSum = 0
	}

	for _, c := range rootNode.children {
		thisSum += FindObjectiveSum(c)
	}

	return thisSum
}

func FindClosestSum(rootNode *TreeNode, targetSum int, depth int) int {
	fmt.Printf("%sNode: %s, Value %d, Depth: %d,  Sum: %d\n", strings.Repeat(" ", depth), rootNode.name, rootNode.value, depth, rootNode.subtreeValue)
	if rootNode.subtreeValue < targetSum || len(rootNode.children) == 0 {
		return -1
	}

	fmt.Printf("%sSubtreevalue larger than targetSum\n", strings.Repeat(" ", depth))

	lowestDirDiff := rootNode.subtreeValue - targetSum
	lowestDirSize := rootNode.subtreeValue
	for _, c := range rootNode.children {
		dirSize := FindClosestSum(c, targetSum, depth+1)

		if dirSize == -1 {
			continue
		}

		dirDiff := dirSize - targetSum

		if dirDiff < lowestDirDiff {
			lowestDirDiff = dirDiff
			lowestDirSize = dirSize
		}
	}

	return lowestDirSize
}

func handleCd(rootNode *TreeNode, currNode *TreeNode, to string) *TreeNode {
	if to == "/" {
		return rootNode
	}

	if to == ".." {
		return currNode.parent
	}

	for _, c := range currNode.children {
		if c.name == to {
			return c
		}
	}

	fmt.Printf("Could not find a node to cd to.. currNode: %+v, to: %s\n", currNode, to)
	return nil
}

func addChild(node *TreeNode, out string) {
	size_and_name := strings.Split(out, " ")
	name := size_and_name[1]
	size := 0
	if size_and_name[0] != "dir" {
		size = helpers.StrToI(size_and_name[0])
	}
	node.children = append(node.children, &TreeNode{
		children: []*TreeNode{},
		parent:   node,
		name:     name,
		value:    size,
	})
}

func sol1(rows []string) string {
	rootNode := &TreeNode{
		children: []*TreeNode{},
		parent:   nil,
		name:     "/",
		value:    0,
	}

	currNode := rootNode

	for _, row := range rows {
		if strings.Contains(row, "$ cd") {
			currNode = handleCd(rootNode, currNode, row[5:])
			continue
		}

		if strings.Index(row, "$") != 0 {
			addChild(currNode, row)
		}
	}

	SumTree(rootNode)
	PrintTree(rootNode, 0)

	return fmt.Sprint(FindObjectiveSum(rootNode))
}

func sol2(rows []string) string {
	rootNode := &TreeNode{
		children: []*TreeNode{},
		parent:   nil,
		name:     "/",
		value:    0,
	}

	currNode := rootNode

	for _, row := range rows {
		if strings.Contains(row, "$ cd") {
			currNode = handleCd(rootNode, currNode, row[5:])
			continue
		}

		if strings.Index(row, "$") != 0 {
			addChild(currNode, row)
		}
	}

	SumTree(rootNode)
	PrintTree(rootNode, 0)

	totDiskSpace := 70000000
	neededSpace := 30000000
	unusedDisk := totDiskSpace - rootNode.subtreeValue
	needToRemove := neededSpace - unusedDisk
	println(needToRemove)

	return fmt.Sprint(FindClosestSum(rootNode, needToRemove, 0))
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("years/2022/7/%s.dat", os.Args[2]))

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
