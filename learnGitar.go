package main

import "fmt"

// func main() {
// 	learnGitar(10, 5, 5, 3)
// }
func learnGitar(x, a, b, n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	// 如果做 memo 的话，需要使用 power, hour 做二维数组
	var dfs func(power int, hour int, learn int)
	dfs = func(power int, hour int, learn int) {
		if hour == n {
			fmt.Println("learn", learn, "ans", ans)
			ans = max(ans, learn)
			return
		}
		// 选择学习或者休息
		if power >= a {
			dfs(power-a, hour+1, learn+power)
		}
		dfs(power+b, hour+1, learn)
	}
	dfs(x, 0, 0)
	fmt.Println("ansssssssssssss", ans)
	return ans
}
