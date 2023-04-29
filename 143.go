package main

func reorderList(head *ListNode) {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	slow, fast := dummyHead, head

	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}

	secondHead := slow.Next
	slow.Next = nil
	secondHead = reverseList(secondHead)

	dummyHead.Next = nil
	tail := dummyHead

	pickTop := true
	for head != nil || secondHead != nil {
		if pickTop {
			tail.Next = head
			head = head.Next
			pickTop = false
		} else {
			tail.Next = secondHead
			secondHead = secondHead.Next
			pickTop = true
		}
		tail = tail.Next
		tail.Next = nil
	}
}
