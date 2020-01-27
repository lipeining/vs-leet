/*
 * @lc app=leetcode.cn id=880 lang=golang
 *
 * [880] 索引处的解码字符串
 */

// @lc code=start
func decodeAtIndex(S string, K int) string {
	stack := ""
	for i:=0;i<len(S);i++ {
		if unicode.IsDigit(rune(S[i])) {
			time,_ := strconv.Atoi(string(S[i]))
			stack = strings.Repeat(stack, time)
		} else {
			stack += string(S[i])
		}
		// fmt.Println(stack)
		if len(stack) >= K {
			return string(stack[K-1])
		}
	}
	return ""
}
// @lc code=end

