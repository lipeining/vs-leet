/*
 * @lc app=leetcode.cn id=791 lang=golang
 *
 * [791] 自定义字符串排序
 */

// @lc code=start
func customSortString(order string, str string) string {
	index := make(map[byte]int)
	for i:= 0; i<len(order);i++{
		index[order[i]] = i+1
	}
	tosort := []byte(str)
	sort.Slice(tosort, func(i, j int)bool {
		return index[tosort[i]] < index[tosort[j]]
	})
	// fmt.Print()
	ans := ""
	for _, chr := range tosort {
		ans += string(chr)
	}
	return ans
}
// @lc code=end

