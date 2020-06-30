package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

	head <-> 5 <-> 3 <-> 1 <-> 4 <-> 2 <-> head

*/

func doublyList() *List {
	doubly := new(List)

	node0 := &ListNode{Value: 2, List: doubly}
	node1 := &ListNode{Value: 4, List: doubly, Next: node0}
	node2 := &ListNode{Value: 1, List: doubly, Next: node1}
	node3 := &ListNode{Value: 3, List: doubly, Next: node2}
	node4 := &ListNode{Value: 5, List: doubly, Next: node3}
	doubly.head = &ListNode{List: doubly, Next: node4}
	node0.Next = doubly.head
	doubly.head.Prev = node0
	node0.Prev = node1
	node1.Prev = node2
	node2.Prev = node3
	node3.Prev = node4
	node4.Prev = doubly.head

	doubly.len = 5
	return doubly
}

func TestInit(t *testing.T) {
	var (
		doubly = doublyList()
		values = []interface{}{5, 3, 1, 4, 2}
		inited = New(values...)
	)
	assert.Equal(t, doubly, inited)
	assert.Equal(t, 5, inited.Len())
}

func TestInsert(t *testing.T) {
	var (
		doubly   = doublyList()
		inserted = doublyList()
	)

	inserted.Clear()
	inserted.InsertAfter(0, nil)
	inserted.InsertBefore(0, nil)
	inserted.PushFront(1)
	inserted.PushFront(3)
	inserted.PushFront(5)
	inserted.PushBack(4)
	inserted.PushBack(2)
	assert.Equal(t, doubly, inserted)
	assert.Equal(t, 5, inserted.Len())
	assert.Equal(t, 5, inserted.First().Value)
	assert.Equal(t, 2, inserted.Last().Value)
}

func TestRemove(t *testing.T) {
	var (
		removed = doublyList()
	)

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
		doubly = doublyList()
		values = []int{1, 2, 3, 4, 5}
	)

	for _, value := range values {
		assert.Equal(t, value, doubly.Find(value).Value)
		assert.Equal(t, value, doubly.FindPrev(value).Value)
	}
	assert.Nil(t, doubly.Find(0))
	assert.Nil(t, doubly.FindPrev(-1))
}

func TestValues(t *testing.T) {
	var (
		doubly = doublyList()
		values = []interface{}{5, 3, 1, 4, 2}
	)

	assert.Equal(t, values, doubly.Values())
}
