package main

import (
	"fmt"
	"sort"
)

func main() {
	// func1()
	// func2()
	func3()
	// func4()
	fmt.Println("anssssssssssssssss end")
}
func func3() {
	maximumScore(2, 4, 6)
	maximumScore(1, 8, 8)
}
func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	l, m, r := arr[0], arr[1], arr[2]
	ans := l
	for l != 0 {
		if r > m {
			r--
		} else {
			m--
		}
		l--
	}
	fmt.Println("arr", arr)
	fmt.Println(l, m, r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans += min(m, r)
	fmt.Println("ansssss", ans)
	return ans
}
