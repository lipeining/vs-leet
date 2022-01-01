package main

import "fmt"

func main() {
	// minTimeToType("bza")
	// minTimeToType("zjpc")
}
func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)
	m := len(matrix[0])
	neg := 0
	sum := int64(0)
	target := 100000 + 1
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] < 0 {
				neg++
			}
			a := abs(matrix[i][j])
			if a < target {
				target = a
			}
			sum += int64(a)
		}
	}
	if neg%2 == 0 {
		return sum
	}
	return sum - int64(target)*2
}
func minTimeToType(word string) int {
	n := len(word)
	ans := 0
	last := byte('a')
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		if word[i] == last {
			// do nothing
		} else if word[i] < last {
			ans += min(int(last-word[i]), (int(word[i])-int(last)+26)%26)
		} else {
			ans += min(int(word[i]-last), (int(last)-int(word[i])+26)%26)
		}
		fmt.Println(ans, last, string(last), word[i], string(word[i]))
		ans++
		last = byte(word[i])
	}
	fmt.Println("anssss", ans)
	return ans
}
