/*
 * @lc app=leetcode.cn id=956 lang=golang
 *
 * [956] 最高的广告牌
 */

// @lc code=start
func tallestBillboard(rods []int) int {
	sum := 0
	for _, rod := range rods {
		sum += rod
	}
	n := len(rods)
	if n <= 1 {
		return 0
	}
	half := sum / 2
	ans := 0
	for end := half; end >= 2; end-- {
		if end%2 == 1 {
			continue
		}
		dp := make([]bool, end+1)
		// 01 背包问题只可以使用一次
		dp[0] = true
		for i := 0; i < n; i++ {
			for j := end; j >= rods[i]; j-- {
				dp[j] = dp[j] || dp[j-rods[i]]
			}
		}
		if dp[end] {
			ans = end
			break
		}
	}
	return ans
}

// @lc code=end

