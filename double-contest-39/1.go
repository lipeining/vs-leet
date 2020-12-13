package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDeletions("aababbab"))
	fmt.Println(minimumDeletions("bbaaaaabb"))
	fmt.Println(minimumDeletions("ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa"))
	// minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9)
	// minimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11)
	// minimumJumps([]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7)
}
func minimumJumps(forbidden []int, a int, b int, x int) int {
	f := make(map[int]bool)
	for _, num := range forbidden {
		f[num] = true
	}
	ans := math.MaxInt32
	var dfs func(pass map[int]bool, path int, now int)
	dfs = func(pass map[int]bool, path int, now int) {
		if now == x {
			ans = min(ans, path)
			// fmt.Println("ans--------")
			return
		}
		fmt.Print(now, "->")
		if pass[now] {
			// fmt.Println("pass--------")
			return
		}
		pass[now] = true
		next := now + a
		if next <= x+b && !f[next] {
			dfs(pass, path+1, next)
		}
		prev := now - b
		if prev >= 0 && !f[prev] {
			dfs(pass, path+1, prev)
		}
		pass[now] = false
		// fmt.Println("false-------")
	}
	pass := make(map[int]bool)
	dfs(pass, 0, 0)
	fmt.Println("ans", ans)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
func minimumDeletions(s string) int {

}
func minimumDeletions2(s string) int {
	n := len(s)
	counter := make(map[byte]int)
	i := 0
	for i < n {
		if s[i] == 'b' {
			break
		}
		i++
	}
	j := n - 1
	for j >= i {
		if s[j] == 'a' {
			break
		}
		j--
	}
	fmt.Println(i, j)
	for i <= j {
		counter[s[i]]++
		i++
	}
	fmt.Println(counter)
	return min(counter['b'], counter['a'])
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
