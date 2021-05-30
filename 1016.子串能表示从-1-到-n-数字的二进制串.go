/*
 * @lc app=leetcode.cn id=1016 lang=golang
 *
 * [1016] 子串能表示从 1 到 N 数字的二进制串
 */

// @lc code=start
func queryString(s string, n int) bool {
	var tmp string
	for i := 1; i <= n; i++ {
		tmp = fmt.Sprintf("%b", i)
		if !strings.Contains(s, tmp) {
			return false
		}
	}
	return true
	// 作者：meng-xiang-zhui-feng
	// 链接：https://leetcode-cn.com/problems/binary-string-with-substrings-representing-1-to-n/solution/go-by-meng-xiang-zhui-feng/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

// @lc code=end

