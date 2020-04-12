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
	singly.size = 5
	return singly
}

func TestInsert(t *testing.T) {
	var (
		singly             = singlyList()
		inserted           = New()
		values             = []int{5, 3, 1, 4, 2}
		node     *ListNode = nil
	)

	assert.Equal(t, true, inserted.Empty(), "singlylinkedlist.Empty")
	assert.Equal(t, node, inserted.Insert(values[0], node), "singlylinkedlist.Insert")
	node = inserted.Head()

	for i, value := range values {
		assert.Equal(t, i, inserted.Size(), "singlylinkedlist.Size")
		node = inserted.Insert(value, node)
	}

	assert.Equal(t, false, inserted.Empty(), "singlylinkedlist.Empty")
	assert.Equal(t, singly, inserted, "singlylinkedlist.Insert")
	assert.Equal(t, 5, inserted.First().Value, "singlylinkedlist.First")
}

func TestDelete(t *testing.T) {
	var (
		deleted = singlyList()
		values  = []int{1, 2, 3, 4, 5}
		cleared = singlyList()
	)
	for i, value := range values {
		assert.Equal(t, len(values)-i, deleted.Size(), "singlylinkedlist.Size")
		assert.Equal(t, true, deleted.Delete(value), "singlylinkedlist.Delete")
	}

	assert.Equal(t, false, deleted.Delete(values[0]), "singlylinkedlist.Delete")
	assert.Equal(t, (*ListNode)(nil), deleted.First(), "singlylinkedlist.First")

	cleared.Clear()
	assert.Equal(t, deleted, cleared, "singlylinkedlist.Clear")
}

func TestFind(t *testing.T) {
	var (
		singly = singlyList()
		values = []int{1, 2, 3, 4, 5}
	)

	for _, value := range values {
		node := singly.Find(value)
		assert.NotEqual(t, nil, node, "singlylinkedlist.Find")
		assert.Equal(t, value, node.Value, "singlylinkedlist.Find")
	}
}

func TestValues(t *testing.T) {
	var (
		singly = singlyList()
		values = []interface{}{5, 3, 1, 4, 2}
	)

	assert.Equal(t, values, singly.Values(), "singlylinked.Values")
}
