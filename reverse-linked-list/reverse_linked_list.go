package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	current := head
	for head.Next != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current.Next = next
	}
	return pre
}
