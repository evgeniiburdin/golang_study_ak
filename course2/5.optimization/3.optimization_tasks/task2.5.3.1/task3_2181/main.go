package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	curNode := head
	for curNode.Next != nil {
		for curNode.Next.Val != 0 {
			curNode.Val += curNode.Next.Val
			curNode.Next = curNode.Next.Next
		}
		if curNode.Next.Next == nil {
			curNode.Next = nil
			break
		}
		curNode = curNode.Next
	}
	return head
}
