/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(A string, B string) []string {
	aList := strings.Split(A, " ")
	bList := strings.Split(B, " ")
	cnt := make(map[string]int)
	for _, s := range aList {
		cnt[s]++
	}
	for _, s := range bList {
		cnt[s]++
	}
	ans := make([]string, 0)
	for k, v := range cnt {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
func uncommonFromSentencesOld(A string, B string) []string {
	aList := strings.Split(A, " ")
	bList := strings.Split(B, " ")
	aMap := getMap(aList)
	bMap := getMap(bList)
	ans := make([]string, 0)
	for k, a := range aMap {
		_, bOk := bMap[k]
		if a == 1 && !bOk {
			ans = append(ans, k)
		}
	}
	for k, b := range bMap {
		_, aOk := aMap[k]
		if b == 1 && !aOk {
			ans = append(ans, k)
		}
	}
	return ans
}
func getMap(list []string) map[string]int {
	t := make(map[string]int)
	for _, word := range list {
		n, ok := t[word]
		if ok {
			t[word] = n + 1
		} else {
			t[word] = 1
		}
	}
	return t
}

// @lc code=end

