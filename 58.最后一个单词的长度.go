/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	// 快慢指针吧
	slow := 0
	fast := 0
	length := len(s)
	res := 0
	for fast < length {
		if s[fast] == ' ' {
			// 假如需要更新旧的值得话，即非连续的空格时
			if fast != slow {
				res = fast - slow
			}
			fast = fast + 1
			slow = fast
		} else {
			fast = fast + 1
		}
	}
	if fast == slow {
		return res
	}
	return fast - slow 
}
// @lc code=end

