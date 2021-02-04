/*
 * @lc app=leetcode.cn id=1187 lang=golang
 *
 * [1187] 使数组严格递增
 */

// @lc code=start
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	// 他人答案
	n := len(arr1)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 1
	sort.Ints(arr2)
	upper := func(target int) int {
		// return -1
		l := 0
		r := len(arr2)
		for l < r-1 {
			mid := (l + r) / 2
			if arr2[mid] > target {
				r = mid
			} else {
				l = mid
			}
		}
		if arr2[l] > n {
			return l
		}
		return l + 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if arr1[i-1] > dp[j][i-1] {
				dp[j][i] = arr1[i-1]
			}
			if j > 0 {
				it := upper(dp[j-1][i-1])
				if it != -1 {
					dp[j][i] = min(dp[j][i], arr1[it])
				}
			}
			if i == n && dp[j][i] != math.MaxInt32 {
				return j
			}
		}
	}
	return -1
}

// @lc code=end

