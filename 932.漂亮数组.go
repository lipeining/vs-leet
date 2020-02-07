/*
 * @lc app=leetcode.cn id=932 lang=golang
 *
 * [932] 漂亮数组
 */

// @lc code=start
func beautifulArray(N int) []int {
	// 带返回值的 dfs
	// 我们将其分治成两个子问题，
	// 其中一个为不超过 (N + 1) / 2 的整数，
	// 并映射到所有的奇数；
	// 另一个为不超过 N / 2 的整数，
	// 并映射到所有的偶数。
	cache := make(map[int][]int)
	var dfs func(N int) []int
	dfs = func(N int) []int {
		if list, ok := cache[N]; ok {
			return list
		}
		ans := make([]int, N)
		if N == 1 {
			ans[0] = 1
			return ans
		}
		t := 0
		for _, x := range dfs((N + 1) / 2) {
			ans[t] = 2*x - 1
			t++
		}
		for _, x := range dfs(N / 2) {
			ans[t] = 2 * x
			t++
		}
		cache[N] = ans
		return ans
	}
	res := dfs(N)
	return res
}

// @lc code=end

// Map<Integer, int[]> memo;
// public int[] beautifulArray(int N) {
// 	memo = new HashMap();
// 	return f(N);
// }

// public int[] f(int N) {
// 	if (memo.containsKey(N))
// 		return memo.get(N);

// 	int[] ans = new int[N];
// 	if (N == 1) {
// 		ans[0] = 1;
// 	} else {
// 		int t = 0;
// 		for (int x: f((N+1)/2))  // odds
// 			ans[t++] = 2*x - 1;
// 		for (int x: f(N/2))  // evens
// 			ans[t++] = 2*x;
// 	}
// 	memo.put(N, ans);
// 	return ans;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/beautiful-array/solution/piao-liang-shu-zu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。