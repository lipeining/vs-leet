/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */
import "sort"
// @lc code=start
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
	sort.Ints(s)
	res := 0
	used := make([]int, len(s))
	// 可以优化这个 used 数组为一个 s 的下标变化
	for _,gi :=range g {
		for j,sj := range s {
			if used[j] == 1 {
				continue
			} else {
				if sj >= gi {
					used[j]=1
					res++
					break
				}
			}
		}
	}
	fmt.Println(used)
	return res
}
// @lc code=end

