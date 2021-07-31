package main

import (
	"fmt"
	"strconv"
)

func main() {}
func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	n := len(students[0])
	sset := make([]bool, m)
	mset := make([]bool, m)
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cnt := func(i, j int) int {
		student := students[i]
		mentor := mentors[j]
		po := 0
		for k := 0; k < n; k++ {
			if student[k] == mentor[k] {
				po++
			}
		}
		return po
	}
	var dfs func(i int, point int)
	dfs = func(i int, point int) {
		if i == m {
			ans = max(ans, point)
			return
		}
		for index := i; index < m; index++ {
			// 第 index 个学生要找一个队
			for j := 0; j < m; j++ {
				if mset[j] {
					continue
				}
				sset[i] = true
				mset[j] = true
				dfs(i+1, point+cnt(i, j))
				sset[i] = false
				mset[j] = false
			}
		}
	}
	dfs(0, 0)
	return ans
}

func maximumNumber(num string, change []int) string {
	ans := ""
	n := len(num)
	i := 0
	for i < n {
		now := int(num[i] - '0')
		c := change[now]
		char := string(num[i])
		if c > now {
			char = strconv.Itoa(c)
			ans += char
			break
		} else {
			ans += char
			i++
		}
	}
	j := i + 1
	for j < n {
		now := int(num[j] - '0')
		c := change[now]
		if c > now {
			ans += strconv.Itoa(c)
			j++
		} else {
			break
		}
	}
	if j < n {
		ans += num[j:]
	}
	return ans
}
func getLucky(s string, k int) int {
	number := ""
	for i := 0; i < len(s); i++ {
		number += strconv.Itoa(int(s[i]-'a') + 1)
	}
	fmt.Println("number", number)
	turn := func(num string) int {
		ans := 0
		for i := 0; i < len(num); i++ {
			ans += int(num[i] - '0')
		}
		return ans
	}
	for step := 0; step < k; step++ {
		t := turn(number)
		if step == k-1 {
			return t
		}
		number = strconv.Itoa(t)
	}
	return 0
}
