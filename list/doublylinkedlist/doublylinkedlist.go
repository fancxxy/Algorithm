package doublylinkedlist

// ListNode 双链表结点
type ListNode struct {
	Value interface{}
	List  *List
	Next  *ListNode
	Prev  *ListNode
}

// List 双链表
type List struct {
	head *ListNode
	tail *ListNode
	len  int
}

// New 创建双链表
func New(args ...interface{}) *List {
	list := new(List)
	list.head = &ListNode{List: list}
	list.head.Next = list.head
	list.head.Prev = list.head
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
	list.head.Next = list.head
	list.head.Prev = list.head
	list.len = 0
}

// First 返回首元素
func (list *List) First() *ListNode {
	return list.head.Next
}

// Last 返回尾元素
func (list *List) Last() *ListNode {
	return list.head.Prev
}

// Find 根据值查找元素，返回找到的第一个元素的地址，没有找到返回nil
func (list *List) Find(value interface{}) *ListNode {
	node := list.head.Next
	for node != list.head && node.Value != value {
		node = node.Next
	}

	if node == list.head {
		return nil
	}
	return node
}

// FindPrev 返回最后一个相等的元素的地址，没有找到返回nil
func (list *List) FindPrev(value interface{}) *ListNode {
	node := list.head.Prev
	for node != list.head && node.Value != value {
		node = node.Prev
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
	if at == nil || at.List != list {
		return nil
	}

	node := &ListNode{Value: value, List: list, Next: at.Next, Prev: at}
	at.Next.Prev = node
	at.Next = node
	list.len++
	return node
}

// InsertBefore 把value插入到at结点之前，返回新插入的结点
func (list *List) InsertBefore(value interface{}, at *ListNode) *ListNode {
	if at == nil || at.List != list {
		return nil
	}

	node := &ListNode{Value: value, List: list, Next: at, Prev: at.Prev}
	at.Prev.Next = node
	at.Prev = node
	list.len++
	return node
}

// PopFront 移除首元素
func (list *List) PopFront() interface{} {
	return list.Remove(list.head.Next)
}

// PopBack 移除尾元素
func (list *List) PopBack() interface{} {
	return list.Remove(list.head.Prev)
}

// Remove 移除at结点，返回节点元素
func (list *List) Remove(at *ListNode) interface{} {
	if at == nil || at == list.head || at.List != list {
		return nil
	}

	next, prev := at.Next, at.Prev
	next.Prev = prev
	prev.Next = next
	list.len--
	return at.Value
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
