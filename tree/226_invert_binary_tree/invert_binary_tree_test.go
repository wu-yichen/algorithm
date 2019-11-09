package invert_binary_tree

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tree := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	node := invertTree(&tree)
	fmt.Println(node.Val)
	fmt.Println(node.Left.Val)
	fmt.Println(node.Right.Val)
	fmt.Println(node.Left.Left.Val)
	fmt.Println(node.Left.Right.Val)
	fmt.Println(node.Right.Left.Val)
	fmt.Println(node.Right.Right.Val)
}
