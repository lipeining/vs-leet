package main

import (
	"strconv"
)

func main() {

}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func interchangeableRectangles(rectangles [][]int) int64 {
	n := len(rectangles)
	m := make(map[string]int)
	toIndex := func(up, down int) string {
		return strconv.Itoa(up) + "/" + strconv.Itoa(down)
	}
	for i := 0; i < n; i++ {
		up := rectangles[i][0]
		down := rectangles[i][1]
		gg := gcd(up, down)
		up /= gg
		down /= gg
		index := toIndex(up, down)
		m[index]++
	}
	ans := int64(0)
	for _, v := range m {
		if v >= 2 {
			ans += int64(v * (v - 1) / 2)
		}
	}
	return ans
}

func reversePrefix(word string, ch byte) string {
	n := len(word)
	j := 0
	for j < n {
		if word[j] == ch {
			break
		}
		j++
	}
	if j == n {
		return word
	}
	ans := ""
	for i := j; i >= 0; i-- {
		ans += string(word[i])
	}
	ans += word[j+1 : n]
	return ans
}
