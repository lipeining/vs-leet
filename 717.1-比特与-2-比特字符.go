/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
	length := len(bits)
	if length <=1 {
		return true
	}
	i:=0
	for i < length - 1 {
		if bits[i] == 0 {
			i++
		} else {
			i+=2
		}
	}
	return i < length
}
// @lc code=end

