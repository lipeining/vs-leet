/*
 * @lc app=leetcode.cn id=893 lang=golang
 *
 * [893] 特殊等价字符串组
 */

// @lc code=start
func numSpecialEquivGroups(A []string) int {
	n := len(A)
	cnt := make(map[string]int)
	format := func(s string) string {
		odd := make([]string, 0)
		even := make([]string, 0)
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				even = append(even, string(s[i]))
			} else {
				odd = append(odd, string(s[i]))
			}
		}
		sort.Strings(odd)
		sort.Strings(even)
		return strings.Join(odd, "") + strings.Join(even, "")
	}
	for i := 0; i < n; i++ {
		s := format(A[i])
		cnt[s]++
	}
	return len(cnt)
}

// @lc code=end

