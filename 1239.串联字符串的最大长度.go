/*
 * @lc app=leetcode.cn id=1239 lang=golang
 *
 * [1239] 串联字符串的最大长度
 */

// @lc code=start
func maxLength(arr []string) int {
	// 别人解答的思路
	// 第 I 个字符串可以选择或者不选择，
	// 从其中得到较大的值
	// 回溯需要剪枝
	counter := &[26]int{}
	return dfs(arr, 0, counter)
}
func dfs(arr []string, curIndex int, counter *[26]int) int {
	length := len(arr)
	if curIndex == length {
		return 0
	}
	tmp := &[26]int{}
	for i := 0; i < 26; i++ {
		tmp[i] = counter[i]
	}
	unique := isUnique(arr[curIndex], tmp)
	if !unique {
		return dfs(arr, curIndex+1, counter)
	}
	len1 := len(arr[curIndex]) + dfs(arr, curIndex+1, tmp)
	len2 := dfs(arr, curIndex+1, counter)
	return max(len1, len2)
}
func isUnique(str string, counter *[26]int) bool {
	for i := 0; i < len(str); i++ {
		counter[str[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if counter[i] > 1 {
			return false
		}
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
