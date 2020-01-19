/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短完整词
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	origin := getMap(licensePlate)
	ans := licensePlate + "temp"
	// fmt.Println(origin)
	for i:=0; i < len(words); i++ {
		if helper(origin, words[i]) && len(ans) > len(words[i]) {
			ans = words[i]
		}
	}
	return ans
}

func helper(origin map[string]int, word string) bool {
	counter := getMap(word)
	fmt.Println(origin, counter)
	for k, v := range origin {
		w,ok := counter[k]
		if !ok {
			return false
		}
		if w < v {
			return false
		}
	}
	return true
}
func getMap(str string) map[string]int {
	counter := make(map[string]int)
	for _,c := range str {
		if unicode.IsUpper(c) || unicode.IsLower(c){
			lower := unicode.ToLower(c)
			n,ok := counter[string(lower)]
			if ok {
				counter[string(lower)] = n + 1
			} else {
				counter[string(lower)] = 1
			}
		}
	}
	return counter
}
// @lc code=end

