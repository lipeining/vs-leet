/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
	// n := len(nums) + 1
	// sum := (n - 1) * n / 2
	// ans := sum
	// for _, num := range nums {
	// 	ans -= num
	// }
	// return ans
}

// @lc code=end

