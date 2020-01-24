/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res,multi := 10, 9
	if n > 10 {
		n = 10
	}
	for i:=1;i<n;i++ {
		multi*= 10-i
		res +=multi
	}
	return res


	// // 初始化的没有 0 开头的数字，但是递归中有0
	// // 相当于全排列，所以需要有 Used
	// ans := 0
	// var dfs func(n int, now int, used []bool, path string)
	// dfs = func(n int, now int, used []bool, path string) {
	// 	if now == n {
	// 		fmt.Println(path)
	// 		ans++
	// 		return
	// 	}
	// 	for i:=0;i<=9;i++ {
	// 		if used[i] {
	// 			continue
	// 		}
	// 		used[i] = true
	// 		dfs(n, now+1, used, path + strconv.Itoa(i))
	// 		used[i] = false
	// 	}
	// }
	// for i:=1;i<=n;i++ {
	// 	used := make([]bool, 10)
	// 	dfs(i, 0, used, "")
	// }
	// // used := make([]bool, 10)
	// // dfs(n, 0, used, "")
	// return ans
}
// @lc code=end

