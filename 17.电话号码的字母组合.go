/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {

	res := make([]string, 0)
	if digits == "" {
		return res	
	}
	list := [8]string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var dfs func(digits string, index int, ans string)
	dfs = func(digits string, index int, ans string) {
		if index == len(digits) {
			res = append(res, ans)
			return
		}
		curNumber, _ := strconv.Atoi(string(digits[index]))
		str := list[curNumber-2]
		for i:=0;i<len(str);i++ {
			dfs(digits, index+1, ans + string(str[i]))
		}
	}
	dfs(digits, 0, "")
	return res
}
// @lc code=end

