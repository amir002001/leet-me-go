package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := head
	finder := tail
	for i := 0; i < n+1; i++ {
		if finder == nil {
			next := head.Next
			head.Next = nil
			return next
		}
		finder = finder.Next
	}

	for finder != nil {
		tail = tail.Next
		finder = finder.Next
	}

	toBeRemoved := tail.Next
	after := toBeRemoved.Next
	tail.Next = after
	toBeRemoved.Next = nil

	return head
}
