package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	convertBST(root, &sum)
	return root
}

func convertBST(node *TreeNode, sum *int) {
	if node == nil {
		return
	}
	convertBST(node.Right, sum)
	*sum += node.Val
	node.Val = *sum
	convertBST(node.Left, sum)
}
