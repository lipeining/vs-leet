package main

import "fmt"

func main() {
	// func1()
	func2()
	func3()
	func4()
	// ans := 0
	// fmt.Println("ans---------------------", ans)
}
func func1() {
	// 输入：n = 4
	// 输出：10
	// 解释：第 4 天后，总额为 1 + 2 + 3 + 4 = 10 。
	// 示例 2：

	// 输入：n = 10
	// 输出：37
	totalMoney(4)
	totalMoney(10)
	totalMoney(20)
}
func totalMoney(n int) int {
	if n <= 1 {
		return n
	}
	days := make([]int, 8)
	days[1] = 1
	ans := 1
	for i := 2; i <= n; i++ {
		day := i % 7
		if day == 0 {
			days[0] = days[6] + 1
		} else if day == 1 {
			days[1] = days[1] + 1
		} else {
			days[day] = days[day-1] + 1
		}
		ans += days[day]
	}
	fmt.Println("ans-----------", ans)
	return ans
}
func func2() {
	// 输入：s = "cdbcbbaaabab", x = 4, y = 5
	// 输出：19
	// 解释：
	// - 删除 "cdbcbbaaabab" 中加粗的 "ba" ，得到 s = "cdbcbbaaab" ，加 5 分。
	// - 删除 "cdbcbbaaab" 中加粗的 "ab" ，得到 s = "cdbcbbaa" ，加 4 分。
	// - 删除 "cdbcbbaa" 中加粗的 "ba" ，得到 s = "cdbcba" ，加 5 分。
	// - 删除 "cdbcba" 中加粗的 "ba" ，得到 s = "cdbc" ，加 5 分。
	// 总得分为 5 + 4 + 5 + 5 = 19 。
	// 示例 2：

	// 输入：s = "aabbaaxybbaabb", x = 5, y = 4
	// 输出：20
	maximumGain("cdbcbbaaabab", 4, 5)
	maximumGain("aabbaaxybbaabb", 5, 4)
}
func maximumGain(s string, x int, y int) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	// 以 I 结尾的字符串的最大分数
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Println("dp", dp)
	ans := max(dp[n], dp[n-1])
	fmt.Println("ans-----------", ans)
	return ans
}
func func3() {

}
func func4() {

}
