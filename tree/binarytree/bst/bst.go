package bst

import "github.com/fancxxy/algorithm/tree/binarytree"

type (
	// TreeNode 二叉搜索树结点
	TreeNode = binarytree.TreeNode

	// Bst 二叉搜索树
	Bst struct {
		Root *TreeNode
	}
)

// Search 搜索
func (tree *Bst) Search(value int) *TreeNode {
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

// Insert 插入新结点
// 插入操作只会插在叶子结点后
func (tree *Bst) Insert(value int) {
	var (
		insert func(*TreeNode, int) *TreeNode
	)

	insert = func(node *TreeNode, value int) *TreeNode {
		if node == nil {
			return &TreeNode{Value: value}
		}

		if value > node.Value {
			node.Right = insert(node.Right, value)
		} else if value < node.Value {
			node.Left = insert(node.Left, value)
		}

		return node
	}

	tree.Root = insert(tree.Root, value)
}

// Delete 删除结点
// 被删除结点没有左结点，直接用右结点代替它
// 被删除结点没有右结点，用左结点代替它
// 被删除结点有左右结点，找到它的后继结点，把后继结点的值赋值给当前结点，删除右子树上的后继结点
func (tree *Bst) Delete(value int) {
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
		return node
	}

	tree.Root = delete(tree.Root, value)
}
