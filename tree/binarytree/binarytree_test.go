package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	in    = []int{3, 2, 4, 1, 7, 6, 8, 5, 0, 9}
	pre   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	post  = []int{3, 4, 2, 7, 8, 6, 0, 9, 5, 1}
	level = [][]int{{1}, {2, 5}, {3, 4, 6, 9}, {7, 8, 0}}

	tree = &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 3},
			Right: &TreeNode{Value: 4},
		},
		Right: &TreeNode{
			Value: 5,
			Left: &TreeNode{
				Value: 6,
				Left:  &TreeNode{Value: 7},
				Right: &TreeNode{Value: 8},
			},
			Right: &TreeNode{
				Value: 9,
				Left:  &TreeNode{Value: 0},
			},
		},
	}
)

func TestPreOrder(t *testing.T) {
	assert := assert.New(t)

	order := tree.PreOrder()
	assert.Equal(pre, order, "PreOrder")

	order = tree.PreOrderRecursive()
	assert.Equal(pre, order, "PreOrderRecursive")
}

func TestPostOrder(t *testing.T) {
	order := tree.PostOrder()
	assert.Equal(t, post, order, "PostOrder")
}

func TestInOrder(t *testing.T) {
	order := tree.InOrder()
	assert.Equal(t, in, order, "InOrder")
}

func TestLevelOrder(t *testing.T) {
	order := tree.LevelOrder()
	assert.Equal(t, level, order, "LevelOrder")
}

func TestBuildTreeFromPostIn(t *testing.T) {
	builded := BuildTreeFromPostIn(post, in)
	assert.Equal(t, tree, builded, "BuildTreeFromPostIn")
}

func TestBuildTreeFromPreIn(t *testing.T) {
	builded := BuildTreeFromPreIn(pre, in)
	assert.Equal(t, tree, builded, "BuildTreeFromPreIn")
}
