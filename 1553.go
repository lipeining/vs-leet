package main

import "fmt"

// func main() {
// 	minDays(10)
// 	minDays(6)
// 	minDays(1)
// 	minDays(56)
// }
func minDays(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2
	memo[3] = 2
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dfs func(num int) int
	dfs = func(num int) int {
		if need, ok := memo[num]; ok {
			return need
		}
		two := num
		three := num
		if num%2 == 0 {
			two = 1 + dfs(num/2)
		}
		if num%3 == 0 {
			three = 1 + dfs(num-2*num/3)
		}
		one := 1 + dfs(num-1)
		ret := min(one, min(two, three))
		memo[num] = ret
		return ret
	}
	ans := dfs(n)
	fmt.Println("anssss", n, ans, memo)
	return ans
}
