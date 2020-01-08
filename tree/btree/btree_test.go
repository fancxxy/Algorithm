package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tree = &TreeNode{
		Elements: []int{4},
		Children: []*TreeNode{
			&TreeNode{
				Elements: []int{2},
				Children: []*TreeNode{
					&TreeNode{
						Elements: []int{1},
						Children: []*TreeNode{},
						Leaf:     true,
					},
					&TreeNode{
						Elements: []int{3},
						Children: []*TreeNode{},
						Leaf:     true,
					},
				},
				Leaf: false,
			},
			&TreeNode{
				Elements: []int{6},
				Children: []*TreeNode{
					&TreeNode{
						Elements: []int{5},
						Children: []*TreeNode{},
						Leaf:     true,
					},
					&TreeNode{
						Elements: []int{7, 8, 9},
						Children: []*TreeNode{},
						Leaf:     true,
					},
				},
				Leaf: false,
			},
		},
		Leaf: false,
	}
)

func TestBTreeInsert(t *testing.T) {
	btree := BTree{Degree: 2}
	for i := 1; i < 10; i++ {
		btree.Insert(i)
	}

	assert.Equal(t, btree.Root, tree, "BTree.Insert")
}
