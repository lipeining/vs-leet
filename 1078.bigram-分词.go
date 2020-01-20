/*
 * @lc app=leetcode.cn id=1078 lang=golang
 *
 * [1078] Bigram 分词
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
	// sep := first + " " + second + " "
	// // aa good girl   first: a, second: good
	// tmpStr := text
	// index := strings.Index(tmpStr, sep)
	// lengthSep := len(sep)
	// ans := make([]string, 0)
	// for {
	// 	next := index+lengthSep
	// 	if next >= len(text) {
	// 		break
	// 	}
	// 	tmpStr = text[next:]
	// 	third := getThird(tmpStr)
	// 	ans = append(ans, third)
	// 	fmt.Println(tmpStr, third)
	// 	newIndex := strings.Index(tmpStr, sep)
	// 	if newIndex == -1 {
	// 		break
	// 	}
	// 	index = next + newIndex
	// }
	// return ans
	ans := make([]string, 0)
	words := strings.Split(text, " ")
	for i:=0;i<len(words)-2;i++ {
		if string(words[i]) == first && string(words[i+1]) == second {
			ans = append(ans, string(words[i+2]))
		}
	}
	return ans
}
func getThird(text string) string {
	str := ""
	for i:=0;i<len(text);i++ {
		if text[i] != ' ' {
			str += string(text[i])
		} else {
			break
		}
	}
	// fmt.Println(text, str)
	return str
}
// @lc code=end

