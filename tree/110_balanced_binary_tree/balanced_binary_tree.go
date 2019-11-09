package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && abs(getMaxDepth(root.Left)-getMaxDepth(root.Right)) < 2
}

func getMaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	maxLeft := getMaxDepth(node.Left)
	maxRight := getMaxDepth(node.Right)
	if maxLeft > maxRight {
		return maxLeft + 1
	}
	return maxRight + 1
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
