/*
Given a binary tree,
check whether it is a mirror of itself (ie, symmetric around its center).
For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
*/
package _01_symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func isEqual(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}
	return isEqual(node1.Left, node2.Right) && isEqual(node1.Right, node2.Left)
}
