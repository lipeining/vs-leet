package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// 输入：command = "G()(al)"
	// 输出："Goal"
	// 解释：Goal 解析器解释命令的步骤如下所示：
	// G -> G
	// () -> o
	// (al) -> al
	// 最后连接得到的结果是 "Goal"
	// 示例 2：

	// 输入：command = "G()()()()(al)"
	// 输出："Gooooal"
	// 示例 3：

	// 输入：command = "(al)G(al)()()G"
	// 输出："alGalooG"
	interpret("G()(al)")
	interpret("G()()()()(al)")
	interpret("(al)G(al)()()G")

	// 	输入：nums = [1,2,3,4], k = 5
	// 输出：2
	// 解释：开始时 nums = [1,2,3,4]：
	// - 移出 1 和 4 ，之后 nums = [2,3]
	// - 移出 2 和 3 ，之后 nums = []
	// 不再有和为 5 的数对，因此最多执行 2 次操作。
	// 示例 2：

	// 输入：nums = [3,1,3,4,3], k = 6
	// 输出：1
	// 解释：开始时 nums = [3,1,3,4,3]：
	// - 移出前两个 3 ，之后nums = [1,4,3]
	// 不再有和为 6 的数对，因此最多执行 1 次操作。
	maxOperations([]int{1, 2, 3, 4}, 5)
	maxOperations([]int{3, 1, 3, 4, 3}, 6)

	// 	输入：n = 1
	// 输出：1
	// 解释：二进制的 "1" 对应着十进制的 1 。
	// 示例 2：

	// 输入：n = 3
	// 输出：27
	// 解释：二进制下，1，2 和 3 分别对应 "1" ，"10" 和 "11" 。
	// 将它们依次连接，我们得到 "11011" ，对应着十进制的 27 。
	// 示例 3：

	// 输入：n = 12
	// 输出：505379714
	// 解释：连接结果为 "1101110010111011110001001101010111100" 。
	// 对应的十进制数字为 118505380540 。
	// 对 109 + 7 取余后，结果为 505379714 。
	concatenatedBinary(1)
	concatenatedBinary(3)
	concatenatedBinary(12)
	concatenatedBinary(42)
	// 	输入：
	// 42
	// 输出：
	// 697078543
	// 预期：
	// 727837408
	// 	输入：nums = [1,2,1,4], k = 2
	// 输出：4
	// 解释：最优的分配是 [1,2] 和 [1,4] 。
	// 不兼容性和为 (2-1) + (4-1) = 4 。
	// 注意到 [1,1] 和 [2,4] 可以得到更小的和，但是第一个集合有 2 个相同的元素，所以不可行。
	// 示例 2：

	// 输入：nums = [6,3,8,1,3,1,2,2], k = 4
	// 输出：6
	// 解释：最优的子集分配为 [1,2]，[2,3]，[6,8] 和 [1,3] 。
	// 不兼容性和为 (2-1) + (3-2) + (8-6) + (3-1) = 6 。
	// 示例 3：

	// 输入：nums = [5,3,3,6,3,3], k = 3
	// 输出：-1
	// 解释：没办法将这些数字分配到 3 个子集且满足每个子集里没有相同数字。
}
func minimumIncompatibility(nums []int, k int) int {
	sort.Ints(nums)
	// n := len(nums)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for _, v := range m {
		if v > k {
			return -1
		}
	}
	return 0
}
func concatenatedBinary(n int) int {
	toMod := int64(1e9 + 7)
	ans := int64(0)
	multi := int64(1)
	for i := n; i >= 1; i-- {
		num := i
		for ; num > 0; num /= 2 {
			lsb := num % 2
			if lsb == 1 {
				ans += multi
			}
			ans %= toMod
			multi *= 2
			multi %= toMod
		}
	}
	fmt.Println("ans", ans)
	return int(ans)
}
func concatenatedBinary2(n int) int {
	str := ""
	tenToBit := func(num int) string {
		s := ""
		if num == 0 {
			return "0"
		}
		// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
		for ; num > 0; num /= 2 {
			lsb := num % 2
			// strconv.Itoa() 将数字强制性转化为字符串
			s = strconv.Itoa(lsb) + s
		}
		return s
	}
	for i := 1; i <= n; i++ {
		str += tenToBit(i)
	}
	// fmt.Println(str)
	toMod := int64(1e9 + 7)
	length := len(str)
	ans := int64(0)
	multi := int64(1)
	for i := length - 1; i >= 0; i-- {
		if str[i] == '1' {
			ans += multi
		} else {
			// do nothing
		}
		ans %= toMod
		multi *= 2
		multi %= toMod
	}
	fmt.Println("ans", ans)
	return int(ans)
}
func maxOperations(nums []int, k int) int {
	// n := len(nums)
	m := make(map[int]int)
	ans := 0
	for _, num := range nums {
		want := k - num
		if wantCount, ok := m[want]; ok && wantCount > 0 {
			m[want]--
			ans++
		} else {
			m[num]++
		}
	}
	// fmt.Println("ans", ans, m)
	return ans
}
func interpret(command string) string {
	n := len(command)
	ans := ""
	i := 0
	for i != n {
		if command[i] == 'G' {
			ans += "G"
			i++
		} else {
			if command[i+1] == ')' {
				ans += "o"
				i += 2
			} else {
				ans += "al"
				i += 4
			}
		}
	}
	fmt.Println("ans", ans)
	return ans
}
