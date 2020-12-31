import "sort"

/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	// 抄答案
	row := len(nums) - 1
	if nums[row] > target {
		return false
	}
	n := len(nums)
	all := 1<<n - 1
	dp := make([]bool, all+1)
	dp[0] = true
	curSum := make([]int, all+1)
	for i := 0; i <= all; i++ {
		if !dp[i] {
			continue
		}
		for j := 0; j < n; j++ {
			// 状态 I 可以跟 j 转化
			// 第 j 位 1<<j 如果在状态 i 中为1的话，说明，已经处理了这个状态
			if i&(1<<j) != 0 {
				continue
			}
			nxt := i | (1 << j)
			// 已经计算过 dp[nxt]
			if dp[nxt] {
				continue
			}
			if nums[j]+curSum[i]%target <= target {
				dp[nxt] = true
				curSum[nxt] = curSum[i] + nums[j]
			} else {
				// j升序，可以剪枝
				break
			}
		}
	}
	return dp[all]
}
func canPartitionKSubsetsDFS(nums []int, k int) bool {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	// 抄答案
	row := len(nums) - 1
	if nums[row] > target {
		return false
	}
	for row >= 0 && nums[row] == target {
		row--
		k--
	}
	groups := make([]int, k)
	return search(groups, row, nums, target)
}
func search(groups []int, row int, nums []int, target int) bool {
	if row < 0 {
		return true
	}
	v := nums[row]
	row--
	for i := 0; i < len(groups); i++ {
		if groups[i]+v <= target {
			groups[i] += v
			if search(groups, row, nums, target) {
				return true
			}
			groups[i] -= v
		}
		if groups[i] == 0 {
			break
		}
	}
	return false
}

// @lc code=end
