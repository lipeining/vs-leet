/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	visited := make(map[string]int)
	want := 10
	length := len(s)
	for i := 0; i <= length-want; i++ {
		sub := s[i : i+want]
		visited[sub]++
	}
	// fmt.Println(visited)
	ans := make([]string, 0)
	for k,v := range visited {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

// @lc code=end
