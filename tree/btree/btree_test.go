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

func TestBTreeSearch(t *testing.T) {
	btree := BTree{Degree: 2, Root: tree}
	node := btree.Search(3)
	assert.Equal(t, node, tree.Children[0].Children[1], "BTree.Search")

	node = btree.Search(10)
	assert.Nil(t, node, "BTree.Search")
}

func TestBTreeInsert(t *testing.T) {
	btree := BTree{Degree: 2}
	for i := 1; i < 10; i++ {
		btree.Insert(i)
	}

	assert.Equal(t, btree.Root, tree, "BTree.Insert")
}

func TestBTreeRemove(t *testing.T) {
	btree1 := BTree{Degree: 2}
	btree2 := BTree{Degree: 2}
	for i := 1; i < 10; i++ {
		btree1.Insert(i)
	}

	for i := 9; i > 0; i-- {
		btree1.Remove(i)
	}

	assert.Equal(t, btree1, btree2, "BTree.Remove")
}
