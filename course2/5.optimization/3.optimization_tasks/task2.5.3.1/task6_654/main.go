package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	suffix, prefix, newNodeValue := separateTree(nums)
	node := &TreeNode{
		Val:   newNodeValue,
		Left:  constructMaximumBinaryTree(prefix),
		Right: constructMaximumBinaryTree(suffix),
	}

	return node
}

func separateTree(nums []int) (suffix, prefix []int, newNodeValue int) {
	if len(nums) == 0 {
		return []int{}, []int{}, 0
	}

	suffix = make([]int, 0)
	prefix = make([]int, 0)

	newNodeValue, newNodeValueIndex := findMaxValueAndItsIndex(nums)

	for i := 0; i < len(nums); i++ {
		if i < newNodeValueIndex {
			prefix = append(prefix, nums[i])
		}
		if i > newNodeValueIndex {
			suffix = append(suffix, nums[i])
		}
	}

	return
}

func findMaxValueAndItsIndex(nums []int) (maxValue, maxValueIndex int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxValueIndex = i
		}
	}

	return
}
