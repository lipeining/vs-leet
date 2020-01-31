/*
 * @lc app=leetcode.cn id=1054 lang=golang
 *
 * [1054] 距离相等的条形码
 */

// @lc code=start
// func rearrangeBarcodes(barcodes []int) []int {
    
// }
func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(barcodes); i++ {
		m[barcodes[i]]++
	}
	n := make([]Node, 0, len(m))
	for k, v := range m {
		n = append(n, Node{k, v})
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i].count > n[j].count
	})
	ret := make([]int, len(barcodes))
	start := 0
	for _, v := range n {
		val, count := v.val, v.count
		for count > 0 {
			ret[start] = val
			start += 2
			count--
			if start >= len(barcodes) {
				start = 1
			}
		}
	}
	return ret
}

type Node struct {
	val int
	count int
}

// 作者：resara
// 链接：https://leetcode-cn.com/problems/distant-barcodes/solution/golang-pai-xu-tian-chong-by-resara/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end

