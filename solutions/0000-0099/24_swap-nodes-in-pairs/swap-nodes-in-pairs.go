package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next

		prev.Next, second.Next, first.Next = second, first, second.Next

		prev = first
	}

	return dummy.Next
}
