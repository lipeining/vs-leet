/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */
import "sort"
// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// nums1 中数字 x 的下一个更大元素是指
	//  x 在 nums2 中对应位置的右边的第一个比 x 大的元素
	// 不可以排序
	// sort.Ints(nums2)
	// smap := make(map[int]int)
	// for i,n := range nums2 {
	// 	smap[n] = i
	// }
	// fmt.Println(smap)
	// fmt.Println(nums2)
	// res := make([]int,0)
	// length := len(nums2)
	// for _,n := range nums1 {
	// 	t := smap[n]
	// 	tmp := -1
	// 	fmt.Println(t, length-1)
	// 	if t != length - 1 {
	// 		tmp = nums2[t+1]
	// 	}
	// 	res = append(res, tmp)
	// }
	// return res
	// 暴力的 O(n^2)
	smap := make(map[int]int)
	for i,n := range nums2 {
		smap[n] = i
	}
	res := make([]int,0)
	length := len(nums2)
	for _,n := range nums1 {
		tmp := -1
		for j:=smap[n]+1;j<length;j++ {
			if nums2[j] > n {
				tmp = nums2[j]
				break
			}
		}
		res = append(res, tmp)
	}
	return res
}
// @lc code=end

