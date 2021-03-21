package main

import (
	"fmt"
	"math"
)

func main() {
	// 输入：nums = [4,2,4,5,6]
	// 输出：17
	// 解释：最优子数组是 [2,4,5,6]
	// 示例 2：

	// 输入：nums = [5,2,1,2,5,2,1,2,5]
	// 输出：8
	// 解释：最优子数组是 [5,2,1] 或 [1,2,5]
	// maximumUniqueSubarray([]int{4, 2, 4, 5, 6})
	// maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})

	// 	输入：nums = [1,-1,-2,4,-7,3], k = 2
	// 输出：7
	// 解释：你可以选择子序列 [1,-1,4,3] （上面加粗的数字），和为 7 。
	// 示例 2：

	// 输入：nums = [10,-5,-2,4,0,3], k = 3
	// 输出：17
	// 解释：你可以选择子序列 [10,4,3] （上面加粗数字），和为 17 。
	// 示例 3：

	// 输入：nums = [1,-5,-20,4,-1,3,-6,-3], k = 2
	// 输出：0
	// maxResult([]int{1, -1, -2, 4, -7, 3}, 2)          // 7
	// maxResult([]int{10, -5, -2, 4, 0, 3}, 3)          // 17
	// maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2) // 0
	// 输入：n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
	// 输出：[false,true]
	// 解释：上图为给定的输入数据。注意到 0 和 1 之间有两条重边，分别为 2 和 16 。
	// 对于第一个查询，0 和 1 之间没有小于 2 的边，所以我们返回 false 。
	// 对于第二个查询，有一条路径（0 -> 1 -> 2）两条边都小于 5 ，所以这个查询我们返回 true 。
	// 输入：n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
	// 输出：[true,false]
	// 解释：上图为给定数据。
	distanceLimitedPathsExist(3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}, [][]int{{0, 1, 2}, {0, 2, 5}})
	distanceLimitedPathsExist(5, [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}, [][]int{{0, 4, 14}, {0, 4, 13}})
}
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	gMap := make(map[int]map[int]int)
	for _, edge := range edgeList {
		v, w, value := edge[0], edge[1], edge[2]
		if lastMap, ok := gMap[v]; ok {
			if wValue, wOk := lastMap[w]; wOk {
				if wValue > value {
					gMap[v][w] = value
				}
			} else {
				gMap[v][w] = value
			}
		} else {
			gMap[v] = map[int]int{w: value}
		}
		if lastMap, ok := gMap[w]; ok {
			if wValue, wOk := lastMap[v]; wOk {
				if wValue > value {
					gMap[w][v] = value
				}
			} else {
				gMap[w][v] = value
			}
		} else {
			gMap[w] = map[int]int{v: value}
		}
	}
	var dfs func(v int, marked []bool, search, limit int) bool
	dfs = func(v int, marked []bool, search, limit int) bool {
		if marked[v] {
			return false
		}
		if v == search {
			return true
		}
		marked[v] = true
		for w, wValue := range gMap[v] {
			if marked[w] {
				continue
			}
			if wValue < limit {
				// fmt.Println("v", v, "w", w, "wValue", wValue)
				if dfs(w, marked, search, limit) {
					return true
				}
			}
		}
		marked[v] = false
		return false
	}
	// fmt.Println("gMap", gMap)
	ans := make([]bool, len(queries))
	for index, query := range queries {
		marked := make([]bool, n)
		ans[index] = dfs(query[0], marked, query[1], query[2])
		// fmt.Println("-------------next")
	}
	// fmt.Println("ans", ans)
	return ans
}

func maxResult(nums []int, k int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	memo := make(map[int]int)
	memo[0] = nums[0]
	var dfs func(index int) int
	dfs = func(index int) int {
		if value, ok := memo[index]; ok {
			return value
		}
		if index == 0 {
			return nums[0]
		}
		node := math.MinInt32
		for j := index - 1; j >= 0 && j >= index-k; j-- {
			node = max(node, dfs(j)+nums[index])
		}
		memo[index] = node
		return node
	}
	ans := dfs(n - 1)
	fmt.Println("ans::::::::::::", ans)
	return ans
}
func maxResultDP(nums []int, k int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if k >= n-1 {
		ans := nums[0]
		for i := 1; i < n-1; i++ {
			// ans = max(ans, ans+nums[i])
			if nums[i] > 0 {
				ans += nums[i]
			}
		}
		fmt.Println("ans:   ", ans)
		return ans + nums[n-1]
	}
	dp := make([]int, n)
	// 调到位置 i 的最大得分是多少
	// 需要根据i-k到i进行计算
	// 超时了 16309 的步数
	for i := 0; i < n; i++ {
		dp[i] = math.MinInt32
	}
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		for j := i - 1; j >= i-k && j >= 0; j-- {
			dp[i] = max(dp[i], dp[j]+nums[i])
		}
		// fmt.Println("dp", dp)
	}
	// fmt.Println("ans", dp[n-1])
	return dp[n-1]
}
func maximumUniqueSubarray(nums []int) int {
	lastMap := make(map[int]int)
	n := len(nums)
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// check := func(checkMap map[int]int) bool {
	// 	for _, v := range checkMap {
	// 		if v > 1 {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }
	// 应该用滑动窗口
	left := 0
	right := 0
	sum := 0
	for right < n {
		num := nums[right]
		sum += num
		lastMap[num]++
		lastCount, ok := lastMap[num]
		for ok && lastCount > 1 {
			// fmt.Println("Left, sum , ans", left, sum, ans)
			sum -= nums[left]
			lastMap[nums[left]]--
			left++
			lastCount, ok = lastMap[num]
		}
		ans = max(ans, sum)
		right++
	}
	// fmt.Println("ans:::::::::::::", ans)
	return ans
}
