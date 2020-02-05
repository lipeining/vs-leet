/*
 * @lc app=leetcode.cn id=779 lang=golang
 *
 * [779] 第K个语法符号
 */

// @lc code=start
func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	return (1 - K%2) ^ kthGrammar(N-1, (K+1)/2)
}

// @lc code=end
