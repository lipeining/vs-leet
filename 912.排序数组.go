/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	partition := func(low, high int) int {
		privot := nums[low]
		for low < high {
			for low < high && nums[high] >= privot {
				high--
			}
			nums[low], nums[high] = nums[high], nums[low]
			for low < high && nums[low] <= privot {
				low++
			}
			nums[low], nums[high] = nums[high], nums[low]
		}
		return low
	}
	randomPartion := func(i, j int) int {
		t := rand.Int()%(j-i+1) + i
		nums[i], nums[t] = nums[t], nums[i]
		return partition(i, j)
	}
	var dfs func(left int, right int)
	dfs = func(left int, right int) {
		if left >= right {
			return
		}
		// p := partition(left, right)
		p := randomPartion(left, right)
		// fmt.Println("now p:", p, nums)
		dfs(left, p-1)
		dfs(p+1, right)
	}
	// fmt.Println("before ", nums)
	dfs(0, len(nums)-1)
	// fmt.Println("after ", nums)
	return nums
}

// @lc code=end

