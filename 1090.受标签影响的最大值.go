import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1090 lang=golang
 *
 * [1090] 受标签影响的最大值
 */

// @lc code=start
type pair struct {
	value int
	label int
}

// 抄题解思路，使用贪心算法
func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	pairs := make([]pair, len(values))
	for i := 0; i < len(values); i++ {
		pairs[i] = pair{value: values[i], label: labels[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})
	// fmt.Println(pairs[0])
	ans := 0
	used := make(map[int]int)
	wanted := 0
	for i := 0; i < len(pairs); i++ {
		if wanted >= num_wanted {
			break
		}
		p := pairs[i]
		now, ok := used[p.label]
		if !ok {
			ans += p.value
			used[p.label] = 1
			wanted++
		} else {
			if now >= use_limit {
				continue
			} else {
				ans += p.value
				used[p.label] = now + 1
				wanted++
			}
		}
	}
	return ans
}

// @lc code=end
