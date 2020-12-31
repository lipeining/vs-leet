/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	// 简单的前缀和可以通过
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := pre[j+1] - pre[i]
			if k == 0 {
				if sum == 0 {
					return true
				}
			} else if sum%k == 0 {
				return true
			}
		}
	}
	return false
	// m := make(map[int]int)
	// // 前缀和的求余数 map
	// m[0] = -1
	// sum := 0
	// for i,num := range nums {
	// 	sum += num
	// 	if k != 0 {
	// 		sum %= k
	// 	}
	// 	if lastIndex, ok := m[sum]; ok {
	// 		if i - lastIndex > 1 {
	// 			return true
	// 		}
	// 	} else {
	// 		m[sum] = i
	// 	}
	// }
	// return false
}

// @lc code=end

