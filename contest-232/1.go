package main

import "fmt"

func main() {
	// maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2)
	// maxAverageRatio([][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4)
	// maxAverageRatio([][]int{{583, 868}, {783, 822}, {65, 262}, {121, 508}, {461, 780}, {484, 668}}, 8)
	// 	[[583,868],[783,822],[65,262],[121,508],[461,780],[484,668]]
	// 8
	// 0.57472

	// 输入：nums = [1,4,3,7,4,5], k = 3
	// 输出：15
	// 解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15 。
	// 示例 2：

	// 输入：nums = [5,5,4,5,4,1,1,1], k = 0
	// 输出：20
	// 解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。
	maximumScore([]int{1, 4, 3, 7, 4, 5}, 3)
	maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0)
}
func maximumScore(nums []int, k int) int {
	// 需要用一种方式记录一段区间内的最小值。
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := l + i - 1
			if j >= n {
				break
			}
			if i > k || k > j {
				break
			}
			left := dp[i][j-1] / (l - 1)
			dp[i][j] = max(dp[i][j], min(left, nums[j])*l)
			ans = max(ans, dp[i][j])
		}
	}
	fmt.Println(dp)
	fmt.Println("ansss", ans)
	return 0
}
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	// dp[i][j] 前 i 个班级使用了 j 个额外学生得到的最大通过率
	// 区间 dp
	n := len(classes)
	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, extraStudents+1)
	}
	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + float64(classes[i-1][0])/float64(classes[i-1][1])
	}
	// k 的部分超时了，如何提高效率呢？
	// 一个班级添加学生之后，会增大通过率
	for i := 1; i <= n; i++ {
		for j := 1; j <= extraStudents; j++ {
			p := float64(classes[i-1][0]+j) / float64(classes[i-1][1]+j)
			dp[i][j] = max(dp[i][j], dp[i-1][0]+p)
			// for k := 0; k <= j; k++ {
			// 	// 当前班级有 k 个额外学生时。
			// 	p := float64(classes[i-1][0]+k) / float64(classes[i-1][1]+k)
			// 	dp[i][j] = max(dp[i][j], dp[i-1][j-k]+p)
			// }
		}
	}
	fmt.Println(dp)
	ans := dp[n][extraStudents] / float64(n)
	fmt.Println("ansssss", ans)
	return ans
}
func findCenter(edges [][]int) int {
	m := len(edges)
	// n := m + 1
	cnt := make(map[int]int)
	for i := 0; i < m; i++ {
		u, v := edges[i][0], edges[i][1]
		cnt[u]++
		cnt[v]++
	}
	ans := 0
	for k, v := range cnt {
		if v == m {
			ans = k
			break
		}
	}
	return ans
}
func func1() {
	fmt.Println("anssssss")
}
func areAlmostEqual(s1 string, s2 string) bool {
	n := len(s1)
	cnt := 0
	l, r := -1, -1
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			cnt++
			if l == -1 {
				l = i
			} else if r == -1 {
				r = i
			} else {
				return false
			}
		}
	}
	// return s1[l] == s2[r]
	if l == -1 && r == -1 {
		return true
	} else if l == -1 && r != -1 {
		return false
	} else if l != -1 && r == -1 {
		return false
	} else {
		return s1[l] == s2[r]
	}
}
