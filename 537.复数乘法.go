/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 */

// @lc code=start
func complexNumberMultiply(a string, b string) string {
	getPair := func(str string) (int, int) {
		parts := strings.Split(str, "+")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1][:len(parts[1])-1])
		return left, right
	}
	i, j := getPair(a)
	m, n := getPair(b)
	x := i*m - j*n
	y := i*n + j*m
	fmt.Println("a", i, j)
	fmt.Println("b", m, n)
	ans := strconv.Itoa(x) + "+" + strconv.Itoa(y) + "i"
	fmt.Println("anssss", ans)
	return ans
}

// @lc code=end

 