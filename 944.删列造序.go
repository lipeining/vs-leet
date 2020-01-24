/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] 删列造序
 */

// @lc code=start
func minDeletionSize(A []string) int {
	length := len(A)
	if length == 0 {
		return 0
	}
	n := len(A[0])
	ans := 0
	for i:=0;i<n;i++ {
		flag := false
		for j:=0;j<length-1;j++ {
			if A[j][i] > A[j+1][i] {
				flag = true
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}
// @lc code=end

