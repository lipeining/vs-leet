import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=826 lang=golang
 *
 * [826] 安排工作以达到最大收益
 */

// @lc code=start

// func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

// }

type Node struct {
	d int
	p int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	ans := 0
	list := make([]*Node, len(difficulty))
	for i := 0; i < len(difficulty); i++ {
		list[i] = &Node{d: difficulty[i], p: profit[i]}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].d < list[j].d
	})
	// fmt.Println(list[0].d, list[1].d)
	sort.Ints(worker)
	i := 0
	now := 0
	for _, w := range worker {
		for i < len(difficulty) && w >= list[i].d {
			now = max(now, list[i].p)
			i++
		}
		ans += now
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
