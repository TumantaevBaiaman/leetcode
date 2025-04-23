package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cashHead := make(map[int]bool)
	if head != nil {
		current := head
		prev := head
		cashHead[current.Val] = true
		current = current.Next
		for current != nil {
			if cashHead[current.Val] {
				prev.Next = current.Next
			} else {
				cashHead[current.Val] = true
				prev = current
			}
			current = current.Next
		}
	}
	return head
}
