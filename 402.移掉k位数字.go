/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	n := len(num)
	remain := n - k
	stack := make([]byte, 0)
	for _, digit := range num {
		for k > 0 && len(stack) != 0 && stack[len(stack)-1] > byte(digit) {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, byte(digit))
	}
	ans := ""
	zero := true
	for i := 0; i < len(stack) && i < remain; i++ {
		if stack[i] == '0' && zero {
			continue
		} else {
			zero = false
			ans += string(stack[i])
		}
	}
	if ans == "" {
		return "0"
	}
	return ans
}

// @lc code=end

