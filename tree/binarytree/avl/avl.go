package avl

import (
	"github.com/fancxxy/algorithm/tree/binarytree"
)

type (
	// TreeNode 平衡二叉树结点
	TreeNode = binarytree.TreeNode

	// Avl 平衡二叉树
	Avl struct {
		Root *TreeNode
	}
)

// Search 查找
func (tree *Avl) Search(value int) *TreeNode {
	curr := tree.Root
	for curr != nil {
		if value > curr.Value {
			curr = curr.Right
		} else if value < curr.Value {
			curr = curr.Left
		} else {
			return curr
		}
	}
	return nil
}

// Insert 插入
func (tree *Avl) Insert(value int) {
	var (
		insert func(*TreeNode, int) *TreeNode
	)
	insert = func(node *TreeNode, value int) *TreeNode {
		// 按照二叉搜索树的方式插入结点
		if node == nil {
			return &TreeNode{Value: value}
		}
		if value > node.Value {
			node.Right = insert(node.Right, value)
		} else if value < node.Value {
			node.Left = insert(node.Left, value)
		} else {
			return node
		}

		// 更新结点高度
		node.Height = max(height(node.Left), height(node.Right)) + 1

		// 获取平衡因子并旋转子树
		bf := balanceFactor(node)
		// LL
		if bf > 1 && value < node.Left.Value {
			return rightRotation(node)
		}
		// RR
		if bf < -1 && value > node.Right.Value {
			return leftRotation(node)
		}
		// LR
		if bf > 1 && value > node.Left.Value {
			node.Left = leftRotation(node.Left)
			return rightRotation(node)
		}
		// RL
		if bf < -1 && value < node.Right.Value {
			node.Right = rightRotation(node.Right)
			return leftRotation(node)
		}
		return node
	}

	tree.Root = insert(tree.Root, value)
}

// Remove 删除
func (tree *Avl) Remove(value int) {
	var (
		delete func(*TreeNode, int) *TreeNode
	)
	delete = func(node *TreeNode, value int) *TreeNode {
		if node == nil {
			return node
		}
		if value > node.Value {
			node.Right = delete(node.Right, value)
		} else if value < node.Value {
			node.Left = delete(node.Left, value)
		} else {
			if node.Left == nil {
				node = node.Right
			} else if node.Right == nil {
				node = node.Left
			} else {
				node.Value = node.Successor().Value
				node.Right = delete(node.Right, node.Value)
			}
		}

		if node == nil {
			return node
		}

		node.Height = max(height(node.Left), height(node.Right)) + 1

		bf := balanceFactor(node)
		// LL
		if bf > 1 && balanceFactor(node.Left) >= 0 {
			return rightRotation(node)
		}
		// RR
		if bf < -1 && balanceFactor(node.Right) <= 0 {
			return leftRotation(node)
		}
		// LR
		if bf > 1 && balanceFactor(node.Left) < 0 {
			node.Left = leftRotation(node.Left)
			return rightRotation(node)
		}
		// RL
		if bf < -1 && balanceFactor(node.Right) < 0 {
			node.Right = rightRotation(node.Right)
			return leftRotation(node)
		}
		return node
	}

	tree.Root = delete(tree.Root, value)
}

func rightRotation(x *TreeNode) *TreeNode {
	y := x.Left
	x.Left = y.Right
	y.Right = x

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

func leftRotation(x *TreeNode) *TreeNode {
	y := x.Right
	x.Right = y.Left
	y.Left = x

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

// 获取平衡因子
func balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
