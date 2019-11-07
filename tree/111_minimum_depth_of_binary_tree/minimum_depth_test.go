package minimum_depth_of_binary_tree

import (
	"testing"
)

func TestMinDepth(t *testing.T) {
	tree := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   8,
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
	tree2 := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	tests := []struct {
		name   string
		t      *TreeNode
		expect int
	}{
		{
			"test1",
			&tree,
			2,
		},
		{
			"test2",
			&tree2,
			2,
		},
	}

	for _, value := range tests {
		depth := minDepth(value.t)
		if depth != value.expect {
			t.Error("wrong")
		}
	}
}
