/*
Given a binary tree,
return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
*/
package binary_tree_level_order_traversal_2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	nodes := []*TreeNode{root}
	var temp [][]int
	for len(nodes) > 0 {
		num := len(nodes)
		var n []int
		for _, node := range nodes {
			n = append(n, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		nodes = nodes[num:]
		temp = append(temp, n)
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}
	return result
}
