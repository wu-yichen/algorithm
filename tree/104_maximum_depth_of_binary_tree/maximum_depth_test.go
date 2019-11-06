package _04_maximum_depth_of_binary_tree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tree := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	depth := maxDepth(&tree)
	fmt.Println(depth)
	if depth != 3 {
		t.Error("wrong")
	}
}
