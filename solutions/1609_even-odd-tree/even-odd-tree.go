package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(root)
	level := 0

	for queue.Len() > 0 {
		size := queue.Len()
		prevVal := 0

		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if level%2 == 0 {
				if node.Val%2 == 0 || (prevVal != 0 && node.Val <= prevVal) {
					return false
				}
			} else {
				if node.Val%2 == 1 || (prevVal != 0 && node.Val >= prevVal) {
					return false
				}
			}

			prevVal = node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		level++
	}

	return true
}
