package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

	head -> 5 -> 3 -> 1 -> 4 -> 2 -> nil

*/

func singlyList() *List {
	singly := new(List)

	node0 := &ListNode{Value: 2, List: singly, Next: nil}
	node1 := &ListNode{Value: 4, List: singly, Next: node0}
	node2 := &ListNode{Value: 1, List: singly, Next: node1}
	node3 := &ListNode{Value: 3, List: singly, Next: node2}
	node4 := &ListNode{Value: 5, List: singly, Next: node3}

	singly.head = &ListNode{List: singly, Next: node4}
	singly.len = 5
	return singly
}

func TestInit(t *testing.T) {
	var (
		singly = singlyList()
		values = []interface{}{5, 3, 1, 4, 2}
		inited = New(values...)
	)
	assert.Equal(t, singly, inited)
	assert.Equal(t, 5, inited.Len())
}

func TestInsert(t *testing.T) {
	var (
		singly   = singlyList()
		inserted = singlyList()
	)

	inserted.Clear()
	inserted.Insert(0, nil)
	inserted.PushFront(1)
	inserted.PushFront(3)
	inserted.PushFront(5)
	inserted.PushBack(2)
	inserted.Insert(4, inserted.Find(1))
	assert.Equal(t, singly, inserted)
}

func TestRemove(t *testing.T) {
	var (
		removed = singlyList()
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
		singly = singlyList()
		values = []int{1, 2, 3, 4, 5}
	)

	for _, value := range values {
		assert.Equal(t, value, singly.Find(value).Value)
	}
}

func TestValues(t *testing.T) {
	var (
		singly = singlyList()
		values = []interface{}{5, 3, 1, 4, 2}
	)

	assert.Equal(t, values, singly.Values())
}
