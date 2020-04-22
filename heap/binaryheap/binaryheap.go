package binaryheap

// Item 二叉堆的数据需要实现的接口
type Item interface {
	Less(Item) bool
}

// BinaryHeap 二叉堆
type BinaryHeap struct {
	values []Item
}

// New 构造二叉堆
func New(values []Item) *BinaryHeap {
	heap := &BinaryHeap{
		values: values,
	}

	heap.Init()
	return heap
}

// Init 初始化堆结构
func (heap *BinaryHeap) Init() {
	size := len(heap.values)
	// size/2 - 1 是有叶子结点的最大编号结点
	for i := size/2 - 1; i >= 0; i-- {
		heap.down(i)
	}
}

// Push 插入新数据
func (heap *BinaryHeap) Push(value Item) {
	heap.values = append(heap.values, value)
	heap.up(len(heap.values) - 1)
}

// Pop 弹出堆顶数据
func (heap *BinaryHeap) Pop() Item {
	if len(heap.values) == 0 {
		return nil
	}

	top := heap.values[0]
	heap.values[0] = heap.values[len(heap.values)-1]
	heap.values = heap.values[:len(heap.values)-1]
	heap.down(0)
	return top
}

func (heap *BinaryHeap) down(i int) {
	values := heap.values
	child := 2*i + 1
	for child < len(values) {
		if child+1 < len(values) && values[child+1].Less(values[child]) {
			child++
		}
		if !values[child].Less(values[i]) {
			break
		}
		values[i], values[child] = values[child], values[i]
		i = child
		child = 2*i + 1
	}
}

func (heap *BinaryHeap) up(i int) {
	values := heap.values
	parent := (i - 1) / 2
	for i != 0 && values[i].Less(values[parent]) {
		values[i], values[parent] = values[parent], values[i]
		i = parent
		parent = (i - 1) / 2
	}
}
