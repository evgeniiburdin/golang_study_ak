package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}

	// Найдем середину списка
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Развернем вторую половину списка
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// Найдем максимальную сумму близнецов
	maxSum := 0
	first, second := head, prev
	for second != nil {
		sum := first.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}
		first = first.Next
		second = second.Next
	}

	return maxSum
}
