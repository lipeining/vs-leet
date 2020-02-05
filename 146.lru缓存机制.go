import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
type LRUCache struct {
	capacity int
	queue    *list.List
	table    map[int]int
}

// l := list.New()
// e4 := l.PushBack(4)
// e1 := l.PushFront(1)
// l.InsertBefore(3, e4)
// l.InsertAfter(2, e1)

// // Iterate through list and print its contents.
// for e := l.Front(); e != nil; e = e.Next() {
//     fmt.Println(e.Value)
// }

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, queue: list.New(), table: make(map[int]int)}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.table[key]
	if ok {
		e := this.queue.Front()
		for e != nil {
			k := e.Value.(int)
			if k == key {
				break
			}
			e = e.Next()
		}
		this.queue.MoveToBack(e)
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.queue.PushBack(key)
	this.table[key] = value
	length := this.queue.Len()
	if length > this.capacity {
		front := this.queue.Front()
		this.queue.Remove(front)
		rmKey := front.Value.(int)
		delete(this.table, rmKey)
	}
	fmt.Println(this.table)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
