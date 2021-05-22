/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	rand.Seed(time.Now().UnixNano())
	partition := func(i, j int) int {
		privot := nums[i]
		for i < j {
			for i < j && nums[j] >= privot {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
			for i < j && nums[i] <= privot {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		return i
	}
	randomPartion := func(i, j int) int {
		t := rand.Int()%(j-i+1) + i
		nums[i], nums[t] = nums[t], nums[i]
		return partition(i, j)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i > j {
			return -1
		}
		p := randomPartion(i, j)
		if p == n-k {
			return nums[p]
		} else if p < n-k {
			return dfs(p+1, j)
		} else {
			return dfs(i, p-1)
		}
	}
	return dfs(0, n-1)
	// sort.Ints(nums)
	// return nums[len(nums)-k]
}

// @lc code=end

