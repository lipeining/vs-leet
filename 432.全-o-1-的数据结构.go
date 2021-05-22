/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 */

// @lc code=start
type AllOne struct {
	m   map[string]int
	rev map[int]map[string]bool
	max []int // 用于最大栈
	min []int // 用于最小栈
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		m:   make(map[string]int),
		rev: make(map[int]map[string]bool),
		max: make([]int, 0),
		min: make([]int, 0),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	this.m[key]++
	val := this.m[key]
	if this.rev[val] == nil {
		this.rev[val] = make(map[string]bool)
	}
	this.rev[val][key] = true
	maxs := len(this.max)
	if maxs == 0 {
		this.max = append(this.max, val)
	} else {
		top := this.max[maxs-1]
		if top < val {
			this.max = append(this.max, val)
		} else {
			this.max = append(this.max, top)
		}
	}
	mins := len(this.min)
	if mins == 0 {
		this.min = append(this.min, val)
	} else {
		top := this.min[mins-1]
		if top > val {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, top)
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if v, ok := this.m[key]; !ok {
		return
	} else {
		delete(this.rev[v], key)
		this.max = this.max[:len(this.max)-1]
		this.min = this.min[:len(this.min)-1]
		if v == 1 {
			delete(this.m, key)
		} else {
			if this.rev[v-1] == nil {
				this.rev[v-1] = make(map[string]bool)
			}
			this.rev[v-1][key] = true
			this.m[key]--
		}
	}

}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	maxs := len(this.max)
	if maxs == 0 {
		return ""
	}
	top := this.max[maxs-1]
	for k := range this.rev[top] {
		return k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	mins := len(this.min)
	if mins == 0 {
		return ""
	}
	top := this.min[mins-1]
	for k := range this.rev[top] {
		return k
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end

