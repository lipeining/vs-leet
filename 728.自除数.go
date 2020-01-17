/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	ans := make([]int,0)
	for i:=left;i<=right;i++ {
		if helper(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
func helper(n int) bool {
	str := strconv.Itoa(n)
	for _,d := range str {
		if d=='0' {
			return false
		}
		t, _ := strconv.Atoi(string(d))
		if n % t != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

