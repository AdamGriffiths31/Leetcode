func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left := head
	right := getMid(head)
	temp := right.Next
	right.Next = nil
	right = temp

	left = sortList(left)
	right = sortList(right)

	return merge(left, right)
}

func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}
		tail = tail.Next

		if left != nil {
			tail.Next = left
		}

		if right != nil {
			tail.Next = right
		}
	}

	return dummy.Next
}