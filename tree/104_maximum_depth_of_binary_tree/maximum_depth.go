/*
Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/
package _04_maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// breadth first search -  use queue !!
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		num := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[num:]
		level++
	}
	return level
}

// depth first search - recursive or stack !!
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	maxLeft := maxDepth(root.Left)
//	maxRight := maxDepth(root.Right)
//	if maxLeft > maxRight {
//		return maxLeft + 1
//	}
//	return maxRight + 1
//}
