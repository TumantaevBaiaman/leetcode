package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	count += countNodes(root.Left)
	count += countNodes(root.Right)

	return count
}
