package main

import (
	"fmt"
	"sort"
)

func main() {
	longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu")
	longestBeautifulSubstring("aeeeiiiioooauuuaeiou")
	longestBeautifulSubstring("a")
	// maxFrequency([]int{1, 2, 4}, 5)
	// maxFrequency([]int{3, 9, 6}, 2)
	// maxFrequency([]int{1, 4, 8, 13}, 5)
	// maxFrequency([]int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966}, 3056)
}
func longestBeautifulSubstring(word string) int {
	n := len(word)
	findStart := func(l int) int {
		for l < n {
			if word[l] != 'a' {
				l++
			} else {
				return l
			}
		}
		return l
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	l := findStart(0)
	r := l + 1
	ans := 0
	next := map[byte]byte{
		'a': 'e',
		'e': 'i',
		'i': 'o',
		'o': 'u',
	}
	w := byte('a')
	for r < n {
		for r < n && word[r] == w {
			r++
		}
		if r == n {
			break
		}
		fmt.Println("r", r, string(word[r]), string(w))
		if w == 'u' {
			fmt.Println("for now l, r", l, r)
			ans = max(ans, r-l)
			l = findStart(r)
			w = byte('a')
			r = l + 1
		} else {
			if word[r] != next[w] {
				l = findStart(r)
				w = byte('a')
				r = l + 1
			} else {
				w = next[w]
			}
		}
	}
	if w == 'u' {
		ans = max(ans, r-l)
	}
	fmt.Println("ansssss", ans)
	return ans
}
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(nums)
	ans := 0
	sum := 0
	l := 0
	r := 0
	for r < n {
		sum += nums[r]
		for sum+k < nums[r]*(r-l+1) {
			sum -= nums[l]
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}
	fmt.Println("ansssss", ans)
	return ans
}
