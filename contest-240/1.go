package main

import "fmt"

func main() {
	// maxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5})
	// maxDistance([]int{2, 2, 2}, []int{10, 10, 1})
	// maxDistance([]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25})
	// maxDistance([]int{5, 4}, []int{3, 2})
}
func maxSumMinProduct(nums []int) int {
	// 应该是
	n := len(nums)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n - 1
	}
	st := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(st) != 0 && nums[st[len(st)-1]] > nums[i] {
			r[st[len(st)-1]] = i - 1
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	st = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(st) != 0 && nums[st[len(st)-1]] > nums[i] {
			l[st[len(st)-1]] = i + 1
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// i 最小时，对应的子数组长度越大越好
	for i := 0; i < n; i++ {
		ans = max(ans, nums[i]*(s[r[i]+1]-s[l[i]]))
	}
	return ans % int(1e9+7)
}
func maxDistance(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	// 二分法
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	j := 0
	for i := 0; i < m; i++ {
		for j < n {
			if nums2[j] >= nums1[i] {
				j++
			} else {
				break
			}
		}
		ans = max(ans, j-i-1)
	}
	fmt.Println("ansssssss", ans)
	return ans
}
func maxDistanceOld(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	// 二分法
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	bin := func(target, l, r int) int {
		// 非递增，不断压缩右边界
		for l < r {
			// fmt.Println(l, r, target)
			mid := l + (r-l)/2
			if nums2[mid] > target {
				l = mid + 1
			} else if nums2[mid] < target {
				r = mid - 1
			} else {
				l = mid
			}
		}
		if nums2[l] >= target {
			return l
		}
		return l - 1
	}
	tryl := 0
	tryr := n - 1
	for i := 0; i < m; i++ {
		target := nums1[i]
		j := bin(target, tryl, tryr)
		if j >= 0 {
			fmt.Println("i-j", i, j)
			ans = max(ans, j-i)
		}
		tryl = j
		tryr = n - 1
	}
	// for i := m - 1; i >= 0; i-- {
	// 	target := nums1[i]
	// 	j := bin(target, tryl, tryr)
	// 	if j >= 0 {
	// 		fmt.Println("i-j", i, j)
	// 		ans = max(ans, j-i)
	// 	}
	// 	tryl = 0
	// 	tryr = j
	// }
	fmt.Println("ansssssss", ans)
	return ans
}
func maximumPopulation(logs [][]int) int {
	cnt := make(map[int]int)
	number := 0
	ans := 2051
	for _, log := range logs {
		b, d := log[0], log[1]
		for j := b; j < d; j++ {
			cnt[j]++
			if cnt[j] > number {
				number = cnt[j]
				ans = j
			} else if cnt[j] == number {
				if ans > j {
					ans = j
				}
			}
		}
	}
	return ans
}
