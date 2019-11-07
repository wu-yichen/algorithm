package search_in_a_binary_search_tree

import (
	"testing"
)

func TestSearchBST(t *testing.T) {
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
				Val:   8,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val:   79,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   78,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	bst := searchBST(&tree, 9)
	if bst.Val != 9 ||
		bst.Left.Val != 79 ||
		bst.Right.Val != 78 {
		t.Error("wrong")
	}
}
