/*
 * @lc app=leetcode.cn id=869 lang=golang
 *
 * [869] 重新排序得到 2 的幂
 */
package main

// func main() {
// 	reorderedPowerOf2(1289)
// }

// @lc code=start
func reorderedPowerOf2(n int) bool {
	check := func(num int) bool {
		cur := 1
		for k := 1; k <= 32; k++ {
			if num == cur {
				return true
			}
			cur <<= 1
		}
		return false
	}
	arr := make([]int, 0)
	it := n
	for it != 0 {
		m := it % 10
		arr = append(arr, m)
		it /= 10
	}
	// 对数组 arr 排列组合得到 num，检查是否可以得到答案
	ans := false
	var dfs func(first bool, path, cnt int, used []bool)
	dfs = func(first bool, path, cnt int, used []bool) {
		if cnt == len(arr) {
			// 已经全部数字使用到了
			power := check(path)
			if power {
				ans = true
			}
			return
		}
		if !first {
			for i := 0; i < len(arr); i++ {
				if arr[i] != 0 {
					used[i] = true
					dfs(true, path*10+arr[i], cnt+1, used)
					used[i] = false
				}
			}
		} else {
			for i := 0; i < len(arr); i++ {
				if !used[i] {
					used[i] = true
					dfs(true, path*10+arr[i], cnt+1, used)
					used[i] = false
				}
			}
		}
	}
	used := make([]bool, len(arr))
	dfs(false, 0, 0, used)
	return ans
}

// @lc code=end
