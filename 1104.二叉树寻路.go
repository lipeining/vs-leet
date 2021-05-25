/*
 * @lc app=leetcode.cn id=1104 lang=golang
 *
 * [1104] 二叉树寻路
 */

// @lc code=start
func pathInZigZagTree(label int) []int {
	var res []int
	// 按照正常完全二叉树的排序计算路径上的数字
	for num := label; num != 0; num /= 2 {
			res = append(res, num)
	}
	n := len(res)
	// 得到的切片中的数字是降序排序的，将其反转，得到升序排序
	for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
			res[i], res[j] = res[j], res[i]
	}
	var i int
	if n % 2 == 0 {
			i = 2
	} else {
			i = 1
	}
	// 对需要进行对称变换的层的值进行变换
	for ; i < n; i += 2 {
			res[i] = 1 << i * 3 - 1 - res[i]
	}
	return res

// 作者：sogayo-2
// 链接：https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/solution/1104-er-cha-shu-xun-lu-go-shuang-100-by-sogayo-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

// @lc code=end

