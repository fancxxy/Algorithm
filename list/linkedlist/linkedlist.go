package linkedlist

// ListNode 单链表结点
type ListNode struct {
	Value interface{}
	next  *ListNode
}

// List 单链表
type List struct {
	head *ListNode
	tail *ListNode
	len  int
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
	list.head.next = nil
	list.tail = list.head
	list.len = 0
}

// First 返回首结点，空链表返回nil
func (list *List) First() *ListNode {
	return list.head.next
}

// Last 返回尾结点，空链表返回nil
func (list *List) Last() *ListNode {
	if list.len == 0 {
		return nil
	}
	return list.tail
}

// Find 查找第一个等于value的结点，没有返回nil
func (list *List) Find(value interface{}) *ListNode {
	node := list.head.next
	for node != nil && node.Value != value {
		node = node.next
	}

	return node
}

// Insert 把value插入到at结点之后，返回新插入的结点
func (list *List) Insert(value interface{}, at *ListNode) *ListNode {
	if at == nil {
		return nil
	}

	node := &ListNode{Value: value, next: at.next}
	at.next = node
	if at == list.tail {
		list.tail = node
	}
	list.len++
	return node
}

// PushFront 插入到首元素
func (list *List) PushFront(value interface{}) *ListNode {
	return list.Insert(value, list.head)
}

// PushBack 插入到尾元素
func (list *List) PushBack(value interface{}) *ListNode {
	if list.len == 0 {
		return list.Insert(value, list.head)
	}
	return list.Insert(value, list.Last())
}

// Remove 移除at结点，返回被移除结点的元素值，不存在返回nil
// 如果不是尾结点，可以把后继结点赋值给at，删除后继结点
func (list *List) Remove(at *ListNode) interface{} {
	if at == nil || at == list.head {
		return nil
	}

	node := list.head
	for node != nil && node.next != at {
		node = node.next
	}
	if node == nil {
		return nil
	}

	node.next = at.next
	list.len--
	if at == list.tail {
		list.tail = node
	}
	return at.Value
}

// PopFront 删除首节点
func (list *List) PopFront() interface{} {
	return list.Remove(list.First())
}

// PopBack 删除尾节点，返回该节点元素
func (list *List) PopBack() interface{} {
	return list.Remove(list.Last())
}

// Values 获取全部元素的值，返回slice
func (list *List) Values() []interface{} {
	node := list.head.next
	slice := make([]interface{}, list.len)
	for i := 0; i < len(slice); i++ {
		slice[i] = node.Value
		node = node.next
	}
	return slice
}

// New 创建单链表
func New(args ...interface{}) *List {
	list := new(List)
	list.head = new(ListNode)
	list.tail = list.head
	if len(args) != 0 {
		for _, arg := range args {
			list.PushBack(arg)
		}
	}
	return list
}
