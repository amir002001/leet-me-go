package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var prev *ListNode
	curr := head
	next := head.Next
	for next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = next.Next
	}
	curr.Next = prev
	return curr
}
