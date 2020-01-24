/*
 * @lc app=leetcode.cn id=842 lang=golang
 *
 * [842] 将数组拆分成斐波那契序列
 */

// @lc code=start
func splitIntoFibonacci(S string) []int {
	// 关键在于 prev next 挑选，这样的话，
	// 处理那个 0 开头的数字呢？
	res := make([]int, 0)
	num := S
	var dfs func(num string, index int, prev string, next string, tobe string, path []int)
	dfs = func(num string, index int, prev string, next string, tobe string, path []int) {
		if index > len(num) {
			res = path
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
		// tobe 验证成功
		nextTobe := getTobe(next, tobe)
		t,_ := strconv.Atoi(tobe)
		tmp:=make([]int, len(path))
		copy(tmp, path)
		tmp = append(tmp, t)
		dfs(num, index+len(nextTobe), next, tobe, nextTobe, tmp)
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
			if checkStr(nextPrev) && checkStr(nextNext) {
				path := make([]int, 2)
				path[0],_ = strconv.Atoi(nextPrev)
				path[1],_ = strconv.Atoi(nextNext)
				nextTobe := getTobe(nextPrev, nextNext)
				dfs(num, x + y, nextPrev, nextNext, nextTobe, path)
			}
		}
	}
	return res
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
func checkStr(str string) bool {
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

