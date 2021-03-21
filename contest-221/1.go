package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("ans-------------------")
	// 	输入：s = "book"
	// 输出：true
	// 解释：a = "bo" 且 b = "ok" 。a 中有 1 个元音，b 也有 1 个元音。所以，a 和 b 相似。
	// 示例 2：

	// 输入：s = "textbook"
	// 输出：false
	// 解释：a = "text" 且 b = "book" 。a 中有 1 个元音，b 中有 2 个元音。因此，a 和 b 不相似。
	// 注意，元音 o 在 b 中出现两次，记为 2 个。
	// 示例 3：

	// 输入：s = "MerryChristmas"
	// 输出：false
	// 示例 4：

	// 输入：s = "AbCdEfGh"
	// 输出：true
	// fmt.Println(halvesAreAlike("book"))
	// fmt.Println(halvesAreAlike("textbook"))
	// fmt.Println(halvesAreAlike("MerryChristmas"))
	// fmt.Println(halvesAreAlike("AbCdEfGh"))
	func2()
	// func3()
	// func4()
}
func func4() {

	// pre := make([]int, n+1)
	// 最开始的是0
	// 到时，取值的时候，对0进行一次异或即可
	// for i := 0; i < n; i++ {
	// 	pre[i+1] = pre[i] ^ nums[i]
	// }
	// fmt.Println(pre)

	// 	输入：nums = [0,1,2,3,4], queries = [[3,1],[1,3],[5,6]]
	// 输出：[3,3,7]
	// 解释：
	// 1) 0 和 1 是仅有的两个不超过 1 的整数。0 XOR 3 = 3 而 1 XOR 3 = 2 。二者中的更大值是 3 。
	// 2) 1 XOR 2 = 3.
	// 3) 5 XOR 2 = 7.
	// 示例 2：

	// 输入：nums = [5,2,4,6,6,3], queries = [[12,4],[8,1],[6,3]]
	// 输出：[15,-1,5]
	maximizeXor([]int{0, 1, 2, 3, 4}, [][]int{{3, 1}, {1, 3}, {5, 6}})
	maximizeXor([]int{5, 2, 4, 6, 6, 3}, [][]int{{12, 4}, {8, 1}, {6, 3}})
}
func maximizeXor(nums []int, queries [][]int) []int {
	n := len(nums)
	sort.Ints(nums)
	m := len(queries)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = -1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < m; i++ {
		xi := queries[i][0]
		mi := queries[i][1]
		for j := 0; j < n; j++ {
			if nums[j] > mi {
				break
			}
			ans[i] = max(ans[i], xi^nums[j])
		}
	}
	fmt.Println("ans::::::::::::", ans)
	return ans
}
func func3() {
	// 	输入：grid = [[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]
	// 输出：[1,-1,-1,-1,-1]
	findBall([][]int{
		{1, 1, 1, -1, -1},
		{1, 1, 1, -1, -1},
		{-1, -1, -1, 1, 1},
		{1, 1, 1, 1, -1},
		{-1, -1, -1, -1, -1},
	})
}
func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
	var fall func(row, col int) int
	fall = func(row, col int) int {
		if row == m {
			return col
		}
		// 如果卡在了盒子边缘
		if grid[row][col] == 1 {
			// check right
			if col == n-1 {
				return -1
			}
			if grid[row][col+1] == -1 {
				return -1
			}
			return fall(row+1, col+1)
		}
		if col == 0 {
			return -1
		}
		// checK left
		if grid[row][col-1] == 1 {
			return -1
		}
		return fall(row+1, col-1)
	}
	for i := 0; i < n; i++ {
		ans[i] = fall(0, i)
	}
	fmt.Println("ans", ans)
	return ans
}

func halvesAreAlike(s string) bool {
	n := len(s)
	arr := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	check := func(char byte) bool {
		for i := 0; i < len(arr); i++ {
			if arr[i] == char {
				return true
			}
		}
		return false
	}
	i := 0
	counter := 0
	for i < n/2 {
		if check(s[i]) {
			counter++
		}
		i++
	}
	for i < n {
		if check(s[i]) {
			counter--
		}
		i++
	}
	return counter == 0
}

func func2() {
	// 	输入：apples = [1,2,3,5,2], days = [3,2,1,4,2]
	// 输出：7
	// 解释：你可以吃掉 7 个苹果：
	// - 第一天，你吃掉第一天长出来的苹果。
	// - 第二天，你吃掉一个第二天长出来的苹果。
	// - 第三天，你吃掉一个第二天长出来的苹果。过了这一天，第三天长出来的苹果就已经腐烂了。
	// - 第四天到第七天，你吃的都是第四天长出来的苹果。
	// 示例 2：

	// 输入：apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
	// 输出：5
	// 解释：你可以吃掉 5 个苹果：
	// - 第一天到第三天，你吃的都是第一天长出来的苹果。
	// - 第四天和第五天不吃苹果。
	// - 第六天和第七天，你吃的都是第六天长出来的苹果。
	eatenApples([]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2})
	eatenApples([]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2})
	// 	[9,10,1,7,0,2,1,4,1,7,0,11,0,11,0,0,9,11,11,2,0,5,5]
	// [3,19,1,14,0,4,1,8,2,7,0,13,0,13,0,0,2,2,13,1,0,3,7]
	eatenApples([]int{9, 10, 1, 7, 0, 2, 1, 4, 1, 7, 0, 11, 0, 11, 0, 0, 9, 11, 11, 2, 0, 5, 5}, []int{3, 19, 1, 14, 0, 4, 1, 8, 2, 7, 0, 13, 0, 13, 0, 0, 2, 2, 13, 1, 0, 3, 7})
}
func eatenApples(apples []int, days []int) int {
	n := len(apples)
	queue := make([][]int, 0)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		if apples[i] == 0 {
			continue
		}
		badDay := days[i] + i - 1
		number := apples[i]
		start := i
		end := min(badDay, number+i-1)
		queue = append(queue, []int{start, end})
	}
	// get aera
	fmt.Println("queue")
	fmt.Println(queue)
	return 0
}
func eatenApplesOld(apples []int, days []int) int {
	n := len(apples)
	dp := make(map[int]int)
	for i := 0; i < n; i++ {
		if apples[i] == 0 {
			continue
		}
		badDay := days[i] + i
		number := apples[i]
		for j := i; number > 0 && j < badDay; j++ {
			dp[j]++
			number--
			fmt.Println("day ", j, "has apple")
		}
		fmt.Println("i on", i, apples[i])
	}
	fmt.Println(dp)
	ans := len(dp)
	fmt.Println("ans-------------", ans)
	return ans
}
