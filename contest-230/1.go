package main

import (
	"math"
	"sort"
)

func main() {
	minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2})
	minOperations([]int{1, 1, 1, 1, 1, 1, 1}, []int{6})
	minOperations([]int{6, 6}, []int{1})
	// 	[5,2,1,5,2,2,2,2,4,3,3,5]
	// [1,4,5,5,6,3,1,3,3]
	// 1
	minOperations([]int{5, 2, 1, 5, 2, 2, 2, 2, 4, 3, 3, 5}, []int{1, 4, 5, 5, 6, 3, 1, 3, 3})
}
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	// 有限背包问题吧
	// 假设好 baseCost 之后，在 toppingCost 中每一个中最多两次的得到 target-baseCost 的方案
	return 0
}
func minOperations(nums1 []int, nums2 []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	sum1, sum2 := 0, 0
	for i := 0; i < len(nums1); i++ {
		sum1 += nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		sum2 += nums2[i]
	}
	if sum1 == sum2 {
		return 0
	}
	if sum1 < sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}
	sort.Slice(nums1, func(a, b int) bool {
		return nums1[a] > nums1[b]
	})
	sort.Slice(nums2, func(a, b int) bool {
		return nums2[a] < nums2[b]
	})
	// fmt.Println(nums1, nums2)
	// 保证 1 大于 2
	need := sum1 - sum2
	// 有可能使用两个数，三个数变小的方式，那么需要提前排序才可以啊
	needLeft := func(left int) int {
		lcnt := 0
		for i := 0; i < len(nums1); i++ {
			if left <= 0 {
				break
			}
			if nums1[i] > 1 {
				left -= min(left, abs(1-nums1[i]))
				lcnt++
			}
		}
		if left > 0 {
			return -1
		}
		return lcnt
	}
	needRight := func(right int) int {
		rcnt := 0
		for i := 0; i < len(nums2); i++ {
			if right <= 0 {
				break
			}
			if nums2[i] < 6 {
				right -= min(right, abs(6-nums2[i]))
				rcnt++
			}
		}
		if right > 0 {
			return -1
		}
		return rcnt
	}
	ans := math.MaxInt32
	for pos := need; pos >= 0; pos-- {
		left := pos
		right := need - pos
		lcnt := needLeft(left)
		rcnt := needRight(right)
		if lcnt == -1 || rcnt == -1 {
			continue
		}
		ans = min(ans, lcnt+rcnt)
		// fmt.Println("anssss", ans, lcnt, left, rcnt, right)
	}
	// fmt.Println("anssss", ans)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
