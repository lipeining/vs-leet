/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i int, j int)bool {
		if len(d[i]) != len(d[j]) {
			return len(d[i]) > len(d[j])
		}
		return d[i] < d[j]
	})
	for i:=0;i<len(d);i++ {
		if helper(s,d[i]) {
			return d[i]
		}
	}
	return ""
}
func helper(s string,d string)bool {
	j:=0
	for i:=0;i<len(s) && j <len(d);i++ {
		if s[i] == d[j] {
			j++
		}
	}
	return j == len(d) 
}
// @lc code=end

