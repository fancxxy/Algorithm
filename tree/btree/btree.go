package btree

type (
	// BTree b树
	BTree struct {
		Degree int
		Root   *TreeNode
	}

	// TreeNode b树的结点
	TreeNode struct {
		Elements []int
		Children []*TreeNode
		Leaf     bool
	}
)

// Search 查找
func (tree *BTree) Search(value int) *TreeNode {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.search(value)
}

func (node *TreeNode) search(value int) *TreeNode {
	index := find(node.Elements, value)
	if index < len(node.Elements) && node.Elements[index] == value {
		return node
	}
	if node.Leaf {
		return nil
	}
	return node.Children[index].search(value)
}

// Insert 插入
func (tree *BTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &TreeNode{
			Elements: make([]int, 0, 2*tree.Degree-1),
			Children: make([]*TreeNode, 0, 2*tree.Degree),
			Leaf:     true,
		}
		tree.Root.Elements = append(tree.Root.Elements, value)
		return
	}

	if len(tree.Root.Elements) == cap(tree.Root.Elements) {
		node := &TreeNode{
			Elements: make([]int, 0, 2*tree.Degree-1),
			Children: make([]*TreeNode, 0, 2*tree.Degree),
		}
		element, splited := tree.Root.split()
		node.Children = append(node.Children, tree.Root, splited)
		node.Elements = append(node.Elements, element)
		tree.Root = node
	}

	tree.Root.insert(value)
}

func (node *TreeNode) insert(value int) {
	index := find(node.Elements, value)
	if node.Leaf {
		node.Elements = insertElement(node.Elements, index, value)
		return
	}

	child := node.Children[index]
	if len(child.Elements) == cap(child.Elements) {
		element, splited := child.split()
		node.Elements = insertElement(node.Elements, index, element)
		node.Children = insertChild(node.Children, index+1, splited)
		if node.Elements[index] < value {
			child = node.Children[index+1]
		}
	}
	child.insert(value)
}

func (node *TreeNode) split() (int, *TreeNode) {
	degree := cap(node.Children) / 2

	element := node.Elements[degree-1]
	splited := &TreeNode{
		Elements: make([]int, degree-1, 2*degree-1),
		Children: make([]*TreeNode, 0, 2*degree),
		Leaf:     node.Leaf,
	}
	copy(splited.Elements[0:], node.Elements[degree:])
	node.Elements = node.Elements[:degree-1]
	if !node.Leaf {
		splited.Children = make([]*TreeNode, degree, 2*degree)
		copy(splited.Children[0:], node.Children[degree:])
		node.Children = node.Children[:degree]
	}
	return element, splited
}

func insertElement(slice []int, index int, value int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func insertChild(slice []*TreeNode, index int, value *TreeNode) []*TreeNode {
	slice = append(slice, nil)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func find(slice []int, value int) int {
	i := len(slice) - 1
	for i >= 0 && value <= slice[i] {
		i--
	}
	return i + 1
}

// Remove 删除
func (tree *BTree) Remove(value int) {
	if tree.Root == nil {
		return
	}

	tree.Root.remove(value)

	if len(tree.Root.Elements) == 0 {
		if tree.Root.Leaf {
			tree.Root = nil
		} else {
			tree.Root = tree.Root.Children[0]
		}
	}
}

func (node *TreeNode) remove(value int) {
	index := find(node.Elements, value)
	degree := cap(node.Children) / 2
	// 找到被删除元素
	if index < len(node.Elements) && node.Elements[index] == value {
		// 在叶子结点，直接删除
		if node.Leaf {
			node.Elements = append(node.Elements[:index], node.Elements[index+1:]...)
			return
		}

		// 在非叶子结点，找前驱、后继结点，条件不足合并俩子结点
		if len(node.Children[index].Elements) >= degree {
			pred := node.predecessor(index)
			node.Elements[index] = pred
			node.Children[index].remove(pred)
		} else if len(node.Children[index+1].Elements) >= degree {
			succ := node.successor(index)
			node.Elements[index] = succ
			node.Children[index].remove(succ)
		} else {
			node.merge(index)
			node.Children[index].remove(value)
		}
		return
	}

	// 没找到，已到叶子结点直接返回
	if node.Leaf {
		return
	}

	// 查询的中间结点，子结点个数如果等于最小度t-1，需要从别处补一个值
	if len(node.Children[index].Elements) < degree {
		if index != 0 && len(node.Children[index-1].Elements) >= degree {
			node.borrowFromPrev(index)
		} else if index != len(node.Elements) && len(node.Children[index+1].Elements) >= degree {
			node.borrowFromNext(index)
		} else {
			if index == len(node.Elements) {
				index--
			}
			node.merge(index)
		}
	}

	node.Children[index].remove(value)
}

func (node *TreeNode) borrowFromPrev(index int) {
	child, sibling := node.Children[index], node.Children[index-1]
	n := len(sibling.Elements)

	child.Elements = insertElement(child.Elements, 0, node.Elements[index-1])
	node.Elements[index-1] = sibling.Elements[n-1]
	sibling.Elements = sibling.Elements[:n-1]

	if !child.Leaf {
		child.Children = insertChild(child.Children, 0, sibling.Children[n])
		sibling.Children = sibling.Children[:n]
	}
}

func (node *TreeNode) borrowFromNext(index int) {
	child, sibling := node.Children[index], node.Children[index+1]

	child.Elements = append(child.Elements, node.Elements[index])
	node.Elements[index] = sibling.Elements[0]
	sibling.Elements = sibling.Elements[1:]

	if !child.Leaf {
		child.Children = append(child.Children, sibling.Children[0])
		sibling.Children = sibling.Children[1:]
	}
}

func (node *TreeNode) merge(index int) {
	child, sibling := node.Children[index], node.Children[index+1]

	child.Elements = append(child.Elements, node.Elements[index])
	child.Elements = append(child.Elements, sibling.Elements...)
	node.Elements = append(node.Elements[:index], node.Elements[index+1:]...)

	if !child.Leaf {
		child.Children = append(child.Children, sibling.Children...)
	}
	node.Children = append(node.Children[:index+1], node.Children[index+2:]...)
}

func (node *TreeNode) predecessor(index int) int {
	curr := node.Children[index]

	for !curr.Leaf {
		curr = curr.Children[len(curr.Elements)]
	}
	return curr.Elements[len(curr.Elements)-1]
}

func (node *TreeNode) successor(index int) int {
	curr := node.Children[index+1]

	for !curr.Leaf {
		curr = curr.Children[0]
	}
	return curr.Elements[0]
}
