/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
        return
    }
	
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head2 := slow.Next
	slow.Next = nil

	// reverse second half
	current, prev := head2, (*ListNode)(nil)
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	reversedHead2 := prev

	var next1, next2 *ListNode
	for head != nil && reversedHead2 != nil {
		next1 = head.Next
		next2 = reversedHead2.Next
		head.Next = reversedHead2
		reversedHead2.Next = next1
		head = next1
		reversedHead2 = next2
	}

	return
}
