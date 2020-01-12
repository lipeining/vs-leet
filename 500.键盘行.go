/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
    smap := map[string]int{
		"q":1,
		"w":1,
		"e":1,
		"r":1,
		"t":1,
		"y":1,
		"u":1,
		"i":1,
		"o":1,
		"p":1,

		"a":2,
		"s":2,
		"d":2,
		"f":2,
		"g":2,
		"h":2,
		"j":2,
		"k":2,
		"l":2,

		"z":3,
		"x":3,
		"c":3,
		"v":3,
		"b":3,
		"n":3,
		"m":3,
	}
	res := make([]string,0)
	for _,word := range words {
		flag := true
		for i:= 0; i <len(word)-1;i++ {
			left := smap[strings.ToLower(string(word[i]))]
			right := smap[strings.ToLower(string(word[i+1]))]
			if left != right {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, word)
		}
	}
	return res
}
// @lc code=end

