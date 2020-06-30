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
	len  int
}

// New 创建单链表
func New(args ...interface{}) *List {
	list := new(List)
	list.head = &ListNode{List: list}
	node := list.head
	if len(args) != 0 {
		for _, arg := range args {
			node = list.Insert(arg, node)
		}
	}
	return list
}

// Len 返回链表长度
func (list *List) Len() int {
	return list.len
}

// Empty 判断链表是否为空
func (list *List) Empty() bool {
	return list.len == 0
}

// Clear 清空链表
func (list *List) Clear() {
	list.head.Next = nil
	list.len = 0
}

// First 返回首元素
func (list *List) First() *ListNode {
	return list.head.Next
}

// Last 返回尾元素
func (list *List) Last() *ListNode {
	node := list.head
	for i := 0; i < list.len; i++ {
		node = node.Next
	}
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

// Insert 把value插入到at结点之后，返回新插入的结点
func (list *List) Insert(value interface{}, at *ListNode) *ListNode {
	if at == nil || at.List != list {
		return nil
	}

	node := &ListNode{Value: value, List: list, Next: at.Next}
	at.Next = node
	list.len++
	return node
}

// PushFront 插入到首元素
func (list *List) PushFront(value interface{}) *ListNode {
	return list.Insert(value, list.head)
}

// PushBack 插入到尾元素
func (list *List) PushBack(value interface{}) *ListNode {
	return list.Insert(value, list.Last())
}

// Remove 移除at结点，成功返回true
// 如果at不是最后一个结点，把后继结点赋值给at，删除后继结点
func (list *List) Remove(at *ListNode) interface{} {
	if at == nil || at.List != list {
		return nil
	}

	// 如果at是尾节点
	if at.Next == nil {
		return list.PopBack()
	}

	ret := at.Value
	at.Value = at.Next.Value
	at.Next = at.Next.Next
	list.len--
	return ret
}

// PopFront 删除首节点
func (list *List) PopFront() interface{} {
	return list.Remove(list.First())
}

// PopBack 删除尾节点，返回该节点元素
func (list *List) PopBack() interface{} {
	if list.len == 0 {
		return nil
	}

	node := list.head
	for i := 0; i < list.len-1; i++ {
		node = node.Next
	}

	tail := node.Next
	node.Next = nil
	list.len--
	return tail.Value
}

// Values 获取全部元素的值，返回slice
func (list *List) Values() []interface{} {
	node := list.head.Next
	slice := make([]interface{}, list.len)
	for i := 0; i < len(slice); i++ {
		slice[i] = node.Value
		node = node.Next
	}
	return slice
}
