/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */
package main

import "fmt"

// func main() {
// 	canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
// 	canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
// 	canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2)
// }

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	prev := false
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			prev = true
		} else {
			if prev {
				prev = false
			} else {
				if i == length-1 || (i+1 < length && flowerbed[i+1] == 0) {
					prev = true
					n--
					// fmt.Println("place on i", i)
				}
			}
		}
	}
	fmt.Println("ans", n <= 0)
	return n <= 0
}
func canPlaceFlowersOld(flowerbed []int, n int) bool {
	res := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] != 1 {
			l := i - 1
			r := i + 1
			if i == 0 {
				l = i
			}
			if i == len(flowerbed)-1 {
				r = i
			}
			if flowerbed[l] == 0 && flowerbed[r] == 0 {
				res++
				flowerbed[i] = 1
			}
		}
	}
	fmt.Println(flowerbed, res)
	return res >= n
}

// @lc code=end
