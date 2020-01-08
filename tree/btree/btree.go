package btree

/*

1.所有叶子结点在同一层
2.b树的最小度为t，度就是结点子树的个数
3.t - 1 <= len(elements) <= 2t - 1
4.t < len(children) < 2t
5.结点len(children) = len(elements) + 1
6.所有结点的key都是按递增顺序排序，关键字k1和k2之间的子树包含的所有关键字k1<key<k2
7.b树增长和收缩都是从root开始
8.查询，插入，删除时间复杂度都是O(logn)

Insert
所有插入操作都是在叶子结点
插入从根结点开始向下递归寻找合适的结点
遍历的所有结点，如果已满，分裂成两个结点
中间key上移到父结点合适位置，分裂后的结点分别获得原始结点前半和后半的key以及child

插入1，创建根结点
		| 1 |

插入2、3，无需分裂
		| 1 | 2 | 3 |

插入4，创建新的根结点，分裂原来的根结点作为叶子结点
		| 2 |
	   /     \
	| 1 |   | 3 | 4 |

插入5
 		| 2 |
	   /     \
   | 1 |    | 3 | 4 | 5 |

插入6，分裂叶子结点，上移元素4到父结点
		| 2 | 4 |
	  /     |     \
   | 1 |  | 3 |  | 5 | 6 |

插入7
		| 2 | 4 |
	   /    |     \
   | 1 |  | 3 |  | 5 | 6 | 7 |

插入8，分裂上移元素6
		| 2 | 4 | 6 |
	  /     |   |     \
  | 1 |  | 3 | | 5 |  | 7 | 8 |

插入9，分裂根结点，继续遍历，插入元素到叶子结点
			   | 4 |
		    /	      \
	   | 2 |         | 6 |
	 /      \       /     \
 | 1 |    | 3 |   | 5 |   | 7 | 8 | 9 |


*/

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
	for ; i >= 0 && value < slice[i]; i-- {
	}
	return i + 1
}
