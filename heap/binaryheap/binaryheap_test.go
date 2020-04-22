package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |


*/

type Int int

func (i Int) Less(j Item) bool {
	return i < j.(Int)
}

func TestHeap(t *testing.T) {
	var (
		values = []int{2, 8, 5, 1, 7, 9, 6, 3, 0, 4}
	)

	var items []Item
	for _, value := range values {
		items = append(items, Int(value))
	}

	heap := New(items)

	for index := range values {
		assert.Equal(t, index, int(heap.Pop().(Int)), "binaryheap.InitPop")
	}

	for _, value := range values {
		heap.Push(Int(value))
	}

	for index := range values {
		assert.Equal(t, index, int(heap.Pop().(Int)), "binaryheap.PushPop")
	}

}
