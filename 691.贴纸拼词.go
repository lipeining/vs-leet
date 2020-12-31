/*
 * @lc app=leetcode.cn id=691 lang=golang
 *
 * [691] 贴纸拼词
 */

// @lc code=start
func minStickers(stickers []string, target string) int {
	// 无限背包问题
	// 背包在外层，值在内层，小到大
	// n := len(stickers)
	m := len(target)
	amount := 1 << m
	dp := make([]int, amount)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	// 长度为0的无需贴纸
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, sticker := range stickers {
		for status := 0; status < amount; status++ {
			if dp[status] == -1 {
				// 无可用贴纸，这个状态
				continue
			}
			curStatus := status
			// 在前段状态下往最终状态继续拼接
			for i := 0; i < len(sticker); i++ {
				for j := 0; j < m; j++ {
					// initialState >> j & 1 == 1 说明这个已经拼接过了
					// sticker[i] == target[j] 说明这个字符可以使用，换下一个字符
					if sticker[i] == target[j] && (curStatus&(1<<j)) == 0 {
						curStatus |= 1 << j
						break
					}
				}
			}
			if dp[curStatus] == -1 {
				dp[curStatus] = dp[status] + 1
			} else {
				dp[curStatus] = min(dp[curStatus], dp[status]+1)
			}
		}
	}
	ans := dp[amount-1]
	return ans
}

// @lc code=end

