/*
 * @lc app=leetcode.cn id=1723 lang=golang
 *
 * [1723] 完成所有工作的最短时间
 */

// @lc code=start
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sum := make([]int, k)
	ans := math.MaxInt32
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(index int, used int, cur int)
	dfs = func(index int, used int, cur int) {
		if cur >= ans {
			return
		}
		// 优先分配给空闲的工人
		if index == n {
			if cur < ans {
				ans = cur
			}
			return
		}
		if used < k {
			sum[used] = jobs[index]
			dfs(index+1, used+1, max(sum[used], cur))
			sum[used] = 0
		}
		for i := 0; i < used; i++ {
			sum[i] += jobs[index]
			dfs(index+1, used, max(sum[i], cur))
			sum[i] -= jobs[index]
		}
	}
	dfs(0, 0, 0)
	return ans
}

// @lc code=end

