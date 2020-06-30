package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Consistent struct {
	replicas int               // 每个结点的虚拟结点数
	circle   map[uint32]string // 结点哈希值和结点名的映射关系
	hashes   []uint32          // 按顺序存储环上结点哈希值
}

func New() *Consistent {
	return &Consistent{
		replicas: 20,
		circle:   make(map[uint32]string),
	}
}

func (c *Consistent) Add(elt string) {
	for i := 0; i < c.replicas; i++ {
		c.circle[c.hashKey(c.eltKey(elt, i))] = elt
	}
	c.updateHashes()
}

func (c *Consistent) Remove(elt string) {
	for i := 0; i < c.replicas; i++ {
		delete(c.circle, c.hashKey(c.eltKey(elt, i)))
	}
	c.updateHashes()
}

func (c *Consistent) Get(name string) string {
	if len(c.circle) == 0 {
		return ""
	}

	key := c.hashKey(name)
	idx := sort.Search(len(c.hashes), func(i int) bool {
		return c.hashes[i] > key
	})

	// 如果没找到，回到环的起点
	if idx == len(c.hashes) {
		idx = 0
	}

	return c.circle[c.hashes[idx]]
}

// eltKey 生成虚拟结点的标签名字 redis1#1 redis1#2, redis2#1
func (c *Consistent) eltKey(elt string, idx int) string {
	return elt + "#" + strconv.Itoa(idx)
}

// hashKey crc32算法生成uint32的hash值
func (c *Consistent) hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

// updateHashes 对环上的所有结点哈希值重新排序
func (c *Consistent) updateHashes() {
	var hashes []uint32
	for key := range c.circle {
		hashes = append(hashes, key)
	}
	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i] < hashes[j]
	})

	c.hashes = hashes
}

func main() {
	c := New()
	c.Add("redis0")
	c.Add("redis1")
	c.Add("redis2")
	c.Add("redis3")

	fmt.Println(c.Get("fan"))
	fmt.Println(c.Get("abcdefg"))
	fmt.Println(c.Get("1234567"))
	fmt.Println(c.Get("!@#$%^&"))
}
