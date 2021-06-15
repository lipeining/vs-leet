/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	for i < len(s) && s[i] >= s[i-1] {
			i++
	}
	if i < len(s) {
			for i > 0 && s[i] < s[i-1] {
					s[i-1]--
					i--
			}
			for i++; i < len(s); i++ {
					s[i] = '9'
			}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/monotone-increasing-digits/solution/dan-diao-di-zeng-de-shu-zi-by-leetcode-s-5908/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}
// @lc code=end

