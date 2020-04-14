package singlylinkedlist

// ListNode 单链表结点
type ListNode struct {
	Value interface{}
	List  *List
	Next  *ListNode
}

// List 链表
type List struct {
	head *ListNode
	size int
}

// New 创建单链表
func New() *List {
	list := new(List)
	list.head = &ListNode{List: list}
	return list
}

// Head 返回头结点
func (list *List) Head() *ListNode {
	return list.head
}

// Size 返回链表长度
func (list *List) Size() int {
	return list.size
}

// First 返回第一个元素，可能为nil
func (list *List) First() *ListNode {
	return list.head.Next
}

// Empty 判断链表是否为空
func (list *List) Empty() bool {
	return list.size == 0
}

// Clear 删除链表所有结点
func (list *List) Clear() {
	list.head.Next = nil
	list.size = 0
}

// Insert 把value插入到at结点之后，返回新插入的结点
func (list *List) Insert(value interface{}, at *ListNode) *ListNode {
	if at == nil || at.List != list {
		return nil
	}

	node := &ListNode{Value: value, List: list, Next: at.Next}
	at.Next = node
	list.size++
	return node
}

// Find 根据值查找元素，返回找到的第一个元素的地址，没有找到返回nil
func (list *List) Find(value interface{}) *ListNode {
	node := list.head.Next
	for node != nil && node.Value != value {
		node = node.Next
	}
	return node
}

// Delete 删除第一个元素值为value结点，成功返回true
func (list *List) Delete(value interface{}) bool {
	node := list.head
	for node.Next != nil && node.Next.Value != value {
		node = node.Next
	}

	if node.Next == nil {
		return false
	}

	return list.Remove(node)
}

// Remove 移除at结点，成功返回true
// 如果at不是最后一个结点，把后继结点赋值给at，删除后继结点
func (list *List) Remove(at *ListNode) bool {
	if at == nil || at.List != list {
		return false
	}

	// 找到at的前驱结点
	if at.Next == nil {
		node := list.head
		for node.Next != nil && node.Next != at {
			node = node.Next
		}
		// 没有找到
		if node.Next == nil {
			return false
		}
		node.Next = nil
	} else {
		at.Value = at.Next.Value
		at.Next = at.Next.Next
	}

	list.size--
	return true
}

// Values 获取全部元素的值，返回slice
func (list *List) Values() []interface{} {
	node := list.head.Next
	slice := make([]interface{}, list.size)
	for i := 0; i < len(slice); i++ {
		slice[i] = node.Value
		node = node.Next
	}
	return slice
}
