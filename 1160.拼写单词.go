/*
 * @lc app=leetcode.cn id=1160 lang=golang
 *
 * [1160] 拼写单词
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	charMap := getMap(chars)
	ans := 0
	for i:=0;i<len(words);i++ {
		wordMap := getMap(words[i])
		flag := true
		for k,v := range wordMap {
			c,ok := charMap[k]
			if ok && c >= v {
				// pass
			} else {
				flag = false
				break
			}
		}
		if flag {
			ans += len(words[i])
		}
	}
	return ans
}
func getMap(str string) map[byte]int {
	cMap := make(map[byte]int)
	for i:=0;i<len(str);i++ {
		n,ok:=cMap[str[i]]
		if ok {
			cMap[str[i]] = n + 1
		} else {
			cMap[str[i]] = 1
		}
	}
	return cMap
}
// @lc code=end

