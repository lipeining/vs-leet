/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	n := len(strs)
	cnt := make(map[string][]string)
	for i := 0; i < n; i++ {
		word := []byte(strs[i])
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		w := string(word)
		cnt[w] = append(cnt[w], strs[i])
	}
	for _, v := range cnt {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end

