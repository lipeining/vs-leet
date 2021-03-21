package main

import "fmt"

func main() {
	// func1()
	// func2()
	func3()
	func4()
}
func func1() {

}
func func2() {
	// 输入：deliciousness = [1,3,5,7,9]
	// 输出：4
	// 解释：大餐的美味程度组合为 (1,3) 、(1,7) 、(3,5) 和 (7,9) 。
	// 它们各自的美味程度之和分别为 4 、8 、8 和 16 ，都是 2 的幂。
	// 示例 2：

	// 输入：deliciousness = [1,1,1,3,3,3,7]
	// 输出：15
	// 解释：大餐的美味程度组合为 3 种 (1,1) ，9 种 (1,3) ，和 3 种 (1,7) 。
	// countPairs([]int{1, 3, 5, 7, 9})
	// countPairs([]int{1, 1, 1, 3, 3, 3, 7})
	// [149,107,1,63,0,1,6867,1325,5611,2581,39,89,46,18,12,20,22,234]
	// 12
	countPairs([]int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234})
}
func countPairs(deliciousness []int) int {
	n := len(deliciousness)
	pro := make(map[int]bool)
	cnt := 1
	for i := 1; i <= 22; i++ {
		pro[cnt] = true
		cnt *= 2
	}
	ans := 0
	toMod := int(1e9 + 7)
	numMap := make(map[int]int)
	for i := 0; i < n; i++ {
		for k := range pro {
			lastNum := k - deliciousness[i]
			if last, ok := numMap[lastNum]; ok {
				// 有一对了
				fmt.Println("pair", k, deliciousness[i], lastNum, "ans add", last)
				ans += last
				ans %= toMod
			}
		}
		numMap[deliciousness[i]]++
	}
	fmt.Println("ans------", ans)
	return ans
}
func func3() {
	// 输入：nums = [1,1,1]
	// 输出：1
	// 解释：唯一一种好的分割方案是将 nums 分成 [1] [1] [1] 。
	// 示例 2：

	// 输入：nums = [1,2,2,2,5,0]
	// 输出：3
	// 解释：nums 总共有 3 种好的分割方案：
	// [1] [2] [2,2,5,0]
	// [1] [2,2] [2,5,0]
	// [1,2] [2,2] [5,0]
	// 示例 3：

	// 输入：nums = [3,2,1]
	// 输出：0
	// 解释：没有好的分割方案。
	waysToSplit([]int{1, 1, 1})
	waysToSplit([]int{1, 2, 2, 2, 5, 0})
	waysToSplit([]int{3, 2, 1})
}
func waysToSplit(nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	toMod := int(1e9 + 7)
	ans := 0
	ans %= toMod
	// i 是左边的下标，j是Mid的下标
	// right的下标一直都是 n
	// 递增队列？
	// for i := 1; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		left := pre[i] - pre[0]
	// 		mid := pre[j] - pre[i]
	// 		right := pre[n] - pre[j]
	// 		// fmt.Println("i,j,n", i, j, n)
	// 		// fmt.Println(left, mid, right)
	// 		if left <= mid && mid <= right {
	// 			ans++
	// 			ans %= toMod
	// 		}
	// 	}
	// }
	for i := 1; i <= n-2; i++ {
		left := pre[i] - pre[0]
		j := n - 1
		for j > i {
			mid := pre[j] - pre[i]
			right := pre[n] - pre[j]
			if left <= mid && mid <= right {
				break
			}
			j--
		}
	}
	// fmt.Println("ans", ans)
	return ans
}
func func4() {

}
