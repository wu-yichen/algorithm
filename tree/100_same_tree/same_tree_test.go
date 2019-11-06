package _00_same_tree

import "testing"

func TestIsSameTree(t *testing.T) {
	tree1 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	tree2 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	tree3 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	tree4 := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	tree5 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	tree6 := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	tests := []struct {
		name   string
		tree   []TreeNode
		expect bool
	}{
		{
			name:   "test1",
			tree:   []TreeNode{tree1, tree2},
			expect: true,
		},
		{
			name:   "test2",
			tree:   []TreeNode{tree3, tree4},
			expect: false,
		},
		{
			name:   "test3",
			tree:   []TreeNode{tree5, tree6},
			expect: false,
		},
	}
	for _, value := range tests {
		isSame := isSameTree(&value.tree[0], &value.tree[1])
		if isSame != value.expect {
			t.Errorf("test %s is wrong", value.name)
		}
	}
}
