package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
                  4
                 / \
                /   \
               2     \
              / \     8
             1   3   / \
                    /   \
                   6     \
                  / \     10
				 5   7   /
				        9
*/

var (
	tree = &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value: 8,
			Left: &TreeNode{
				Value: 6,
				Left:  &TreeNode{Value: 5},
				Right: &TreeNode{Value: 7},
			},
			Right: &TreeNode{
				Value: 10,
				Left:  &TreeNode{Value: 9},
			},
		},
	}
)

func TestBstSearch(t *testing.T) {
	bst := &Bst{Root: tree}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, bst.Search(i).Value, i, "Bst.Search")
	}
}
func TestBstInsert(t *testing.T) {
	values := []int{4, 2, 8, 1, 3, 10, 6, 5, 7, 9}
	bst := &Bst{}

	for _, value := range values {
		bst.Insert(value)
	}

	assert.Equal(t, bst.Root, tree, "Bst.Insert")
}

func TestBstDelete(t *testing.T) {
	deleted := &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value: 9,
			Left: &TreeNode{
				Value: 6,
				Left:  &TreeNode{Value: 5},
				Right: &TreeNode{Value: 7},
			},
			Right: &TreeNode{
				Value: 10,
			},
		},
	}

	values := []int{4, 2, 8, 1, 3, 10, 6, 5, 7, 9}
	bst := &Bst{}

	for _, value := range values {
		bst.Insert(value)
	}

	bst.Delete(8)
	assert.Equal(t, bst.Root, deleted, "Bst.Delete")
}
