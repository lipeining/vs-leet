/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 */

// @lc code=start
func rotatedDigits(N int) int {
	ans := 0
	for i:=1; i<=N;i++ {
		if helper(i, false) {
			ans++
		}
	}
	return ans
}
func helper(n int, flag bool) bool {
	if n == 0 {
		return flag
	}
	str := n%10
	if str == 3 || str==4 || str==7 {
		return false
	} else if str == 0 || str == 1 || str == 8{
		return helper(n/10, flag)
	}
	return helper(n/10, true)
}
// @lc code=end

