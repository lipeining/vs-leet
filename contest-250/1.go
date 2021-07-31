package main

import (
	"fmt"
	"strings"
)

func main() {
	// addRungs([]int{1, 3, 5, 10}, 2)
	// addRungs([]int{3, 6, 8, 10}, 3)
	// addRungs([]int{3, 4, 6, 7}, 2)
	// addRungs([]int{5}, 10)
	// addRungs([]int{3}, 1)
	maxPoints([][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}})
	maxPoints([][]int{{1, 5}, {2, 3}, {4, 2}})
}
func maxPoints(points [][]int) int64 {
	m := len(points)
	n := len(points[0])
	dp := make([][]int64, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int64, n)
	}
	max := func(a, b int64) int64 {
		if a > b {
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
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = int64(points[i][j])
			} else {
				// 从上一行中选择最合适的
				for y := 0; y < n; y++ {
					dp[i][j] = max(dp[i][j], dp[i-1][y]+int64(points[i][j])-int64(abs(j-y)))
				}
			}
		}
	}
	var ans int64
	for j := 0; j < n; j++ {
		ans = max(ans, dp[m-1][j])
	}
	// fmt.Println(dp, ans)
	return ans
}
func addRungs(rungs []int, dist int) int {
	n := len(rungs)
	dp := make([]int, n)
	need := func(a int, b int) int {
		if a <= b {
			return 0
		}
		d := a / b
		if a%b == 0 {
			d--
		}
		return d
	}
	dp[0] = need(rungs[0], dist)
	for i := 1; i < n; i++ {
		cur := need(rungs[i]-rungs[i-1], dist)
		dp[i] = dp[i-1] + cur
	}
	fmt.Println("dp", dp)
	return dp[n-1]
}
func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	bm := make([]bool, 26)
	for i := 0; i < len(brokenLetters); i++ {
		index := int(brokenLetters[i] - 'a')
		bm[index] = true
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			index := int(words[i][j] - 'a')
			if bm[index] {
				ans++
				break
			}
		}
	}
	ans = len(words) - ans
	return ans
}
