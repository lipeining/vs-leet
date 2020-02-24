import "sort"

/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] 救生艇
 */

// @lc code=start
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	// weight := 0
	// 1 2 4 5 limit=6
	left := 0
	length := len(people)
	right := length - 1
	for left <= right {
		ans++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}
	return ans
}

// @lc code=end
