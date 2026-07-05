/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	fast, slow := head.Next.Next, head.Next
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
