/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
    // for i:=0;i<len(A);i++ {
	// 	if i%2 != A[i]%2 {
	// 		toSwap := A[i]
	// 		for toSwap%2 != A[toSwap] % 2 {
	// 			tmp:=A[toSwap]
	// 			A[toSwap] = A[i]
	// 			A[i] = tmp
	// 			toSwap = tmp
	// 		}
	// 	}
	// }
	// return A
	// 原地交换可能出现超级大的数字不可以
// 	public int[] sortArrayByParityII(int[] A) {
//         int j = 1;
//         for (int i = 0; i < A.length; i += 2)
//             if (A[i] % 2 == 1) {
//                 while (A[j] % 2 == 1)
//                     j += 2;

//                 // Swap A[i] and A[j]
//                 int tmp = A[i];
//                 A[i] = A[j];
//                 A[j] = tmp;
//             }

//         return A;
//     }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/sort-array-by-parity-ii/solution/an-qi-ou-pai-xu-shu-zu-ii-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	j:=1
	for i:=0;i<len(A);i+=2{
		if A[i] % 2 == 1 {
			for A[j]%2==1 {
				j+=2
			}
			tmp:=A[i]
			A[i]=A[j]
			A[j]=tmp
		}
	}
	return A
}
// @lc code=end

