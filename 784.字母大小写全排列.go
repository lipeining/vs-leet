/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
func letterCasePermutation(S string) []string {
	n := len(S)
	ans := make([]string, 0)
	var dfs func(i int, path string)
	dfs = func(i int, path string) {
		if i == n {
			ans = append(ans, path)
			return
		}
		right := ""
		left := string(S[i])
		if unicode.IsLower(rune(S[i])) {
			right = string(unicode.ToUpper(rune(S[i])))
		} else if unicode.IsUpper(rune(S[i])) {
			right = string(unicode.ToLower(rune(S[i])))
		} else {
			right = string(rune(S[i]))
		}
		if left == right {
			dfs(i+1, path+left)
		} else {
			dfs(i+1, path+left)
			dfs(i+1, path+right)
		}
	}
	dfs(0, "")
	return ans
}

// @lc code=end

