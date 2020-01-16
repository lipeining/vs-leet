/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	// 原地交换算法
	// 或者使用排序的算法
	for index,num := range nums {
		toSwap := num
		for toSwap!=index+1 && nums[toSwap-1] != toSwap {
			tmp := nums[index]
			nums[index] = nums[toSwap-1]
			nums[toSwap-1] = tmp

			// 
			toSwap = nums[index]
		}
	}
	fmt.Println(nums)
	res := make([]int, 0)
	for index,num := range nums {
		if num != index + 1 {
			res = append(res, num)
			res = append(res, index + 1)
			break
		}
	}
	return res
}
// @lc code=end

