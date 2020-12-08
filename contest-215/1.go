package main

import (
	"fmt"
	"math"
)

func main() {
	// ans := 0
	// fmt.Println("ans", ans)
	// fmt.Println(closeStrings("abbzzca", "babzzcz"), "-0-------")
	// fmt.Println(closeStrings("cabbba", "abbccc"), "-1---")
	// fmt.Println(closeStrings("cabbba", "aabbss"), "-2---")

	// 输入：nums = [1,1,4,2,3], x = 5
	// 输出：2
	// 解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
	// 示例 2：
	minOperations([]int{1, 1}, 3)
	minOperations([]int{1, 1, 4, 2, 3}, 5)
	// 输入：nums = [5,6,7,8,9], x = 4
	// 输出：-1
	// 示例 3：
	minOperations([]int{5, 6, 7, 8, 9}, 4)
	// 输入：nums = [3,2,20,1,1,3], x = 10
	// 输出：5
	// 解释：最佳解决方案是移除后三个元素和前两个元素（总共 5 次操作），将 x 减到 0 。
	minOperations([]int{3, 2, 20, 1, 1, 3}, 10)
	minOperations([]int{1241, 8769, 9151, 3211, 2314, 8007, 3713, 5835, 2176, 8227, 5251, 9229, 904, 1899, 5513, 7878, 8663, 3804, 2685, 3501, 1204, 9742, 2578, 8849, 1120, 4687, 5902, 9929, 6769, 8171, 5150, 1343, 9619, 3973, 3273, 6427, 47, 8701, 2741, 7402, 1412, 2223, 8152, 805, 6726, 9128, 2794, 7137, 6725, 4279, 7200, 5582, 9583, 7443, 6573, 7221, 1423, 4859, 2608, 3772, 7437, 2581, 975, 3893, 9172, 3, 3113, 2978, 9300, 6029, 4958, 229, 4630, 653, 1421, 5512, 5392, 7287, 8643, 4495, 2640, 8047, 7268, 3878, 6010, 8070, 7560, 8931, 76, 6502, 5952, 4871, 5986, 4935, 3015, 8263, 7497, 8153, 384, 1136}, 894887480)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minOperations(nums []int, x int) int {
	n := len(nums)
	ans := math.MaxInt32
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	// 过了 90% 个用例，超时了 77/86
	// fmt.Println(pre)
	// 从左边挑选 i 个，右边挑选 j 个
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			left := pre[i] - pre[0]
			right := pre[n] - pre[n-j]
			now := left + right
			if now == x {
				fmt.Println(i, j, left, right, now)
				ans = min(ans, i+j)
			}
			if now > x {
				break
			}
		}
	}
	fmt.Println("ans", ans, "----------------")
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
func minOperations2(nums []int, x int) int {
	// 贪心 dfs
	n := len(nums)
	ans := math.MaxInt32
	var dfs func(left, right, now, action int)
	dfs = func(left, right, now, action int) {
		if now < 0 {
			// fmt.Println("--next--")
			return
		}
		if now == 0 {
			ans = min(ans, action)
			// fmt.Println("--next--")
			return
		}
		if left > right {
			return
		}
		if now >= nums[left] {
			dfs(left+1, right, now-nums[left], action+1)
			// fmt.Println("-")
		}
		if now >= nums[right] {
			dfs(left, right-1, now-nums[right], action+1)
		}
	}
	dfs(0, n-1, x, 0)
	fmt.Println("ans", ans)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func closeStrings(word1 string, word2 string) bool {
	n := len(word1)
	if n != len(word2) {
		return false
	}
	c1 := make(map[byte]int)
	c2 := make(map[byte]int)
	for i := 0; i < n; i++ {
		c1[word1[i]]++
		c2[word2[i]]++
	}
	fmt.Println(c1, c2)
	// 两者都有相同的 key， 相同数量的 key 要相等
	cn1 := make(map[int]int)
	cn2 := make(map[int]int)
	for k, v := range c1 {
		if _, ok := c2[k]; !ok {
			return false
		}
		cn1[v]++
	}
	for k, v := range c2 {
		if _, ok := c1[k]; !ok {
			return false
		}
		cn2[v]++
	}

	fmt.Println(cn1, cn2)
	// 保证 cn1 == cn2
	for k, v := range cn1 {
		if v2, ok := cn2[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}
