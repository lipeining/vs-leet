import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1048 lang=golang
 *
 * [1048] 最长字符串链
 */

// @lc code=start
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	// fmt.Println(words)
	length := len(words)
	dp := make([]int, length)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if helper(words[i], words[j]) {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}
	ans := 0
	for i := 0; i < length; i++ {
		ans = max(ans, dp[i])
	}
	return ans + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func helper(left, right string) bool {
	if len(left) != len(right)-1 {
		return false
	}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] == right[j] {
			i++
		}
		j++
	}
	return i == len(left)
}

// @lc code=end
