/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */
package main

import (
	"container/list"
)

// func main() {
// 	lru := NewLRUCache(2)
// 	print := func() {
// 		for e := lru.list.Front(); e != nil; e = e.Next() {
// 			fmt.Print(e.Value, "->")
// 		}
// 		fmt.Println("endddddddd")
// 	}
// 	lru.Put(1, 1) // 缓存是 {1=1}
// 	print()
// 	lru.Put(2, 2) // 缓存是 {1=1, 2=2}
// 	print()
// 	fmt.Println(lru.Get(1)) // 返回 1
// 	print()
// 	lru.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// 	print()
// 	fmt.Println(lru.Get(2)) // 返回 -1 (未找到)
// 	print()
// 	lru.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// 	print()
// 	fmt.Println(lru.Get(1)) // 返回 -1 (未找到)
// 	print()
// 	fmt.Println(lru.Get(3)) // 返回 3
// 	print()
// 	fmt.Println(lru.Get(4)) // 返回 4
// 	print()
// }

// @lc code=start

type LRUCache struct {
	hash     map[int]*list.Element
	list     *list.List
	capacity int
}

// func NewLRUCache(capacity int) LRUCache {
// 	return LRUCache{
// 		hash:     make(map[int]*list.Element),
// 		list:     list.New(),
// 		capacity: capacity,
// 	}
// }

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:     make(map[int]*list.Element),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.hash[key]; !ok {
		return -1
	} else {
		this.list.MoveToFront(e)
		return e.Value.(int)
	}
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.hash[key]; ok {
		// update
		this.list.Remove(e)
		ne := this.list.PushFront(value)
		this.hash[key] = ne
	} else {
		// insert
		ie := this.list.PushFront(value)
		this.hash[key] = ie
		if this.list.Len() > this.capacity {
			last := this.list.Back()
			this.list.Remove(last)
			for k, v := range this.hash {
				if v == last {
					delete(this.hash, k)
					break
				}
			}
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
