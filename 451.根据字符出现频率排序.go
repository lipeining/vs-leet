/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
func frequencySort(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	keys := make([]byte, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	ans := ""
	for _, k := range keys {
		ans += strings.Repeat(string(k), m[k])
	}
	return ans
}

// @lc code=end

