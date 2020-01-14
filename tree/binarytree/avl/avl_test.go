package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tree = &TreeNode{
		Value:  4,
		Height: 3,
		Left: &TreeNode{
			Value:  2,
			Height: 1,
			Left:   &TreeNode{Value: 1},
			Right:  &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value:  7,
			Height: 2,
			Left: &TreeNode{
				Value:  6,
				Height: 1,
				Left:   &TreeNode{Value: 5},
			},
			Right: &TreeNode{
				Value:  9,
				Height: 1,
				Left:   &TreeNode{Value: 8},
				Right:  &TreeNode{Value: 10},
			},
		},
	}
)

func TestAvlSearch(t *testing.T) {
	avl := &Avl{Root: tree}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, avl.Search(i).Value, i, "Avl.Search")
	}
}

func TestAvlInsert(t *testing.T) {
	values := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	avl := &Avl{}

	for _, value := range values {
		avl.Insert(value)
	}

	assert.Equal(t, avl.Root, tree, "Avl.Insert")
}

func TestAvlRemove(t *testing.T) {
	var (
		removed = &TreeNode{
			Value:  4,
			Height: 2,
			Left: &TreeNode{
				Value:  2,
				Height: 1,
				Left:   &TreeNode{Value: 1},
				Right:  &TreeNode{Value: 3},
			},
			Right: &TreeNode{
				Value:  6,
				Height: 1,
				Left:   &TreeNode{Value: 5},
				Right:  &TreeNode{Value: 7},
			},
		}
	)
	values := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	avl := &Avl{}

	for _, value := range values {
		avl.Insert(value)
	}

	avl.Remove(8)
	avl.Remove(9)
	avl.Remove(10)

	assert.Equal(t, avl.Root, removed, "avl.Remove")
}
