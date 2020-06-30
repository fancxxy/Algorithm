package doublylinkedlist

// ListNode 双链表结点
type ListNode struct {
	Value interface{}
	list  *List
	next  *ListNode
	prev  *ListNode
}

// List 双链表
type List struct {
	head *ListNode
	len  int
}

// New 创建双链表
func New(args ...interface{}) *List {
	list := new(List)
	list.head = &ListNode{list: list}
	list.head.next = list.head
	list.head.prev = list.head
	if len(args) != 0 {
		for _, arg := range args {
			list.PushBack(arg)
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
	list.head.next = list.head
	list.head.prev = list.head
	list.len = 0
}

// First 返回首元素
func (list *List) First() *ListNode {
	if list.len == 0 {
		return nil
	}
	return list.head.next
}

// Last 返回尾元素
func (list *List) Last() *ListNode {
	if list.len == 0 {
		return nil
	}
	return list.head.prev
}

// Find 返回第一个相等的元素的地址，没有找到返回nil
func (list *List) Find(value interface{}) *ListNode {
	node := list.head.next
	for node != list.head && node.Value != value {
		node = node.next
	}

	if node == list.head {
		return nil
	}
	return node
}

// FindPrev 返回最后一个相等的元素的地址，没有找到返回nil
func (list *List) FindPrev(value interface{}) *ListNode {
	node := list.head.prev
	for node != list.head && node.Value != value {
		node = node.prev
	}

	if node == list.head {
		return nil
	}
	return node
}

// PushFront 插入到首元素
func (list *List) PushFront(value interface{}) *ListNode {
	return list.InsertAfter(value, list.head)
}

// PushBack 插入到尾元素
func (list *List) PushBack(value interface{}) *ListNode {
	return list.InsertBefore(value, list.head)
}

// InsertAfter 把value插入到at结点之后，返回新插入的结点
func (list *List) InsertAfter(value interface{}, at *ListNode) *ListNode {
	if at == nil || at.list != list {
		return nil
	}

	node := &ListNode{Value: value, list: list, next: at.next, prev: at}
	at.next.prev = node
	at.next = node
	list.len++
	return node
}

// InsertBefore 把value插入到at结点之前，返回新插入的结点
func (list *List) InsertBefore(value interface{}, at *ListNode) *ListNode {
	if at == nil || at.list != list {
		return nil
	}

	node := &ListNode{Value: value, list: list, next: at, prev: at.prev}
	at.prev.next = node
	at.prev = node
	list.len++
	return node
}

// PopFront 移除首元素
func (list *List) PopFront() interface{} {
	return list.Remove(list.head.next)
}

// PopBack 移除尾元素
func (list *List) PopBack() interface{} {
	return list.Remove(list.head.prev)
}

// Remove 移除at结点，返回节点元素
func (list *List) Remove(at *ListNode) interface{} {
	if at == nil || at == list.head || at.list != list {
		return nil
	}

	next, prev := at.next, at.prev
	next.prev = prev
	prev.next = next
	list.len--
	return at.Value
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
