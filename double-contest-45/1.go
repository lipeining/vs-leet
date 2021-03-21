package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// func2()
	// func3()
	func4()
}
func func4() {
	// 	输入：events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
	// 	输出：7
	// 	解释：选择绿色的活动会议 0 和 1，得到总价值和为 4 + 3 = 7 。
	// 	输入：events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
	// 输出：10
	// 解释：参加会议 2 ，得到价值和为 10 。
	// 你没法再参加别的会议了，因为跟会议 2 有重叠。你 不 需要参加满 k 个会议。

	// 输入：events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
	// 输出：9
	// 解释：尽管会议互不重叠，你只能参加 3 个会议，所以选择价值最大的 3 个会议。
	maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2)
	maxValue([][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}, 2)
	maxValue([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 3)
}
func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	fmt.Println(events)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// ans := dp[1][1]
	// for i := 1; i <= n; i++ {
	// 	dp[i][1] = max(dp[i][1], events[i-1][2])
	// 	ans = max(ans, dp[i][1])
	// 	for j := i + 1; j <= n; j++ {
	// 		dp[j][1] = max(dp[j][1], dp[i][1])
	// 		ans = max(ans, dp[j][1])
	// 		for x := 2; x <= k; x++ {
	// 			if events[j-1][0] > events[i-1][1] {
	// 				// fmt.Println("j can come from i", j, i)
	// 				dp[j][x] = max(dp[j][x], dp[i][x-1] + events[j-1][2])
	// 				ans = max(ans, dp[j][x])
	// 			}
	// 		}
	// 	}
	// }
	// // fmt.Println("dddddpppp", dp)
	// fmt.Println("anssssss", ans)
	// return ans
	// dp[1][1] = events[0][2]
	// for i := 1; i <= n; i++ {
	// 	for j := 1; j < i; j++ {
	// 		if events[i-1][0] >= events[j-1][1] {
	// 			dp[i][k] = max(dp[i][j], dp[j][k-1]+events[i-1][2])
	// 			fmt.Println("for i", i, "connect j", j, "result in", dp[i][k])
	// 		} else {
	// 			fmt.Println("for i not from j", i, j, dp[i][k])
	// 		}
	// 	}
	// }
	// fmt.Println("anssssss", dp)
	// return 0
}
func func3() {
	minimumLength("ca")
	minimumLength("cabaabac")
	minimumLength("aabccabba")
}
func minimumLength(s string) int {
	n := len(s)
	fmt.Println("s is ", s, n)
	if n <= 1 {
		return n
	}
	if s[0] != s[n-1] {
		return n
	}
	it := s[0]
	i := 0
	for i < n {
		if s[i] != it {
			break
		}
		i++
	}
	j := n - 1
	for j >= i {
		if s[j] != it {
			break
		}
		j--
	}
	fmt.Println(i, j)
	return minimumLength(s[i : j+1])
}
func func2() {
	// 输入：nums = [1,-3,2,3,-4]
	// 输出：5
	// 解释：子数组 [2,3] 和的绝对值最大，为 abs(2+3) = abs(5) = 5 。
	// 示例 2：

	// 输入：nums = [2,-5,1,-4,3,-2]
	// 输出：8
	// 解释：子数组 [-5,1,-4] 和的绝对值最大，为 abs(-5+1-4) = abs(-8) = 8 。
	maxAbsoluteSum([]int{1, -3, 2, 3, -4})
	maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2})
}
func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	g := math.MinInt32
	l := math.MaxInt32
	cg := 0
	cl := 0
	for i := 0; i < n; i++ {
		cg = max(cg+nums[i], nums[i])
		g = max(g, cg)
		cl = min(cl+nums[i], nums[i])
		l = min(l, cl)
	}
	fmt.Println(cg, g)
	fmt.Println(cl, l)
	ans := max(abs(g), abs(l))
	fmt.Println(ans)
	return ans
	// pre := make([]int, n+1)
	// for i := 0; i < n; i++ {
	// 	pre[i+1] = pre[i] + nums[i]
	// }
	// 以 i 结尾的最大列表，最小列表
	// ans := math.MinInt32
	// for i := 0; i <= n; i++ {
	// 	for j := i; j <= n; j++ {
	// 		a := pre[j] - pre[i]
	// 		// fmt.Println("form ", i, "to", j, "a is", a)
	// 		ans = max(ans, abs(a))
	// 	}
	// }
	// fmt.Println("ansssssss", ans)
	// return ans
}
