package template

func canPartition(nums []int) bool {
	// 背包问题
	sum := 0
	length := len(nums)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	// 是否有 dp[i][j] = sum/2
	// 是否有选择一部分数组得到和为N的结果
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	// boolean[] dp = new boolean[W + 1];
	// dp[0] = true;
	// for (int num : nums) {                 // 0-1 背包一个物品只能用一次
	//     for (int i = W; i >= num; i--) {   // 从后往前，先计算 dp[i] 再计算 dp[i-num]
	//         dp[i] = dp[i] || dp[i - num];
	//     }
	// }
	// return dp[W];
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := half; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[half]
}

// 0-1背包问题是指每一种物品都只有一件，可以选择放或者不放。现在假设有n件物品，背包承重为m。
// 对于这种问题，我们可以采用一个二维数组去解决：f[i][j]，其中i代表加入背包的是前i件物品，j表示背包的承重，
// f[i][j]表示当前状态下能放进背包里面的物品的最大总价值。那么，f[n][m]就是我们的最终结果了。
// f[i][j] = max(f[i][j] = f[i - 1][j] , f[i - 1][j - weight[i]] + value[i])   j - weight[i]>=0
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func pack01(weight, value []int, m int) int {
	n := len(weight)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j >= weight[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][m]
}
func pack01Simple(weight, value []int, m int) int {
	n := len(weight)
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := m; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[m]
}

// 完全背包
func pack(weight, value []int, m int) int {
	n := len(weight)
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := weight[i]; j <= m; j++ {
			// 背包中 物品是无限的，需要拿当前的背包计算价值，空间有限
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[m]
}

// 多重背包问题限定了一种物品的个数，解决多重背包问题，只需要把它转化为0-1背包问题即可。
// 比如，有2件价值为5，重量为2的同一物品，我们就可以分为物品a和物品b，
// a和b的价值都为5，重量都为2，但我们把它们视作不同的物品。

func multiPack(weight, value, nums []int, m int) int {
	n := len(weight)
	k := n + 1
	for i := 1; i <= n; i++ {
		for nums[i] != 1 {
			weight[k] = weight[i]
			value[k] = value[i]
			k++
			nums[i]--
		}
	}
	dp := make([]int, m+1)
	for i := 1; i <= k; i++ {
		for j := m; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[m]
}
