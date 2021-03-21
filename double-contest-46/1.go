package main

import (
	"fmt"
	"unicode"
)

func main() {
	// func2()
	// func3()
	// highestPeak([][]int{
	// 	{0, 1},
	// 	{0, 0},
	// })
	// highestPeak([][]int{
	// 	{0, 0, 1},
	// 	{1, 0, 0},
	// 	{0, 0, 0},
	// })
	// canChoose([][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0})
	// canChoose([][]int{{10, -2}, {1, 2, 3, 4}}, []int{1, 2, 3, 4, 10, -2})
	// canChoose([][]int{{1, 2, 3}, {3, 4}}, []int{7, 7, 1, 2, 3, 4, 7, 7})
	func4()
}
func func4() {
}
func getCoprimes(nums []int, edges [][]int) []int {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	n := len(nums)
	ans := make([]int, n)
	parent := make(map[int]int)
	for _, edge := range edges {
		parent[edge[1]] = edge[0]
	}
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	for i := 1; i < n; i++ {
		node := i
		for {
			p := parent[node]
			if gcd(nums[p], nums[i]) == 1 {
				ans[i] = p
				break
			}
			if p == 0 {
				break
			}
			node = p
		}
	}
	return ans
}
func canChoose(groups [][]int, nums []int) bool {
	// 将 groups 的元素能够在 nums 中一一找出就可以了，但是找出来时，不能相交。
	n := len(groups)
	length := len(nums)
	j := 0
	check := func(x int, b []int) bool {
		for i := 0; i < len(groups[x]); i++ {
			if groups[x][i] != b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		need := len(groups[i])
		fix := false
		for j+need <= length {
			if check(i, nums[j:j+need]) {
				fmt.Println("on ", j, "fix")
				j = j + need
				fix = true
				break
			}
			j++
		}
		if !fix {
			fmt.Println("fail on", i, j)
			return false
		}
	}
	fmt.Println("pass")
	return true
}
func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])
	queue := make([]int, 0)
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
			if isWater[i][j] == 1 {
				queue = append(queue, i*m+j)
				ans[i][j] = 0
			}
		}
	}
	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			x := queue[i] / m
			y := queue[i] % m
			fmt.Println("on", x, "-", y, "value", ans[x][y])
			for _, d := range dir {
				nx := x + d[0]
				ny := y + d[1]
				if nx < 0 || ny < 0 || nx >= m || ny >= n {
					continue
				}
				if ans[nx][ny] != -1 {
					continue
				}
				ans[nx][ny] = ans[x][y] + 1
				fmt.Println("nx-ny", nx, "-", ny, "to be", ans[nx][ny])
				queue = append(queue, nx*m+ny)
			}
		}
		queue = queue[size:]
	}
	fmt.Println(ans)
	return ans
}
func longestNiceSubstring(s string) string {
	n := len(s)
	ans := ""
	check := func(str string) bool {
		m := make(map[byte]bool)
		for i := 0; i < len(str); i++ {
			m[str[i]] = true
		}
		for k, _ := range m {
			if unicode.IsLower(rune(k)) {
				_, ok := m[byte(unicode.ToUpper(rune(k)))]
				if !ok {
					return false
				}
			} else {
				_, ok := m[byte(unicode.ToLower(rune(k)))]
				if !ok {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			str := s[i:j]
			if check(str) {
				if len(str) > len(ans) {
					ans = str
				}
			}
		}
	}
	return ans
}
