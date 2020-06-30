package singlylinkedlist

// ListNode 单链表结点
type ListNode struct {
	Value interface{}
	list  *List
	next  *ListNode
}

// List 链表
type List struct {
	head *ListNode
	len  int
}

// New 创建单链表
func New(args ...interface{}) *List {
	list := new(List)
	list.head = &ListNode{list: list}
	list.head.next = list.head
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

	node := list.head
	for node.next != list.head {
		node = node.next
	}
	return node
}

// Find 返回第一个相等的元素的地址，没有找到返回nil
func (list *List) Find(value interface{}) *ListNode {
	if list.len == 0 {
		return nil
	}

	node := list.head.next
	for node != list.head && node.Value != value {
		node = node.next
	}

	if node == list.head {
		return nil
	}
	return node
}

// Insert 把value插入到at结点之后，返回新插入的结点
func (list *List) Insert(value interface{}, at *ListNode) *ListNode {
	if at == nil || at.list != list {
		return nil
	}

	node := &ListNode{Value: value, list: list, next: at.next}
	at.next = node
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

// Remove 移除at结点，成功返回true
// 把后继结点赋值给at，删除后继结点
func (list *List) Remove(at *ListNode) interface{} {
	if at == nil || at == list.head || at.list != list {
		return nil
	}

	ret := at.Value
	at.Value = at.next.Value
	at.next = at.next.next
	list.len--
	// 如果at是尾节点
	if at.next == list.head {
		list.head = at
	}
	return ret
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
