package main

import "strconv"

func main() {}
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	m := make(map[int]bool)
	twoToTen := func(str string) int {
		ans := 0
		mul := 1
		for i := n - 1; i >= 0; i-- {
			bit := str[i] == '1'
			if bit {
				ans += mul
			}
			mul *= 2
		}
		return ans
	}
	tenToTwo := func(num int) string {
		result := ""

		if n == 0 {
			return "0"
		}

		for ; num > 0; num /= 2 {
			lsb := num % 2
			result = strconv.Itoa(lsb) + result
		}

		return result
	}
	for i := 0; i < n; i++ {
		cov := twoToTen(nums[i])
		m[cov] = true
	}
	for i := 0; i <= n; i++ {
		if !m[i] {
			str := tenToTwo(i)
			for len(str) != n {
				str = "0" + str
			}
			return str
		}
	}
	return ""
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func findGCD(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	l, r := nums[0], nums[0]
	for _, num := range nums {
		l = min(l, num)
		r = max(r, num)
	}
	return gcd(l, r)
}
