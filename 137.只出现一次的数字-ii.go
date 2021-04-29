/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	ans := int32(0)
	for j := 0; j < 32; j++ {
		sum := int32(0)
		for _, num := range nums {
			sum += (int32(num) >> j) & 1
		}
		if sum%3 == 1 {
			ans |= 1 << j
		}
	}
	return int(ans)
}

// @lc code=end

