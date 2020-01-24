/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	ans := false
	// 关键在于 prev next 挑选，这样的话，
	// 处理那个 0 开头的数字呢？
	var dfs func(num string, index int, prev string, next string, tobe string)
	dfs = func(num string, index int, prev string, next string, tobe string) {
		if index >= len(num) {
			ans = true
			return
		}
		// fmt.Println(num, index, prev, next, tobe)
		// 取一个合理的数
		if index + len(tobe) > len(num) {
			return
		}
		if !checkEqual(num[index:index+len(tobe)],tobe) {
			return
		}
		nextTobe := getTobe(next, tobe)
		dfs(num, index+len(nextTobe), next, tobe, nextTobe)
	}
	// 这里可以初始化多种搜索路径，需要符合正常的数字开头即可
	for x:=1;x<len(num);x++ {
		for y:=1;y<len(num);y++ {
			nextPrev := ""
			nextNext := ""
			if x < len(num) {
				nextPrev = num[0:x]
			}
			if x+y < len(num) {
				nextNext = num[x:x+y]
			}
			nextTobe := getTobe(nextPrev, nextNext)
			if check(nextPrev) && check(nextNext) {
				dfs(num, x + y, nextPrev, nextNext, nextTobe)
			}
		}
	}
	return ans
}
func checkEqual(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for j:=0;j<len(b);j++ {
		if a[j]!=b[j]{
			return false
		}
	}
	return true
}
func check(str string) bool {
	// 允许 0 不允许 0 开头
	if str == "" {
		return false
	}
	if len(str) > 1 && str[0]=='0' {
		return false
	}
	return true
}
func getTobe(prev string, next string) string {
	p,_ := strconv.Atoi(prev)
	n,_ := strconv.Atoi(next)
	t := p+n
	ans := strconv.Itoa(t)
	// fmt.Println(prev, next, ans)
	return ans
}
// @lc code=end

