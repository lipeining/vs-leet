/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func tribonacci(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	ans := 0
	for i:=3;i<=n;i++ {
		tmp:= t0+t1+t2
		t0 = t1
		t1 = t2
		t2 = tmp
		ans = tmp
	}
	return ans
}
// @lc code=end

