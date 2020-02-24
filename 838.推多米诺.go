import "strings"

/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] 推多米诺
 */

// @lc code=start
func pushDominoes(dominoes string) string {
	prev, now := 0, 0
	length := len(dominoes)
	ans := make([]string, length)
	// prev == now 不处理
	// 需要枚举 L R 的四种情况
	for now < length {
		char := dominoes[now]
		if char == 'R' {

		} else if char == 'L' {
			if dominoes[prev] == char {

			} else {

			}
		} else {

		}
	}
	// 剩下的 prev
	return strings.Join(ans, "")
}

// @lc code=end
