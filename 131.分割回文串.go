/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	ans := make([][]string, 0)
	var dfs func(s string, index int, path []string)
	dfs = func(s string, index int, path []string) {
		if index == len(s) {
			ans = append(ans, path)
			return
		}
		// 提前判断是否是 回文串 再前进
		// 按照长度来划分
		for i:=1;i<=len(s)-index;i++ {
			if helper(s[index:index+i]) {
				tmp := make([]string, len(path))
				copy(tmp, path)
				tmp = append(tmp, s[index:index+i])
				dfs(s, index+i, tmp)
			}
		}
	}
	path := make([]string, 0)
	dfs(s, 0, path)
	return ans
}
func helper(s string) bool {
	length := len(s)
	for i:=0;i<length/2;i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}
// @lc code=end

