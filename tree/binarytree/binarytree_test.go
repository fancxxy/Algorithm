package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
                  1
                 / \
                /   \
               2     \
              / \     5
             3   4   / \
                    /   \
                   6     \
                  / \     9
				 7   8   /
				        0
*/

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
	assert.Equal(pre, order, "TreeNode.PreOrder")

	order = tree.PreOrderRecursive()
	assert.Equal(pre, order, "TreeNode.PreOrderRecursive")
}

func TestPostOrder(t *testing.T) {
	order := tree.PostOrder()
	assert.Equal(t, post, order, "TreeNode.PostOrder")
}

func TestInOrder(t *testing.T) {
	order := tree.InOrder()
	assert.Equal(t, in, order, "TreeNode.InOrder")
}

func TestLevelOrder(t *testing.T) {
	order := tree.LevelOrder()
	assert.Equal(t, level, order, "TreeNode.LevelOrder")

	order = tree.LevelOrderRecursive()
	assert.Equal(t, level, order, "TreeNode.LevelOrder")
}

func TestBuildTreeFromPostIn(t *testing.T) {
	builded := BuildTreeFromPostIn(post, in)
	assert.Equal(t, tree, builded, "BuildTreeFromPostIn")
}

func TestBuildTreeFromPreIn(t *testing.T) {
	builded := BuildTreeFromPreIn(pre, in)
	assert.Equal(t, tree, builded, "BuildTreeFromPreIn")
}

func TestInvertTree(t *testing.T) {
	inverted := &TreeNode{
		Value: 1,
		Right: &TreeNode{
			Value: 2,
			Right: &TreeNode{Value: 3},
			Left:  &TreeNode{Value: 4},
		},
		Left: &TreeNode{
			Value: 5,
			Right: &TreeNode{
				Value: 6,
				Right: &TreeNode{Value: 7},
				Left:  &TreeNode{Value: 8},
			},
			Left: &TreeNode{
				Value: 9,
				Right: &TreeNode{Value: 0},
			},
		},
	}

	assert.Equal(t, inverted, tree.InvertTree(), "TreeNode.InvertTree")
	assert.Equal(t, inverted, tree.InvertTree().InvertTreeRecursive(), "TreeNode.InvertTreeRecursive")
}

func TestPredecessor(t *testing.T) {
	assert.Equal(t, tree.Predecessor().Value, 4, "TreeNode.Predecessor")
	assert.NotEqual(t, tree.Predecessor().Value, 3, "TreeNode.Predecessor")
}

func TestSuccessor(t *testing.T) {
	assert.Equal(t, tree.Successor().Value, 7, "TreeNode.Successor")
	assert.NotEqual(t, tree.Successor().Value, 8, "TreeNode.Successor")
}
