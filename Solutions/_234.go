func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	var rev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nextSlow := slow.Next
		slow.Next = rev
		rev = slow
		slow = nextSlow
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != rev.Val {
			return false
		}
		slow = slow.Next
		rev = rev.Next
	}

	return true
}