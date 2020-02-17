/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	pMap := getMap(p)
	want := len(p)
	ans := make([]int, 0)
	if len(s) < want {
		return ans
	}
	sMap := getMap(s[:want])
	// fmt.Println(sMap, pMap)
	i := want
	for i < len(s) {
		// fmt.Println(sMap, pMap)
		if check(sMap, pMap) {
			ans = append(ans, i-want)
		}
		sMap[string(s[i-want])]--
		sMap[string(s[i])]++
		i++
	}
	// fmt.Println(sMap, pMap)
	if check(sMap, pMap) {
		ans = append(ans, i-want)
	}
	return ans
}
func check(sMap, pMap map[string]int) bool {
	for k, v := range pMap {
		if v2, ok := sMap[k]; ok {
			if v2 != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
func getMap(str string) map[string]int {
	m := make(map[string]int)
	for _, char := range str {
		m[string(char)]++
	}
	return m
}
// @lc code=end

