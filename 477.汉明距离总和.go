/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	cnt := make([]int, 32)
	for i := 0; i < n; i++ {
		num := nums[i]
		for j := 0; j < 32 && num != 0; j++ {
			cnt[j] += num & 1
			num >>= 1
		}
	}
	ans := 0
	for _, c := range cnt {
		ans += c * (n - c)
	}
	return ans
}

// func totalHammingDistance(nums []int) int {
// 	n := len(nums)
// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			ans += hammingDistance(nums[i], nums[j])
// 		}
// 	}
// 	return ans
// }
// func hammingDistance(x int, y int) int {
// 	num := x ^ y
// 	ans := 0
// 	for num != 0 {
// 		num &= num - 1
// 		ans++
// 	}
// 	return ans
// }

// @lc code=end

