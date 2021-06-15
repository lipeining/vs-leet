/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	getId := func(x, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1
	}
	set := make(map[int]int)
	for i, num := range nums {
		id := getId(num, t+1)
		if _, ok := set[id]; ok {
			return true
		}
		// left right
		if l, ok := set[id-1]; ok && abs(l-num) <= t {
			return true
		}
		if r, ok := set[id+1]; ok && abs(r-num) <= t {
			return true
		}
		set[id] = num
		if i >= k {
			delete(set, getId(nums[i-k], t+1))
		}
	}
	return false
}

// @lc code=end

