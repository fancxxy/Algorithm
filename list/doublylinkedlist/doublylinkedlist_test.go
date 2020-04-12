package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

	head <-> 5 <-> 3 <-> 1 <-> 4 <-> 2 <-> nil

*/

func doublyList() *List {
	doubly := new(List)

	node0 := &ListNode{Value: 2, List: doubly, Next: nil}
	node1 := &ListNode{Value: 4, List: doubly, Next: node0}
	node2 := &ListNode{Value: 1, List: doubly, Next: node1}
	node3 := &ListNode{Value: 3, List: doubly, Next: node2}
	node4 := &ListNode{Value: 5, List: doubly, Next: node3}
	doubly.head = &ListNode{List: doubly, Next: node4}

	node4.Prev = doubly.head
	node3.Prev = node4
	node2.Prev = node3
	node1.Prev = node2
	node0.Prev = node1

	doubly.size = 5
	return doubly
}

func TestInsert(t *testing.T) {
	var (
		doubly             = doublyList()
		inserted           = New()
		values             = []int{3, 1, 4, 2}
		node     *ListNode = nil
	)

	assert.Equal(t, true, inserted.Empty(), "doublylinkedlist.Empty")
	assert.Equal(t, node, inserted.Insert(values[0], node), "doublylinkedlist.Insert")
	node = inserted.Head()

	for i, value := range values {
		assert.Equal(t, i, inserted.Size(), "doublylinkedlist.Size")
		node = inserted.Insert(value, node)
	}

	assert.NotEqual(t, (*ListNode)(nil), inserted.Insert(5, inserted.Head()), "doublylinkedlist.Insert")
	assert.Equal(t, false, inserted.Empty(), "doublylinkedlist.Empty")
	assert.Equal(t, doubly, inserted, "doublylinkedlist.Insert")
	assert.Equal(t, 5, inserted.First().Value, "doublylinkedlist.First")
}

func TestDelete(t *testing.T) {
	var (
		deleted = doublyList()
		values  = []int{1, 2, 3, 4, 5}
		cleared = doublyList()
	)
	for i, value := range values {
		assert.Equal(t, len(values)-i, deleted.Size(), "doublylinkedlist.Size")
		assert.Equal(t, true, deleted.Delete(value), "doublylinkedlist.Delete")
	}

	assert.Equal(t, false, deleted.Delete(values[0]), "doublylinkedlist.Delete")
	assert.Equal(t, (*ListNode)(nil), deleted.First(), "doublylinkedlist.First")

	cleared.Clear()
	assert.Equal(t, deleted, cleared, "doublylinkedlist.Clear")
}

func TestFind(t *testing.T) {
	var (
		doubly = doublyList()
		values = []int{1, 2, 3, 4, 5}
	)

	for _, value := range values {
		node := doubly.Find(value)
		assert.NotEqual(t, nil, node, "doublylinkedlist.Find")
		assert.Equal(t, value, node.Value, "doublylinkedlist.Find")
	}
}

func TestValues(t *testing.T) {
	var (
		doubly = doublyList()
		values = []interface{}{5, 3, 1, 4, 2}
	)

	assert.Equal(t, values, doubly.Values(), "doublylinked.Values")
}
