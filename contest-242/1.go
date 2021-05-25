package main

import (
	"fmt"
	"math"
)

func main() {
	// minSpeedOnTime([]int{1, 3, 2}, 6)
	// minSpeedOnTime([]int{1, 3, 2}, 2.7)
	// minSpeedOnTime([]int{1, 3, 2}, 1.9)
	canReach("011010", 2, 3)
	canReach("01101110", 2, 3)
	canReach("00111010", 3, 5)
	// stoneGameVIII([]int{-1, 2, -3, 4, -5})
	// stoneGameVIII([]int{7, -6, 5, 10, 5, -2, -6})
	// stoneGameVIII([]int{-10, -12})
	// stoneGameVIII([]int{-10, -12, -10, -12}) // 12
}
func stoneGameVIII(stones []int) int {
	n := len(stones)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + stones[i]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n+1)
	dp[n] = prefix[n]
	ans := dp[n]
	for i := n - 1; i > 1; i-- {
		dp[i] = prefix[i] - ans
		ans = max(ans, dp[i])
	}
	return ans
}
func stoneGameVIIIDFS(stones []int) int {
	n := len(stones)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + stones[i]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(index int) int
	dfs = func(index int) int {
		if index == n {
			return 0
		}
		if index == n-1 {
			return prefix[n] - prefix[0]
		}
		origin := prefix[index] - prefix[0]
		now := math.MinInt32
		for j := index + 2; j <= n; j++ {
			cur := prefix[j] - prefix[index]
			next := dfs(j)
			fmt.Println("index", "j", "origin", "cur", "next")
			fmt.Println(index, j, origin, cur, next)
			now = max(now, origin+cur-next)
		}
		fmt.Println("onnnnnnnnnnnnnn now", index, now)
		return now
	}
	ans := dfs(0)
	fmt.Println("anssssss", ans)
	return ans
}
func canReachDFS(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	memo := make([]bool, n)
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if s[index] == '1' {
			memo[index] = false
			return false
		}
		if index > 0 && memo[index] {
			return memo[index]
		}
		l := index + minJump
		r := index + maxJump
		if l > n-1 {
			memo[index] = false
			return false
		}
		if r >= n-1 {
			memo[index] = true
			return true
		}
		for j := l; j <= r; j++ {
			if dfs(j) {
				memo[index] = true
				return true
			}
		}
		memo[index] = false
		return false
	}
	ans := dfs(0)
	fmt.Println("anssssss", ans, memo)
	return ans
}
func canReachWD(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	memo := make([]bool, n)
	memo[0] = true
	// 没有考虑到超过边界的情况
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if memo[index] {
			return true
		}
		if s[index] == '1' {
			return false
		}
		l := index - maxJump
		if l < 0 {
			l = 0
		}
		r := index - minJump
		if r < 0 {
			r = 0
		}
		for j := l; j <= r; j++ {
			if dfs(j) {
				memo[index] = true
				return true
			}
		}
		memo[index] = false
		return false
	}
	ans := dfs(n - 1)
	fmt.Println("anssssss", ans, memo)
	return ans
}
func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	dp := make([]bool, n)
	if s[n-1] == '1' {
		return false
	}
	dp[0] = true
	pre := make([]int, n+1)
	for i := 0; i < minJump; i++ {
		pre[i] = 1
	}
	for i := minJump; i < n; i++ {
		l := i - maxJump
		r := i - minJump
		if s[i] == '1' {
			pre[i] = pre[i-1]
		} else {
			total := pre[r] - 0
			if l >= 1 {
				total = pre[r] - pre[l-1]
			}
			if total > 0 {
				dp[i] = true
				pre[i] = pre[i-1] + 1
			} else {
				pre[i] = pre[i-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1]
}
func minSpeedOnTime(dist []int, hour float64) int {
	// 二分法
	n := len(dist)
	if float64(n-1) > hour {
		return -1
	}
	check := func(speed int) bool {
		total := float64(0)
		for i := 0; i < n; i++ {
			dis := dist[i]
			t := dis / speed
			if dis%speed == 0 {
				total += float64(t)
			} else {
				if i != n-1 {
					total += float64(t + 1)
				} else {
					total += float64(dis) / float64(speed)
				}
			}
		}
		// fmt.Println("on speed", speed, "t", total, total < hour)
		return total <= hour
	}
	l := 1
	r := int(1e9)
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println("lllllllllllllll", l)
	return l
}
func checkZeroOnes(s string) bool {
	one := 0
	zero := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)
	i := 0
	for i < n {
		now := s[i]
		j := i
		for j < n {
			if s[j] == now {
				j++
			} else {
				break
			}
		}
		size := j - i + 1
		if now == '0' {
			zero = max(zero, size)
		} else {
			one = max(one, size)
		}
		i = j
	}
	return one > zero
}
