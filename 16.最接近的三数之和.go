import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	// 官方思路
	length := len(nums)
	// if length < 3 {
	// 	return ans
	// }
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < length-2; i++ {
		// if nums[i] > 0 {
		// 	break
		// }
		// if i > 0 && nums[i] == nums[i-1] {
		// 	continue
		// }
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if helper(sum, target, ans) {
				ans = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return ans
			}
		}
	}
	return ans
}
func helper(sum, target, ans int) bool {
	if math.Abs(float64(sum-target)) < math.Abs(float64(ans-target))  {
		return true
	}
	return false
}

// @lc code=end
