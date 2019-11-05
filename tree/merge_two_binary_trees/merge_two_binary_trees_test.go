package merge_two_binary_trees

import (
	"testing"
)

func TestMergeTrees(t *testing.T) {
	tree1 := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   13,
			Left:  nil,
			Right: nil,
		},
	}
	tree2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	mergeTrees(&tree1, &tree2)
	if tree1.Val != 3 ||
		tree1.Left.Val != 6 ||
		tree1.Right.Val != 16 {
		t.Error("wrong")
	}
}
