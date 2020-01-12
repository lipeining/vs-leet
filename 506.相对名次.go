/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */
import "sort"
// @lc code=start
func findRelativeRanks(nums []int) []string {
	smap := make(map[int]int)
	for i,n := range nums {
		smap[n] = i
	}
	sort.Ints(nums)
	length := len(nums)
	res := make([]string, length)
	for i:= 0;i<length;i++ {
		p := length - i
		originIndex := smap[nums[i]]
		str := strconv.Itoa(p)
		if p == 1 {
			str = "Gold Medal"
		} else if p == 2 {
			str = "Silver Medal"
		} else if p == 3 {
			str = "Bronze Medal"
		} else {
			// do nothing
		}
		res[originIndex] = str
		// fmt.Println(originIndex, str)
	}
	return res
}
// @lc code=end

