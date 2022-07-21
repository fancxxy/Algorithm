package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingBuffer(t *testing.T) {
	r := New(8, true)
	assert.Equal(t, 5, r.Write([]T{1, 2, 3, 4, 5}))
	assert.Equal(t, 5, r.Len())
	assert.Equal(t, 3, r.Write([]T{6, 7, 8}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, 3, r.Write([]T{'a', 'b', 'c'}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, []T{4, 5, 6}, r.Read(3))
	assert.Equal(t, 5, r.Len())
	assert.Equal(t, []T{7, 8, 'a', 'b', 'c'}, r.Read(10))
	assert.Equal(t, 0, r.Len())
	assert.Equal(t, []T{}, r.Read(1))
	assert.Equal(t, 5, r.Write([]T{1, 2, 3, 4, 5}))
	assert.Equal(t, 5, r.Len())
	assert.Equal(t, []T{1, 2}, r.Read(2))
	assert.Equal(t, 3, r.Len())
	assert.Equal(t, []T{}, r.Read(0))
	assert.Equal(t, 0, r.Write(nil))
	assert.Equal(t, 2, r.Write([]T{'d', 'e'}))
	assert.Equal(t, 5, r.Len())
	assert.Equal(t, 3, r.Write([]T{'a', 'b', 'c'}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, []T{3, 4, 5}, r.Read(3))
	assert.Equal(t, 5, r.Len())
	assert.Equal(t, 6, r.Write([]T{'a', 'b', 'c', 'd', 'e', 'f'}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, 4, r.Write([]T{1, 2, 3, 4}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, 5, r.Write([]T{5, 6, 7, 8, 9}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, []T{2, 3, 4, 5, 6}, r.Read(5))
	assert.Equal(t, 3, r.Len())
	assert.Equal(t, 3, r.Write([]T{1, 2, 3}))
	assert.Equal(t, 6, r.Len())
	assert.Equal(t, 8, r.Write([]T{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}))
	assert.Equal(t, 8, r.Len())
	assert.Equal(t, []T{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}, r.Read(10))
}
