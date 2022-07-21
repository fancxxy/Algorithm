package ringbuffer

type T = interface{}

type RingBuffer struct {
	buf     []T // 缓存数组
	size    int // 数组长度
	read    int // 下一次读取位置
	write   int // 下一次写入位置
	overlap bool
}

func New(size int, overlap bool) *RingBuffer {
	if size <= 0 {
		size = 16
	}

	size--
	size |= size >> 1
	size |= size >> 2
	size |= size >> 4
	size |= size >> 8
	size |= size >> 16
	size++

	return &RingBuffer{
		buf:     make([]T, size),
		size:    size,
		overlap: overlap,
	}
}

func (r *RingBuffer) Read(n int) []T {
	if n <= 0 {
		return []T{}
	}

	// 缓存数组为空，返回空数组
	if r.read == r.write {
		return []T{}
	}

	//  把读写索引映射到真实索引
	var (
		read  = r.read & (r.size - 1)
		write = r.write & (r.size - 1)
		count int
	)

	// 第一种情况
	if write > read {
		// 可读取的数量
		count = write - read
		if count > n {
			count = n
		}
		p := make([]T, count)

		copy(p, r.buf[read:read+count])
		r.read = (r.read + count) & (2*r.size - 1)
		return p
	}

	// 第二种情况
	count = r.size - read + write
	if count > n {
		count = n
	}
	p := make([]T, count)

	if read+count < r.size {
		copy(p, r.buf[read:read+count])
	} else {
		copy(p, r.buf[read:r.size])
		copy(p[r.size-read:], r.buf[0:count-r.size+read])
	}

	r.read = (r.read + count) & (2*r.size - 1)
	return p
}

func (r *RingBuffer) Write(p []T) int {
	if len(p) == 0 {
		return 0
	}

	//  把读写索引映射到真实索引
	var (
		read  = r.read & (r.size - 1)
		write = r.write & (r.size - 1)
		// count实际能写入的数据个数，space空闲位置
		count, space int
	)

	// 缓存数组满，如果不能覆盖旧数据，直接返回0
	if r.read^r.size == r.write {
		if !r.overlap {
			return 0
		}

		space = 0
	} else {
		// 计算空闲位置个数
		if write >= read {
			space = r.size - write + read
		} else {
			space = read - write
		}
	}

	// 保证一次写入的数据量小于缓存数组大小，在不允许覆盖的情况下写入的数据量最大不超过空闲位置数
	if r.overlap && len(p) > r.size {
		p = p[:r.size]
	} else if !r.overlap && len(p) > space {
		p = p[:space]
	}

	count = len(p)

	if write+count <= r.size {
		copy(r.buf[write:], p)
	} else {
		copy(r.buf[write:], p[:r.size-write])
		copy(r.buf[0:], p[r.size-write:])
	}

	// 如果写入的数据量大于空闲位置数，说明缓存数组满了，需要移动读索引
	if count > space {
		r.read = (r.read + count - space) & (2*r.size - 1)
	}

	r.write = (r.write + count) & (2*r.size - 1)
	return count
}

func (r *RingBuffer) Len() int {
	if r.read == r.write {
		return 0
	}

	if r.read < r.write {
		return r.write - r.read
	}

	return 2*r.size - r.read + r.write
}

func (r *RingBuffer) Reset() {
	r.read = 0
	r.write = 0
	r.buf = make([]T, r.size)
}
