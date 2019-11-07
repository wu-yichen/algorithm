/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.
*/
package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left) + 1
	right := minDepth(root.Right) + 1

	if left == 1 && right != 1 {
		return right
	}
	if right == 1 && left != 1 {
		return left
	}
	if left < right {
		return left
	}
	return right
}
