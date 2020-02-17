/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H指数
 */

// @lc code=start
func hIndex(citations []int) int {
	length := len(citations)
	if length == 0 {
		return 0
	}
	// [100]: 1 [11, 15]: 2
	sort.Ints(citations)
	// ans := 0
	// if citations[length-1] > 0 {
	// 	ans = 1
	// }
	// for i:=length-1;i>=0;i-- {
	// 	has := length-i-1
	// 	if has >= citations[i] {
	// 		if ans < citations[i] {
	// 			ans = citations[i]
	// 		}
	// 	}
	// }
	// return ans
	fmt.Println(citations)
	i := 0
	for i < length {
		if citations[length-i-1] < i {
			break
		}
		i++
	}
	return i
}
// @lc code=end

