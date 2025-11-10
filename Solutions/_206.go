func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		nextNode := current.Next
		current.Next = previous
		previous = current
		current = nextNode
	}
	return previous
}
