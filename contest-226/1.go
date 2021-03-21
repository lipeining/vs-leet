package main

import (
	"fmt"
)

func main() {
	// func1()
	// func2()
	func3()
	// func4()
	fmt.Println("anssssssssssssssss end")
}
func func3() {
	// 	输入：candiesCount = [7,4,5,3,8], queries = [[0,2,2],[4,2,4],[2,13,1000000000]]
	// 输出：[true,false,true]
	// 提示：
	// 1- 在第 0 天吃 2 颗糖果(类型 0），第 1 天吃 2 颗糖果（类型 0），第 2 天你可以吃到类型 0 的糖果。
	// 2- 每天你最多吃 4 颗糖果。即使第 0 天吃 4 颗糖果（类型 0），第 1 天吃 4 颗糖果（类型 0 和类型 1），你也没办法在第 2 天吃到类型 4 的糖果。换言之，你没法在每天吃 4 颗糖果的限制下在第 2 天吃到第 4 类糖果。
	// 3- 如果你每天吃 1 颗糖果，你可以在第 13 天吃到类型 2 的糖果。
	// 示例 2：

	// 输入：candiesCount = [5,2,6,4,1], queries = [[3,1,2],[4,10,3],[3,10,100],[4,100,30],[1,3,1]]
	// 输出：[false,true,true,false,false]
	canEat([]int{1, 1, 1, 1, 1}, [][]int{{0, 2, 1}, {4, 2, 2}, {4, 2, 1}})
	canEat([]int{1, 1, 1, 2, 1}, [][]int{{0, 2, 1}, {4, 2, 2}, {4, 2, 2}})
	canEat([]int{7, 4, 5, 3, 8}, [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 100000000}})
	canEat([]int{5, 2, 6, 4, 1}, [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}})
}

// [16,38,8,41,30,31,14,45,3,2,24,23,38,30,31,17,35,4,9,42,28,18,37,18,14,46,11,13,19,3,5,39,24,48,20,29,4,19,36,11,28,49,38,16,23,24,4,22,29,35,45,38,37,40,2,37,8,41,33,8,40,27,13,4,33,5,8,14,19,35,31,8,8]
// [[35,669,5],[72,822,74],[47,933,94],[62,942,85],[42,596,11],[56,1066,18],[54,571,45],[39,890,100],[3,175,26],[48,1489,37],[40,447,52],[30,584,7],[26,1486,38],[21,1142,21],[9,494,96],[56,759,81],[13,319,16],[20,1406,57],[11,1092,19],[24,670,67],[38,1702,33],[5,676,32],[50,1386,77],[36,1551,87],[29,1445,13],[58,977,13],[7,887,64],[37,1396,23],[0,765,69],[40,1083,86],[43,1054,49],[48,690,92],[28,1201,56],[47,948,43],[57,233,25],[32,1293,65],[0,1646,34],[43,1467,39],[39,484,23],[21,1576,69],[12,1222,68],[9,457,83],[32,65,9],[10,1424,42],[35,534,3],[23,83,22],[33,501,33],[25,679,51],[2,321,42],[1,240,68],[7,1297,42],[45,480,72],[26,1472,9],[6,649,90],[26,361,57],[49,1592,7],[11,158,95],[35,448,24],[41,1654,10],[61,510,43],[31,1230,95],[11,1471,12],[37,43,84],[56,1147,48],[69,1368,65],[22,170,24],[56,192,80],[34,1207,69],[1,1226,22],[37,1633,50],[11,98,58],[17,125,13],[0,1490,5],[37,1732,43],[45,793,14],[16,578,72],[50,241,78]]
func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + candiesCount[i]
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		t, day, limit := q[0], q[1], q[2]
		before := pre[t] - pre[0]
		total := pre[t+1] - pre[0]
		totalLimit := (day + 1) * limit
		fmt.Println("for t", t, "has before", before, "total", total)
		fmt.Println("totalLimit", totalLimit)
		if before >= totalLimit || total <= day {
			continue
		}
		ans[i] = true
	}
	fmt.Println("ansss", ans)
	return ans
}
func func4() {
	// 	输入：s = "abcbdd"
	// 输出：true
	// 解释："abcbdd" = "a" + "bcb" + "dd"，三个子字符串都是回文的。
	// 示例 2：

	// 输入：s = "bcbddxy"
	// 输出：false
	// 解释：s 没办法被分割成 3 个回文子字符串。
	checkPartitioning("abcbdd")
	checkPartitioning("bcbddxy")
}
func checkPartitioning(s string) bool {
	n := len(s)
	// dp[i][1] 到 i 为 1个回文的情况
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, 4)
	}
	dp[0][1] = true
	dp[1][2] = true
	checkPair := func(start, end int) bool {
		left := start
		right := end
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		// fmt.Println("for start, end", start, end, true)
		return true
	}
	for i := 2; i < n; i++ {
		dp[i][1] = checkPair(0, i)
		for j := 0; j < i; j++ {
			dp[i][2] = dp[i][2] || (dp[j][1] && checkPair(j+1, i))
			dp[i][3] = dp[i][3] || (dp[j][2] && checkPair(j+1, i))
			fmt.Println("in i", i, dp[i])
		}
	}
	ans := dp[n-1][3]
	fmt.Println("ansssss", dp, ans)
	return ans
}
func func1() {
	countBalls(1, 10)
	countBalls(5, 15)
	countBalls(19, 28)
}
func countBalls(lowLimit int, highLimit int) int {
	counter := make(map[int]int)
	target := 0
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	getSum := func(n int) int {
		now := 0
		for n != 0 {
			m := n % 10
			now += m
			n /= 10
		}
		return now
	}
	for i := lowLimit; i <= highLimit; i++ {
		sum := getSum(i)
		counter[sum]++
		if counter[sum] > counter[target] {
			target = sum
		}
		fmt.Println("for i", i, target, sum)
	}
	fmt.Println(counter, target)
	return counter[target]
}
func func2() {
	// 	输入：adjacentPairs = [[2,1],[3,4],[3,2]]
	// 输出：[1,2,3,4]
	// 解释：数组的所有相邻元素对都在 adjacentPairs 中。
	// 特别要注意的是，adjacentPairs[i] 只表示两个元素相邻，并不保证其 左-右 顺序。
	// 示例 2：

	// 输入：adjacentPairs = [[4,-2],[1,4],[-3,1]]
	// 输出：[-2,4,1,-3]
	// 解释：数组中可能存在负数。
	// 另一种解答是 [-3,1,4,-2] ，也会被视作正确答案。
	// 示例 3：

	// 输入：adjacentPairs = [[100000,-100000]]
	// 输出：[100000,-100000]
	restoreArray([][]int{{2, 1}, {3, 4}, {3, 2}})
	restoreArray([][]int{{4, -2}, {1, 4}, {-3, 1}})
	restoreArray([][]int{{100000, -100000}})
}

// func restoreArray(adjacentPairs [][]int) []int {
// 	in := make(map[int]int)
// 	for _, p := range adjacentPairs {
// 		a, b := p[0], p[1]
// 		in[a]++
// 		in[b]++
// 	}
// 	m := len(adjacentPairs)
// 	n := m + 1
// 	ans := make([]int, n)
// 	queue := make([]int, 0)
// 	for k, v := range in {
// 		if v == 1 {
// 			queue = append(queue, k)
// 		}
// 	}
// 	fmt.Println(queue)
// 	i := 0
// 	for len(queue) != 0 {
// 		size := len(queue)
// 		fmt.Println("size ", size)
// 		// 应该只有两个
// 		if size == 2 {
// 			ans[i] = queue[0]
// 			ans[n-i-1] = queue[1]
// 		} else {
// 			ans[i] = queue[0]
// 		}
// 		i++
// 		queue = queue[size:]
// 	}
// 	// for i:=0;i<n/2;i++ {

// 	// }
// }
func restoreArray(adjacentPairs [][]int) []int {
	table := make(map[int]map[int]bool)
	for _, p := range adjacentPairs {
		// table[p[0]] = append(table[p[0]], p[1])
		// table[p[1]] = append(table[p[1]], p[0])
		a, b := p[0], p[1]
		if table[a] == nil {
			table[a] = make(map[int]bool)
		}
		table[a][b] = true
		if table[b] == nil {
			table[b] = make(map[int]bool)
		}
		table[b][a] = true
	}
	// 如果相邻的只有一个的话，说明是最边的元素
	m := len(adjacentPairs)
	n := m + 1
	ans := make([]int, n)
	queue := make([]int, 0)
	visited := make(map[int]bool)
	for k, v := range table {
		if len(v) == 1 {
			queue = append(queue, k)
			visited[k] = true
		}
	}
	fmt.Println(queue)
	i := 0
	for len(queue) != 0 {
		size := len(queue)
		fmt.Println("size ", size, queue)
		// 应该只有两个
		if size == 2 {
			ans[i] = queue[0]
			ans[n-i-1] = queue[1]
			for next := range table[queue[0]] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
			for next := range table[queue[1]] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		} else {
			ans[i] = queue[0]
		}
		i++
		queue = queue[size:]
	}
	fmt.Println("anssssss", ans)
	return ans
}
