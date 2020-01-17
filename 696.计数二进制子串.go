/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	groups := make([]int,len(s))
	t := 0
	groups[0] = 1
	for i:=1;i<len(s);i++ {
		if s[i] != s[i-1] {
			t++
			groups[t] = 1
		} else {
			groups[t]++
		}
	}
	fmt.Println(groups)
	ans := 0
	for i:=1;i<len(groups);i++ {
		if groups[i] > groups[i-1] {
			ans += groups[i-1]
		} else {
			ans += groups[i]
		}
	}
	return ans
}
// @lc code=end

