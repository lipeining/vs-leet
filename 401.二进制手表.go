/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */
package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	readBinaryWatch(1)
// 	readBinaryWatch(3)
// }

// @lc code=start
func readBinaryWatch(num int) []string {
	hours := []int{1, 2, 4, 8, 0, 0, 0, 0, 0, 0}
	minutes := []int{0, 0, 0, 0, 1, 2, 4, 8, 16, 32}
	ans := make([]string, 0)
	var dfs func(num int, index int, hour, minute int)
	dfs = func(num int, index int, hour, minute int) {
		if hour >= 12 || minute >= 60 {
			return
		}
		if num == 0 {
			left := strconv.Itoa(hour)
			right := strconv.Itoa(minute)
			if minute < 10 {
				right = "0" + right
			}
			time := left + ":" + right
			ans = append(ans, time)
			return
		}
		for i := index; i < 10; i++ {
			dfs(num-1, i+1, hour+hours[i], minute+minutes[i])
		}
	}
	dfs(num, 0, 0, 0)
	// fmt.Println(ans)
	return ans
}
func readBinaryWatch2(num int) []string {
	ans := make([]string, 0)
	hourMap := map[int]int{0: 1, 1: 2, 2: 4, 3: 8}
	minuteMap := map[int]int{4: 1, 5: 2, 6: 4, 7: 8, 8: 16, 9: 32}
	toString := func(used []bool) string {
		hour := 0
		for i := 0; i < 4; i++ {
			if used[i] {
				hour += hourMap[i]
			}
		}
		minute := 0
		for i := 4; i < 10; i++ {
			if used[i] {
				minute += minuteMap[i]
			}
		}
		if hour >= 12 {
			return ""
		}
		if minute >= 60 {
			return ""
		}
		left := strconv.Itoa(hour)
		right := strconv.Itoa(minute)
		if minute < 10 {
			right = "0" + right
		}
		return left + ":" + right
	}
	var dfs func(cnt int, used []bool)
	dfs = func(cnt int, used []bool) {
		if cnt == num {
			time := toString(used)
			if time != "" {
				ans = append(ans, time)
			}
			return
		}
		for i := 0; i < 4+6; i++ {
			if !used[i] {
				used[i] = true
				dfs(cnt+1, used)
				used[i] = false
			}
		}
	}
	used := make([]bool, 4+6)
	dfs(0, used)
	fmt.Println("ansss", ans)
	return ans
}

// @lc code=end
