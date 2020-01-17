/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	grid := make(map[int][]int)
	// arr : len start end
	for i:=0;i<len(nums);i++ {
		n := nums[i]
		arr,ok:=grid[n]
		if ok {
			arr[0]++
			arr[2]=i
			grid[n]=arr
		} else {
			tmp := make([]int,3)
			tmp[0]=1
			tmp[1]=i
			tmp[2]=i
			grid[n]=tmp
		}
	}
	max := 0
	ans := len(nums)
	fmt.Println(grid)
	for _,arr := range grid {
		if arr[0] == max {
			max = arr[0]
			len := arr[2]-arr[1] + 1
			if len < ans {
				ans = len
			}
		} else if arr[0] > max {
			max = arr[0]
			len := arr[2]-arr[1] + 1
			ans = len
		}
	}
	return ans
}
// @lc code=end

