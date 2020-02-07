package template

import (
	"fmt"
	"sort"
)

// 39 40 46 47 77 78 90
// 39
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) == 0 {
		return ans
	}
	// 不使用集合的话，如何处理呢？
	// 参考题解，使用下标作为筛选条件
	var dfs func(candidates []int, target int, sum int, index int, path []int)
	dfs = func(candidates []int, target int, sum int, index int, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			fmt.Println(ans, path)
			ans = append(ans, path)
			return
		}
		for i := index; i < len(candidates); i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, candidates[i])
			dfs(candidates, target, sum+candidates[i], i, tmp)
		}
	}
	path := make([]int, 0)
	dfs(candidates, target, 0, 0, path)
	return ans
}

// 40
func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) == 0 {
		return ans
	}
	// 不使用集合的话，如何处理呢？
	// 参考题解，使用下标作为筛选条件
	var dfs func(candidates []int, target int, sum int, index int, used []bool, path []int)
	dfs = func(candidates []int, target int, sum int, index int, used []bool, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			fmt.Println(ans, path)
			ans = append(ans, path)
			return
		}
		for i := index; i < len(candidates); i++ {
			if used[i] {
				continue
			}
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, candidates[i])
			used[i] = true
			dfs(candidates, target, sum+candidates[i], i+1, used, tmp)
			used[i] = false
		}
	}
	sort.Ints(candidates)
	path := make([]int, 0)
	used := make([]bool, len(candidates))
	dfs(candidates, target, 0, 0, used, path)
	return ans
}

// 46
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, used []bool, path []int)
	dfs = func(nums []int, used []bool, path []int) {
		if len(path) == len(nums) {
			ans = append(ans, path)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			used[i] = true
			dfs(nums, used, tmp)
			used[i] = false
		}
	}
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, path)
	return ans
}

// 47
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, used []bool, path []int)
	dfs = func(nums []int, used []bool, path []int) {

		if len(path) == len(nums) {
			ans = append(ans, path)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			used[i] = true
			dfs(nums, used, tmp)
			used[i] = false
		}
	}
	sort.Ints(nums)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, path)
	return ans
}

// 77
func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, k int, index int, path []int)
	dfs = func(nums []int, k int, index int, path []int) {
		if len(path) == k {
			ans = append(ans, path)
			return
		}
		for i := index; i < len(nums); i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			dfs(nums, k, i+1, tmp)
		}
	}
	path := make([]int, 0)
	dfs(nums, k, 0, path)
	return ans
	// var dfs func(nums []int, k int, index int, used []bool, path []int)
	// dfs = func(nums []int, k int, index int,used []bool, path []int) {
	// 	if len(path) == k {
	// 		ans = append(ans, path)
	// 		return
	// 	}
	// 	for i:=index;i<len(nums);i++ {
	// 		tmp := make([]int, len(path))
	// 		copy(tmp, path)
	// 		tmp = append(tmp, nums[i])
	// 		used[i]=true
	// 		dfs(nums, k, i+1, used, tmp)
	// 		used[i]=false
	// 	}
	// }
	// used := make([]bool, len(nums))
	// path := make([]int, 0)
	// dfs(nums, k, 0, used, path)
	// return ans
}

// 78
func subsets(nums []int) [][]int {
	// nums := make([]int, 0)
	// for i:=1;i<=n;i++ {
	// 	nums = append(nums, i)
	// }
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, k int, index int, path []int)
	dfs = func(nums []int, k int, index int, path []int) {
		if len(path) == k {
			ans = append(ans, path)
			return
		}
		for i := index; i < len(nums); i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			dfs(nums, k, i+1, tmp)
		}
	}
	path := make([]int, 0)
	for i := 0; i <= len(nums); i++ {
		dfs(nums, i, 0, path)
	}
	return ans
}

// 90
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, k int, index int, path []int)
	dfs = func(nums []int, k int, index int, path []int) {
		if len(path) == k {
			ans = append(ans, path)
			return
		}
		for i := index; i < len(nums); i++ {
			if i > index && nums[i-1] == nums[i] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			dfs(nums, k, i+1, tmp)
		}
	}
	path := make([]int, 0)
	sort.Ints(nums)
	for i := 0; i <= len(nums); i++ {
		dfs(nums, i, 0, path)
	}
	return ans
}
