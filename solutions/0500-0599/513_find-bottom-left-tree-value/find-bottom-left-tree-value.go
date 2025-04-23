package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	var leftMostValue int

	for queue.Len() > 0 {
		size := queue.Len()
		leftMostValue = queue.Front().Value.(*TreeNode).Val

		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return leftMostValue
}
