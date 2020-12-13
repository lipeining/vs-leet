package main

import (
	"fmt"
)

func main() {
	fmt.Println("ans")
	// 	输入：nums = [2,3,5]
	// 输出：[4,3,5]
	// 	result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4，
	// result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3，
	// result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5。
	// 输入：nums = [1,4,6,8,10]
	// 输出：[24,15,13,15,21]
	getSumAbsoluteDifferences([]int{2, 3, 5})
	getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10})
	// 	输入：aliceValues = [1,3], bobValues = [2,1]
	// 输出：1
	// 解释：
	// 如果 Alice 拿石子 1 （下标从 0开始），那么 Alice 可以得到 3 分。
	// Bob 只能选择石子 0 ，得到 2 分。
	// Alice 获胜。
	// 示例 2：

	// 输入：aliceValues = [1,2], bobValues = [3,1]
	// 输出：0
	// 解释：
	// Alice 拿石子 0 ， Bob 拿石子 1 ，他们得分都为 1 分。
	// 打平。
	// 示例 3：

	// 输入：aliceValues = [2,4,3], bobValues = [1,6,7]
	// 输出：-1
	// 解释：
	// 不管 Alice 怎么操作，Bob 都可以得到比 Alice 更高的得分。
	// 比方说，Alice 拿石子 1 ，Bob 拿石子 2 ， Alice 拿石子 0 ，Alice 会得到 6 分而 Bob 得分为 7 分。
	// Bob 会获胜。
	// stoneGameVI([]int{1, 3}, []int{2, 1})       // 1
	// stoneGameVI([]int{1, 2}, []int{3, 1})       // 0
	// stoneGameVI([]int{2, 4, 3}, []int{1, 6, 7}) // -1
	// 	[9,8,3,8]
	// [10,6,9,5]
	// 1
	stoneGameVI([]int{9, 8, 3, 8}, []int{10, 6, 9, 5})
	// 	[5,6,1,9,4,7,1,7,3,7,9,7,1,2,3,9,4,7,6]
	// [1,2,3,5,2,8,5,9,1,6,4,2,10,4,8,1,10,5,6]
	// 1
	// 只能用 dp
	stoneGameVI([]int{5, 6, 1, 9, 4, 7, 1, 7, 3, 7, 9, 7, 1, 2, 3, 9, 4, 7, 6}, []int{1, 2, 3, 5, 2, 8, 5, 9, 1, 6, 4, 2, 10, 4, 8, 1, 10, 5, 6})
}
func stoneGameVI(aliceValues []int, bobValues []int) int {
	// 回溯算法，需要记录双方的得分和已经使用的石头，石头允许回溯
	n := len(aliceValues)
	var dfs func(aliceFlag bool, used []bool, cnt, alice, bob int) int
	dfs = func(aliceFlag bool, used []bool, cnt, alice, bob int) int {
		if cnt == n {
			if alice > bob {
				return 1
			} else if alice < bob {
				return -1
			} else {
				return 0
			}
		}
		hasZero := false
		// prefix := strings.Repeat("#", cnt)
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			// 这里要考虑不让对方拿到这个高的分数
			if aliceFlag {
				// fmt.Println(prefix+"alice pick", i)
				ret := dfs(false, used, cnt+1, alice+aliceValues[i], bob)
				// fmt.Println(prefix+"alice pick", i, "and result", ret)
				if ret > 0 {
					// 这里会导致他人无法复用这个 i, 别人有机会绕过这个结果的
					used[i] = false
					return ret
				} else if ret == 0 {
					hasZero = true
				}
			} else {
				// fmt.Println(prefix+"bob pick", i)
				ret := dfs(true, used, cnt+1, alice, bob+bobValues[i])
				// fmt.Println(prefix+"bob pick", i, "and result", ret)
				if ret < 0 {
					used[i] = false
					return ret
				} else if ret == 0 {
					hasZero = true
				}
			}
			used[i] = false
		}
		if hasZero {
			return 0
		}
		if aliceFlag {
			return -1
		}
		return 1
	}
	used := make([]bool, n)
	ans := dfs(true, used, 0, 0, 0)
	fmt.Println("ans----------------------------------------------", ans)
	return ans
}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	tmp := 0
	// 	for j := 0; j < n; j++ {
	// 		if j <= i {
	// 			tmp += nums[i] - nums[j]
	// 		} else {
	// 			tmp += nums[j] - nums[i]
	// 		}
	// 	}
	// 	ans[i] = tmp
	// }
	// fmt.Println("ans", ans)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	// fmt.Println("pre", pre)
	for i := 0; i < n; i++ {
		left := pre[i] - pre[0]
		right := pre[n] - pre[i]
		// fmt.Println(left, right)
		ans[i] = right - (n-i)*nums[i] + i*nums[i] - left
		// ans[i] =
	}
	fmt.Println("ans", ans)
	return ans
}
