/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(S string) []int {
	var last [26]int
	length := len(S)
	for i := 0; i < length; i++ {
		last[S[i]-'a'] = i
	}
	j, anchor := 0, 0
	ans := make([]int, 0)
	for i := 0; i < length; i++ {
		j = max(j, last[S[i]-'a'])
		if i == j {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
// public List<Integer> partitionLabels(String S) {
// 	int[] last = new int[26];
// 	for (int i = 0; i < S.length(); ++i)
// 		last[S.charAt(i) - 'a'] = i;

// 	int j = 0, anchor = 0;
// 	List<Integer> ans = new ArrayList();
// 	for (int i = 0; i < S.length(); ++i) {
// 		j = Math.max(j, last[S.charAt(i) - 'a']);
// 		if (i == j) {
// 			ans.add(i - anchor + 1);
// 			anchor = i + 1;
// 		}
// 	}
// 	return ans;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/partition-labels/solution/hua-fen-zi-mu-qu-jian-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
