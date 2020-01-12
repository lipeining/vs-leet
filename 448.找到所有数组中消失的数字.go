/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	// 原地交换
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
			res = append(res, index + 1)
		}
	}
	return res
}
// @lc code=end

