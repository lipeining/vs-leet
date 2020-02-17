/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	length := len(nums)
	ans := 0
	m := make(map[int]int)
	sum := 0
	m[0] = 1
	for i := 0; i < length; i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			ans += v
		}
		m[sum]++
	}
	return ans
	// length := len(nums)
	// ans := 0
	// for start := 0; start < length; start++ {
	// 	sum := 0
	// 	for end := start; end < length; end++ {
	// 		sum += nums[end]
	// 		if sum == k {
	// 			ans++
	// 		}
	// 	}
	// }
	// return ans
}

// @lc code=end
