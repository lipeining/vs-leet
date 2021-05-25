/*
 * @lc app=leetcode.cn id=960 lang=golang
 *
 * [960] 删列造序 III
 */

// @lc code=start
func minDeletionSize(A []string) int {
	W := len(A[0])
	dp := make([]int, W)
	for i:=0;i<W;i++ {
		dp[i] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i:=W-2;i>=0;i-- {
		for j:=i+1;j<W;j++ {
			flag := true
			if _, row := range A {
				if row[i] > row[j] {
					flag = false
				}
			}
			if flag {
				dp[i] = max(dp[i], 1 + dp[j])
			}
		}
	}
	kept := 0
	for _, x := range dp {
		kept = max(kept, x)
	}
	return w - kept
}
func minDeletionSizeOld(A []string) int {
	w := len(A[0])
	dp := make([]int, w)
	for i := 0; i < w; i++ {
		dp[i] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := w - 2; i >= 0; i-- {
		for j := i + 1; j < w; j++ {
			flag := false
			for _, row := range A {
				if row[i] > row[j] {
					flag = true
					break
				}
			}
			if flag {
				break
			}
			dp[i] = max(dp[i], 1+dp[j])
		}
	}
	kept := 0
	for _, x := range dp {
		kept = max(kept, x)
	}
	return w - kept
}

// @lc code=end

