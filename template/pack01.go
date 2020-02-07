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
