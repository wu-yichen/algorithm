package binary_tree_level_order_traversal_2

import (
	"fmt"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
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

	actual := levelOrderBottom(&tree)
	for _, value := range actual {
		for _, v := range value {
			fmt.Print(v, ",")
		}
		fmt.Println()

	}
}
