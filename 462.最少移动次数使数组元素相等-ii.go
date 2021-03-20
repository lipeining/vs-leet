/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start
func minMoves2(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	mid := nums[n/2]
	cnt := 0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for _, num := range nums {
		cnt += abs(mid - num)
	}
	return cnt
}

// @lc code=end

