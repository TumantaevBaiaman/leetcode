package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter := 0
	depth(root, &diameter)
	return diameter
}

func depth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	leftDepth := depth(node.Left, diameter)
	rightDepth := depth(node.Right, diameter)
	*diameter = int(math.Max(float64(*diameter), float64(leftDepth+rightDepth)))
	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}
