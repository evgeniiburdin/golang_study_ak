package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	var sum int

	for len(queue) > 0 {
		levelSize := len(queue)
		sum = 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return sum
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 8}

	sum := deepestLeavesSum(root)
	fmt.Println("Sum of deepest leaves:", sum)
}
