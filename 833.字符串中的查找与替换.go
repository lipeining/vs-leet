/*
 * @lc app=leetcode.cn id=833 lang=golang
 *
 * [833] 字符串中的查找与替换
 */

// @lc code=start
func findReplaceString(s string, indexes []int, sources []string, targets []string) string {
	// n := len(s)
	m := len(indexes)
	type pair struct {
		index  int
		source string
		target string
	}
	plist := make([]pair, 0)
	for i := 0; i < m; i++ {
		index := indexes[i]
		source := sources[i]
		target := targets[i]
		if s[index:index+len(source)] == source {
			plist = append(plist, pair{
				index,
				source,
				target,
			})
		}
	}
	sort.Slice(plist, func(i, j int) bool {
		return plist[i].index < plist[j].index
	})
	ans := ""
	now := 0
	for i := 0; i < len(plist); i++ {
		p := plist[i]
		ans += s[now:p.index]
		ans += p.target
		now = p.index + len(p.source)
	}
	ans += s[now:]
	return ans
}

// @lc code=end

