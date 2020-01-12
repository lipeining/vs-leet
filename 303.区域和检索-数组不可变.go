/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	nums []int
	prefixSum []int
}


func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums) + 1)
	for index,num := range nums {
		if index == 0 {
			prefixSum[index] = num
		} else {
			prefixSum[index] = prefixSum[index-1] + num
		}
	}
	return NumArray{nums, prefixSum}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.prefixSum[j]
	}
    return this.prefixSum[j]-this.prefixSum[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

