package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func linkedList() *List {
	head := &ListNode{}
	node0 := &ListNode{Value: 5}
	node1 := &ListNode{Value: 3}
	node2 := &ListNode{Value: 1}
	node3 := &ListNode{Value: 4}
	tail := &ListNode{Value: 2}

	head.next = node0
	node0.next = node1
	node1.next = node2
	node2.next = node3
	node3.next = tail

	return &List{head: head, tail: tail, len: 5}
}

func TestInit(t *testing.T) {
	var (
		list   = linkedList()
		values = []interface{}{5, 3, 1, 4, 2}
		inited = New(values...)
	)
	assert.Equal(t, list, inited)
	assert.Equal(t, 5, inited.Len())
}

func TestInsert(t *testing.T) {
	var (
		list     = linkedList()
		inserted = linkedList()
	)

	inserted.Clear()
	inserted.Insert(0, nil)
	inserted.PushFront(1)
	inserted.PushFront(3)
	inserted.PushFront(5)
	inserted.PushBack(2)
	inserted.Insert(4, inserted.Find(1))
	assert.Equal(t, list, inserted)
}

func TestRemove(t *testing.T) {
	var (
		removed = linkedList()
	)

	assert.Nil(t, removed.Remove(new(ListNode)))
	assert.Equal(t, 5, removed.PopFront())
	assert.Equal(t, 2, removed.PopBack())
	assert.Equal(t, 4, removed.Remove(removed.Find(4)))
	assert.Equal(t, 3, removed.PopFront())
	assert.Equal(t, 1, removed.PopFront())
	assert.Equal(t, 0, removed.Len())
	assert.Nil(t, removed.PopFront())
	assert.Nil(t, removed.PopBack())
	assert.True(t, removed.Empty())
}

func TestFind(t *testing.T) {
	var (
		list   = linkedList()
		values = []int{1, 2, 3, 4, 5}
	)

	for _, value := range values {
		assert.Equal(t, value, list.Find(value).Value)
	}
	assert.Nil(t, list.Find(0))
}

func TestValues(t *testing.T) {
	var (
		list   = linkedList()
		values = []interface{}{5, 3, 1, 4, 2}
	)

	assert.Equal(t, values, list.Values())
}
