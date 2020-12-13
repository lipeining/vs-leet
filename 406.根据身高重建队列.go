/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || people[i][0] == people[j][0] && people[i][1] > people[j][1]
	})
	n := len(people)
	ans := make([][]int, n)
	for _, person := range people {
		pos := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				pos--
				if pos == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}

// @lc code=end

