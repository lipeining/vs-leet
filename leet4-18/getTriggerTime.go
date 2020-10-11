package main

import "fmt"

func main() {
	// 	[[2,8,4],[2,5,0],[10,9,8]]
	// [[2,11,3],[15,10,7],[9,17,12],[8,1,14]]
	// increase := [][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}}
	// requirements := [][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}}
	// ans := getTriggerTime(increase, requirements)
	// fmt.Print(ans)
	increase := [][]int{{1,1,1}}
	requirements := [][]int{{0,0,0}}
	ans := getTriggerTime(increase, requirements)
	fmt.Print(ans)
}
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	// 排序，使用二分查找方式得到对应的值
	//  time 是单调递增的一个 CRH 数组
	// 对于每一个 requirements ，二分查找 time 中符合条件的最左元素?
	// fmt.Println(increase)
	// fmt.Println(requirements)
	days := len(increase)
	time := make([][]int, days+1)
	time[0] = []int{0, 0, 0}
	for i := 1; i <= days; i++ {
		time[i] = make([]int, 3)
		time[i][0] = time[i-1][0] + increase[i-1][0]
		time[i][1] = time[i-1][1] + increase[i-1][1]
		time[i][2] = time[i-1][2] + increase[i-1][2]
	}
	n := len(requirements)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	// fmt.Println(time)
	for j := 0; j < len(requirements); j++ {
		// 二分查找符合条件的最左元素
		wantC, wantR, wantH := requirements[j][0], requirements[j][1], requirements[j][2]
		left, right := 0, days
		for left < right {
			mid := left + (right-left)/2
			C, R, H := time[mid][0], time[mid][1], time[mid][2]
			// fmt.Println("want:", wantC, wantR, wantH)
			// fmt.Println("now:", C, R, H)
			// fmt.Println(left, mid, right)
			if C >= wantC && R >= wantR && H >= wantH {
				right = mid
			} else {
				// 向右
				left = mid + 1
			}
		}
		C, R, H := time[left][0], time[left][1], time[left][2]
		if C >= wantC && R >= wantR && H >= wantH {
			ans[j] = left
		}
	}
	return ans
}

// func getTriggerTime(increase [][]int, requirements [][]int) []int {
// 	// 排序，使用二分查找方式得到对应的值
// 	//  time 是单调递增的一个 CRH 数组
// 	// 对于每一个 requirements ，二分查找 time 中符合条件的最左元素?

// 	n := len(requirements)
// 	ans := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		ans[i] = -1
// 	}
// 	for i := 0; i <= len(increase); i++ {
// 		// 找到可以触发的地方
// 		// 启用减法
// 		for j := 0; j < len(requirements); j++ {
// 			if ans[j] == -1 {
// 				if requirements[j][0] <= 0 && requirements[j][1] <= 0 && requirements[j][2] <= 0 {
// 					ans[j] = i
// 				}
// 				if i < len(increase) {
// 					C, R, H := increase[i][0], increase[i][1], increase[i][2]
// 					requirements[j][0] -= C
// 					requirements[j][1] -= R
// 					requirements[j][2] -= H
// 				}
// 			}
// 		}
// 	}
// 	return ans
// }

// func getTriggerTime(increase [][]int, requirements [][]int) []int {
// 	C,R,H := 0,0,0
// 	n := len(increase)
// 	ans := make([]int, n)
// 	for i:=0;i<len(increase);i++ {
// 		C += increase[i][0]
// 		R += increase[i][1]
// 		H += increase[i][2]
// 		// 找到可以触发的地方
// 		index := make([]int, 0)
// 		for j,r := range requirements {
// 			if C >= r[0] && R >= r[1] && H >= r[2] && ans[j] == 0 {
// 				index = append(index, j)
// 			}
// 		}
// 		for _,j := range index {
// 			ans[j] = i
// 		}
// 	}
// 	return ans
// }
