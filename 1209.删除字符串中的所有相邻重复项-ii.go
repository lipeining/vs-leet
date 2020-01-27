/*
 * @lc app=leetcode.cn id=1209 lang=golang
 *
 * [1209] 删除字符串中的所有相邻重复项 II
 */

// @lc code=start
func removeDuplicates(s string, k int) string {
	stack := make([]string, 0)
	for i:=0;i<len(s);i++ {
		length := len(stack)
		if length >= k-1 {
			if stack[length-1] != string(s[i]) {
				stack = append(stack, string(s[i]))
				continue
			}
			flag := true
			for j:=length-1;j>= length-k+1;j-- {
				if stack[j] != string(s[i]) {
					flag = false
					break
				}
			}
			if flag {
				stack = stack[:length-k+1]
			} else {
				stack = append(stack, string(s[i]))
			}
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	return strings.Join(stack, "")
}
// @lc code=end

