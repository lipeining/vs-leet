/*
 * @lc app=leetcode.cn id=552 lang=golang
 *
 * [552] 学生出勤记录 II
 */
package main

import "fmt"

// func main() {
// 	// checkRecord(1)
// 	// checkRecord(2)
// 	checkRecord(3) // 19
// }

// @lc code=start
func checkRecord(n int) int {
	toMod := int(1e9 + 7)
	a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 := 1, 0, 0, 0, 0, 0
	for i := 0; i <= n; i++ {
		a0l0Next := (a0l0 + a0l1 + a0l2) % toMod
		a0l2 = a0l1
		a0l1 = a0l0
		a0l0 = a0l0Next
		a1l0Next := (a0l0 + a1l0 + a1l1 + a1l2) % toMod
		a1l2 = a1l1
		a1l1 = a1l0
		a1l0 = a1l0Next
	}
	fmt.Println("ans", a1l0)
	return a1l0
}
func checkRecordNANL(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 3
	}
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = make([]int, 3)
		}
	}
	ans := 0
	toMod := int(1e9 + 7)
	// 前 i 个长度的字符串中 A,L,P 结尾的连续长度为 0,1,2
	// dp[0][0][0]
	// 第二位 分别为 A:0, L:1, P:2
	dp[1][0][1] = 1
	dp[1][1][1] = 1
	dp[1][2][1] = 1
	// 理解错误题意了，整个字符串中有两次 A 就不及格了
	for i := 2; i <= n; i++ {
		// 0 to ignore
		// dp[i][0][0]
		// dp[i][1][0]
		// dp[i][2][0]
		dp[i][0][1] = (dp[i-1][1][1] + dp[i-1][1][2] + dp[i-1][2][1] + dp[i-1][2][2]) % toMod
		// dp[i][0][2] 应该为 0
		dp[i][1][1] = (dp[i-1][0][1] + dp[i-1][2][1] + dp[i-1][2][2]) % toMod
		dp[i][1][2] = dp[i-1][1][1] % toMod

		dp[i][2][1] = (dp[i-1][0][1] + dp[i-1][1][1] + dp[i-1][1][2]) % toMod
		dp[i][2][2] = (dp[i-1][2][1] + dp[i-1][2][2]) % toMod
		fmt.Println("iiiiiiiiiiiiii", i)
		fmt.Println("A", dp[i][0])
		fmt.Println("L", dp[i][1])
		fmt.Println("P", dp[i][2])
	}
	for char := 0; char < 3; char++ {
		for k := 0; k < 3; k++ {
			ans += dp[n][char][k]
			ans %= toMod
		}
	}
	fmt.Println("ans---------------------", ans)
	return ans
}

// @lc code=end
