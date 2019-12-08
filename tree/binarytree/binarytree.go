package binarytree

// TreeNode 二叉树结点
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 前序遍历 DFS
func (root *TreeNode) PreOrder() []int {
	var (
		stack []*TreeNode
		order []int
	)

	if root == nil {
		return order
	}

	for stack = append(stack, root); len(stack) != 0; {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, curr.Value)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	return order
}

// PreOrderRecursive 前序遍历递归
func (root *TreeNode) PreOrderRecursive() []int {
	var (
		order     []int
		traversal func(*TreeNode)
	)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		order = append(order, node.Value)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)

	return order
}

// InOrder 中序遍历
func (root *TreeNode) InOrder() []int {
	var (
		stack []*TreeNode
		order []int
	)

	for curr := root; curr != nil || len(stack) != 0; {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, curr.Value)
		curr = curr.Right
	}

	return order
}

// PostOrder 后序遍历
// 前序遍历镜像 根->右->左，结果逆序 左->右->根
func (root *TreeNode) PostOrder() []int {
	var (
		stack []*TreeNode
		order []int
	)

	if root == nil {
		return order
	}

	for stack = append(stack, root); len(stack) != 0; {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, curr.Value)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	for left, right := 0, len(order)-1; left < right; left, right = left+1, right-1 {
		order[left], order[right] = order[right], order[left]
	}

	return order
}

// LevelOrder 层次遍历 BFS
func (root *TreeNode) LevelOrder() [][]int {
	var (
		order [][]int
		queue []*TreeNode
		level int
	)

	if root == nil {
		return order
	}

	for queue = append(queue, root); len(queue) != 0; {
		order = append(order, []int{})
		length := len(queue)
		for i := 0; i < length; i++ {
			order[level] = append(order[level], queue[i].Value)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
		level++
	}

	return order
}

// LevelOrderRecursive 层次遍历
func (root *TreeNode) LevelOrderRecursive() [][]int {
	var (
		traversal func(*TreeNode, int)
		order     [][]int
	)

	traversal = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(order) == level {
			order = append(order, []int{})
		}
		order[level] = append(order[level], node.Value)
		traversal(node.Left, level+1)
		traversal(node.Right, level+1)
	}
	traversal(root, 0)

	return order
}

// BuildTreeFromPreIn 从前序和中序构造二叉树
func BuildTreeFromPreIn(preorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{Value: preorder[0]}
	if len(inorder) == 1 {
		return node
	}

	idx := index(node.Value, inorder)
	node.Left = BuildTreeFromPreIn(preorder[1:idx+1], inorder[:idx])
	node.Right = BuildTreeFromPreIn(preorder[idx+1:], inorder[idx+1:])

	return node
}

// BuildTreeFromPostIn 从后序和中序构造二叉树
func BuildTreeFromPostIn(postorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{Value: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return node
	}

	idx := index(node.Value, inorder)
	node.Left = BuildTreeFromPostIn(postorder[:idx], inorder[:idx])
	node.Right = BuildTreeFromPostIn(postorder[idx:len(postorder)-1], inorder[idx+1:])

	return node
}

// InvertTree 翻转二叉树
func (root *TreeNode) InvertTree() *TreeNode {
	var queue []*TreeNode
	if root == nil {
		return root
	}

	for queue = append(queue, root); len(queue) != 0; {
		curr := queue[0]
		curr.Left, curr.Right = curr.Right, curr.Left
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
		queue = queue[1:]
	}

	return root
}

// InvertTreeRecursive 翻转二叉树
func (root *TreeNode) InvertTreeRecursive() *TreeNode {
	var invert func(*TreeNode) *TreeNode

	invert = func(node *TreeNode) *TreeNode {
		if node == nil {
			return node
		}

		node.Left, node.Right = node.Right, node.Left
		node.Left = invert(node.Left)
		node.Right = invert(node.Right)

		return node
	}

	return invert(root)
}

func index(val int, slice []int) int {
	for i, v := range slice {
		if val == v {
			return i
		}
	}

	return 0
}
