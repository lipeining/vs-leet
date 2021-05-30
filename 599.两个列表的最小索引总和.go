/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	left, right := make(map[string]int), make(map[string]int)
	for i, str := range list1 {
		left[str] = i
	}
	for i, str := range list2 {
		right[str] = i
	}
	sum := 2001
	ans := make([]string, 0)
	for k, v := range left {
		if v2, ok := right[k]; ok {
			if v+v2 < sum {
				sum = v + v2
				ans = []string{k}
			} else if v+v2 == sum {
				ans = append(ans, k)
			}
		}
	}
	return ans
}

// @lc code=end

