/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	// 最开始使用前缀和求得对应的
	// sum := make([]int, 0)
	// for index,num := range nums {
	// 	if index == 0 {
	// 		sum = append(sum, num)
	// 	} else {
	// 		sum = append(sum, num+sum[index-1])
	// 	}
	// }
	// fmt.Println(sum)
	length := len(nums)
	sum := make([]int, length)
	for index,num := range nums {
		if index == 0 {
			sum[0] = num
		} else {
			sum[index] = num + sum[index-1]
		}
	}
	// fmt.Println(sum)
	max := float64(sum[k-1]) / float64(k)
	for i:=k;i<length;i++ {
		avg := float64(sum[i] - sum[i-k]) / float64(k)
		if avg > max {
			max = avg
		}
	}
	return max
}
// @lc code=end

