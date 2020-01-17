/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 */

// @lc code=start
func repeatedStringMatch(A string, B string) int {
	lenA := len(A)
	lenB := len(B)

	k := lenB / lenA
	if lenB % lenA > 0 {
		k++
	}
	if strings.Contains(strings.Repeat(A, k), B) {
		return k
	}
	if strings.Contains(strings.Repeat(A, k + 1), B) {
		return k+1
	}
	return -1

	// // first := strings.Index(B, A)
	// // last := strings.LastIndex(B, A)
	// // 快慢指针找到重复的序列先
	// i:=0
	// j:=0
	// counter:=0
	// first:=0
	// last:=0
	// for i<len(B) {
	// 	if B[i] != A[j] {
	// 		i++
	// 	} else {
	// 		first=i
	// 		counter, last = helper(A, B, i)
	// 		break
	// 	}
	// }
	// fmt.Println(first, counter, last)
	// if counter == -1 {
	// 	return counter
	// }
	// // "abababaaba"
	// // "aabaaba"
	// // 0 0 0
	// // 这种特殊的情况在于 B 为 A*2 的字串
	// // if counter == 0 && first == 0 && last == 0 {
	// // 	if len(B) > len(A) {
	// // 		return -1
	// // 	}
	// // 	if strings.Contains(strings.Repeat(A))
	// // }
	// if first != 0 {
	// 	if !contains(A, B, 0, first) {
	// 		return -1
	// 	} else {
	// 		counter++
	// 	}
	// }
	// if last != len(B) {
	// 	if !contains(A, B, last, len(B)) {
	// 		return -1
	// 	} else {
	// 		counter++
	// 	}
	// }
	// return counter
}
func contains(A string, B string, start int, end int) bool {
	w := B[start:end]
	return strings.Contains(A, w)
}
func helper(A string, B string, start int) (c int, l int) {
	counter := 0
	lenA := len(A)
	i:=start
	for i < len(B) - lenA {
		for k:=0; k<lenA; k++ {
			if B[i+k] != A[k] {
				return -1, 0
			}
		}
		counter++
		i+=lenA
	}
	return counter, i
}
// @lc code=end

