/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	// ans := make([]int, 1)
	// ans[0]=0
	// head := 1
	// for i:=0;i<n;i++ {
	// 	for j := len(ans)-1;j>=0;j-- {
	// 		ans = append(ans, ans[j] + head)
	// 	}
	// 	head <<= 1
	// }
	// return ans
	ans := make([]int, 0)
	var dfs func(now string, x int)
	dfs = func(now string, x int) {
		if len(now) == n {
			tmp,_ := strconv.ParseInt(now, 2, 0)
			ans = append(ans, int(tmp))
			return
		}
		if x == 0 {
			dfs(now+"0", 0)
			dfs(now+"1", 1)
		} else {
			dfs(now+"1", 0)
			dfs(now+"0", 1)
		}
	}
	dfs("", 0)
	return ans
// 	def grayCode(self, n: int) -> List[int]:
// 	if n==0:
// 		return [0]

// 	res=[]
// 	def back(now,x):
// 		if len(now)==n:
// 			res.append(int(now,2))
// 		elif x==0:
// 			back(now+'0',0)
// 			back(now+'1',1)
// 		else:
// 			back(now+'1',0)
// 			back(now+'0',1)
	
// 	back('',0)
// 	return res

// 作者：dannnn
// 链接：https://leetcode-cn.com/problems/gray-code/solution/jian-dan-de-si-lu-44ms-by-dannnn/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// 	public List<Integer> grayCode(int n) {
//         List<Integer> res = new ArrayList<Integer>() {{ add(0); }};
//         int head = 1;
//         for (int i = 0; i < n; i++) {
//             for (int j = res.size() - 1; j >= 0; j--)
//                 res.add(head + res.get(j));
//             head <<= 1;
//         }
//         return res;
//     }

// 作者：jyd
// 链接：https://leetcode-cn.com/problems/gray-code/solution/gray-code-jing-xiang-fan-she-fa-by-jyd/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

}
// @lc code=end

