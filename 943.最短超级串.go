/*
 * @lc app=leetcode.cn id=943 lang=golang
 *
 * [943] 最短超级串
 */
package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	shortestSuperstring([]string{"alex", "loves", "leetcode"})
// 	shortestSuperstring([]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"})
// }

// @lc code=start
func shortestSuperstring(A []string) string {
	// 长度排序，
	// 状态压缩 dp
	n := len(A)
	// 官方解答
	overlaps := make([][]int, n)
	for i := 0; i < n; i++ {
		overlaps[i] = make([]int, n)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m := min(len(A[i]), len(A[j]))
			for k := m; k >= 0; k-- {
				if strings.HasSuffix(A[i], A[j][:k]) {
					overlaps[i][j] = k
					break
				}
			}
		}
	}
	// dp[mask][i] = 最短长度的，以i结尾
	m := 1 << n
	dp := make([][]int, m)
	parent := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		parent[i] = make([]int, n)
	}
	for mask := 0; mask < m; mask++ {
		for j := 0; j < n; j++ {
			parent[mask][j] = -1
		}
		for bit := 0; bit < n; bit++ {
			if (mask>>bit)&1 > 0 {
				// 如果 mask 的第 bit 位为 1
				pmask := mask ^ (1 << bit)
				if pmask == 0 {
					// mask 与  1 << bit 一样的话，
					// 说明目前只有一个字符串被挑中
					continue
				}
				for i := 0; i < n; i++ {
					if (pmask>>i)&1 > 0 {
						// 如果子 pmask end with i
						val := dp[pmask][i] + overlaps[i][bit]
						if val > dp[mask][bit] {
							dp[mask][bit] = val
							parent[mask][bit] = i
						}
					}
				}
			}
		}
	}
	perm := make([]int, n)
	seen := make([]bool, n)
	t := 0
	mask := m - 1
	p := 0
	// p 最后一个元素
	for j := 0; j < n; j++ {
		if dp[(1<<n)-1][j] > dp[(1<<n)-1][p] {
			p = j
		}
	}
	for p != -1 {
		perm[t] = p
		t++
		seen[p] = true
		p2 := parent[mask][p]
		mask ^= 1 << p
		p = p2
	}
	// reverse perm
	for i := 0; i < t/2; i++ {
		v := perm[i]
		perm[i] = perm[t-1-i]
		perm[t-1-i] = v
	}
	for i := 0; i < n; i++ {
		if !seen[i] {
			perm[t] = i
			t++
		}
	}
	ans := A[perm[0]]
	for i := 1; i < n; i++ {
		overlap := overlaps[perm[i-1]][perm[i]]
		ans += A[perm[i]][:overlap]
	}
	fmt.Println("ans", ans)
	return ans
}

// @lc code=end
