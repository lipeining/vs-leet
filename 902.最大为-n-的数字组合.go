/*
 * @lc app=leetcode.cn id=902 lang=golang
 *
 * [902] 最大为 N 的数字组合
 */
package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 100)
// 	atMostNGivenDigitSet([]string{"1", "4", "9"}, 1000000000)
// 	atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 777777)
// 	// 	["1","3","5","7"]
// 	// 777777
// 	// 5460
// 	atMostNGivenDigitSet([]string{"1", "3", "5", "7"}, 776477)
// 	// 	["1","3","5","7"]
// 	// 776477
// 	// 5396

// 	atMostNGivenDigitSet2([]string{"1", "3", "5", "7"}, 100)
// 	atMostNGivenDigitSet2([]string{"1", "4", "9"}, 1000000000)
// 	atMostNGivenDigitSet2([]string{"1", "3", "5", "7"}, 777777)
// 	// 	["1","3","5","7"]
// 	// 777777
// 	// 5460
// 	atMostNGivenDigitSet2([]string{"1", "3", "5", "7"}, 776477)
// 	// 	["1","3","5","7"]
// 	// 776477
// 	// 5396
// }

// @lc code=start
func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	length := len(s)
	dp := make([]int, length+1)
	dp[length] = 1
	pow := func(num int) int {
		number := 1
		for i := 1; i <= num; i++ {
			number *= len(digits)
		}
		return number
	}
	for i := length - 1; i >= 0; i-- {
		si := int(s[i] - '0')
		for _, d := range digits {
			numd, _ := strconv.Atoi(d)
			if numd < si {
				dp[i] += pow(length - i - 1)
			} else if numd == si {
				dp[i] += dp[i+1]
			}
		}
	}
	for i := 1; i < length; i++ {
		dp[0] += pow(i)
	}
	fmt.Println("ansssss", dp[0])
	return dp[0]
}
func atMostNGivenDigitSet2(digits []string, n int) int {
	length := 10
	dp := make([]int, length)
	digitMap := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	// i 个位数时，对应符合要求的组合数量
	check := func(num int) (less, equal int) {
		for _, digit := range digits {
			if digitMap[digit] < num {
				less++
			} else if digitMap[digit] == num {
				equal++
			}
		}
		return
	}
	count := func(num int) int {
		dLen := len(digits)
		cnt := 1
		for i := 1; i <= num; i++ {
			cnt *= dLen
		}
		return cnt
	}
	zero := n % 10
	l, e := check(zero)
	dp[0] = l + e
	fmt.Println(dp)
	for i := 1; i < length; i++ {
		n /= 10
		if n == 0 {
			break
		}
		mod := n % 10
		le, eq := check(mod)
		cnt := count(i)
		if eq == 1 {
			dp[i] += dp[i-1]
		} else {
			dp[i] += cnt
		}
		// dp[i] = eq*dp[i-1] + le*cnt
		fmt.Println("for i", i, dp[i], le, eq, cnt)
	}
	fmt.Println(dp)
	ans := 0
	for _, cnt := range dp {
		ans += cnt
	}
	fmt.Println("ansssssss", ans)
	return ans
}

// @lc code=end
