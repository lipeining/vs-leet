package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// memLeak(2, 2)
	// memLeak(9, 6)
	// memLeak(8, 11)
	printFunc := func(box [][]byte) {
		m := len(box)
		n := len(box[0])
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Print("\"" + string(box[i][j]) + "\"" + ",")
			}
			fmt.Println("")
		}
		fmt.Println("eneneenenen")
	}
	printFunc(rotateTheBox([][]byte{{'#', '.', '#'}}))
	printFunc(rotateTheBox([][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}))
	printFunc(rotateTheBox([][]byte{
		{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'},
	}))
}
func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, m)
	}
	// 从 ans 的最底层开始填充
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			row := box[m-j-1]
			y := i
			now := row[y]
			stone := false
			// 默认为 原来位置的值，向左查找石头或者障碍
			for y >= 0 {
				if row[y] == '#' {
					stone = true
					row[y] = '.'
					break
				} else if row[y] == '*' {
					break
				}
				y--
			}
			if stone {
				now = byte('#')
			}
			ans[i][j] = now
		}
	}
	return ans
}
func memLeak(memory1 int, memory2 int) []int {
	l := 1
	r := memory1 + memory2
	sum := memory1 + memory2
	check := func(num int) bool {
		return num*(num+1)/2 > sum
	}
	for l < r {
		mid := l + (r-l)/2
		// fmt.Println(l, r)
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println("l,r", l, r)
	i := 1
	for i < l {
		if memory1 >= memory2 {
			if memory1 < i {
				break
			}
			memory1 -= i
		} else {
			if memory2 < i {
				break
			}
			memory2 -= i
		}
		i++
	}
	ans := []int{i, memory1, memory2}
	fmt.Println("anssss", ans)
	return ans
}
func sortSentence(s string) string {
	list := strings.Split(s, " ")
	toIndex := func(input string) int {
		numKey := string(input[len(input)-1])
		num, _ := strconv.Atoi(numKey)
		return num
	}
	sort.Slice(list, func(i, j int) bool {
		return toIndex(list[i]) < toIndex(list[j])
	})
	ans := ""
	for i := 0; i < len(list); i++ {
		ans += list[i][:len(list[i])-1]
		if i != len(list)-1 {
			ans += " "
		}
	}
	return ans
}
