package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var treeNodesAmount int
	getTreeNodesAmount(root, &treeNodesAmount)

	values := make([]int, 0, treeNodesAmount)
	treeToSlice(root, &values)
	sort.Ints(values)

	return binarySearchTree(values)
}

func getTreeNodesAmount(root *TreeNode, nodesAmount *int) {
	if root == nil {
		return
	}
	*nodesAmount++
	getTreeNodesAmount(root.Left, nodesAmount)
	getTreeNodesAmount(root.Right, nodesAmount)
}

func treeToSlice(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}
	*slice = append(*slice, root.Val)
	treeToSlice(root.Left, slice)
	treeToSlice(root.Right, slice)
}

func binarySearchTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   values[len(values)/2],
		Left:  binarySearchTree(values[:len(values)/2]),
		Right: binarySearchTree(values[len(values)/2+1:])}
}
