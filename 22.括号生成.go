/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// ans := make([]string, 0)
	// if n == 0 {
	// 	return ans
	// }
	// var dfs func(n int, i int, str string)
	// dfs = func(n int, i int, str string) {
	// 	if n==i {
	// 		i:=0
	// 		for i<len(ans) {
	// 			if ans[i] == str {
	// 				break
	// 			}
	// 			i++
	// 		}
	// 		if i == len(ans) {
	// 			ans = append(ans, str)
	// 		}	
	// 		return
	// 	}
	// 	// 包裹 平行 两种方法
	// 	// 如何剪枝？
	// 	dfs(n, i+1, "("+ str + ")")
	// 	dfs(n, i+1, "()"+str)
	// 	dfs(n, i+1, str + "()")
	// }
	// dfs(n, 0, "")
	// return ansans := make([]string, 0)
	if n == 0 {
		return ans
	}
	var dfs func(l int, r int, str string)
	dfs = func(l int, r int, str string) {
		if l == 0 && r == 0 {
			ans = append(ans, str)
			return
		}
		// 包裹 平行 两种方法
		if l > r {
			return
		}
		if l > 0 {
			dfs(l-1, r, str + "(")
		}
		if r > 0 {
			dfs(l, r-1, str + ")")
		}
	}
	dfs(n, n, "")
	return ans
	
}
// @lc code=end

