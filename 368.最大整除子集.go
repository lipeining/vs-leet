import "sort"

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	sz := len(nums)
	mx, end := 0, -1
	dp := make([]int, sz)
	last := make([]int, sz)
	for i := 0; i < sz; i++ {
		dp[i] = 1
		last[i] = -1
	}
	sort.Ints(nums)
	for i := 0; i < sz; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
				last[i] = j
			}
		}
		if dp[i] > mx {
			mx = dp[i]
			end = i
		}
	}
	ans := make([]int, 0)
	for i := end; i != -1; i = last[i] {
		ans = append(ans, nums[i])
	}
	return ans
}

// @lc code=end

// vector<int> largestDivisibleSubset(vector<int>& nums) {
// 	int sz = nums.size(),mx = 0,end = -1;
// 	vector<int> dp(sz,1),last(sz,-1),res;
// 	sort(nums.begin(),nums.end());
// 	for(int i = 0;i<sz;i++){
// 		for(int j = 0;j<i;j++){
// 			if(nums[i]%nums[j] == 0 && dp[i]<=dp[j]){
// 				dp[i] = dp[j]+1;
// 				last[i] = j;
// 			}
// 		}
// 		if(dp[i]>mx){
// 			mx = dp[i];
// 			end = i;
// 		}
// 	}
// 	for(int i = end;i!=-1;i = last[i]){//倒序输出
// 		res.push_back(nums[i]);
// 	}
// 	return res;
// }

// 作者：leolee
// 链接：https://leetcode-cn.com/problems/largest-divisible-subset/solution/ge-ren-ti-jie-dpcon2-by-leolee/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。