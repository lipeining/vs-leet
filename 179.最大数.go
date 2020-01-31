/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i:=0;i<len(nums);i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
    sort.Slice(strs, func(i int, j int)bool {
		left,_ := strconv.Atoi(strs[i]+strs[j])
		right,_ := strconv.Atoi(strs[j]+strs[i])
		return left > right	
	})
	// fmt.Println(strs)
	ans := ""
	for i:=0;i<len(strs);i++ {
		ans += strs[i]
	}
	zero,_ := strconv.Atoi(ans)
	if zero == 0 {
		return "0"
	}
	return ans
}
// @lc code=end

