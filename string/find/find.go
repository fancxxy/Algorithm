package find

// Native 朴素匹配
func Native(s, substr string) int {
	m, n := len(s), len(substr)

	for i := 0; i <= m-n; i++ {
		var j int
		for ; j < n; j++ {
			if s[i+j] != substr[j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return -1
}

// RabinKarp 算法
func RabinKarp(s, substr string) int {
	const base = 128
	hash := func(s string) (uint32, uint32) {
		var value, pow uint32 = 0, 1
		for i := 0; i < len(s); i++ {
			value = value*base + uint32(s[i])
			pow *= base
		}
		return value, pow
	}

	n := len(substr)
	if n == 0 {
		return 0
	} else if n == len(s) {
		if s == substr {
			return 0
		}
		return -1
	} else if n > len(s) {
		return -1
	}

	value, pow := hash(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*base + uint32(s[i])
	}
	if h == value && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h = h*base + uint32(s[i]) - uint32(s[i-n])*pow
		i++
		if h == value && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}
