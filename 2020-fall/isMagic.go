package fall

import "fmt"

func isMagic(target []int) bool {
	n := len(target)
	base := make([]int, n)
	for i := 0; i < n; i++ {
		base[i] = i + 1
	}
	init := flash(base)
	k := 0
	for i := 0; i < n; i++ {
		if target[i] != init[i] {
			break
		}
		k = i + 1
	}
	if k == 0 {
		return false
	}
	// 检查这个 k 是否符合之后的运算
	return check(init[k:], target[k:], k)
}
func check(list, target []int, k int) bool {
	nextList := flash(list)
	fmt.Println(list, target, k)
	for i := 0; i < k && i < len(target); i++ {
		if nextList[i] != target[i] {
			return false
		}
	}
	if len(target) < k {
		return true
	}
	return check(nextList[k:], target[k:], k)
}
func flash(arr []int) []int {
	n := len(arr)
	ans := make([]int, len(arr))
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			ans[n/2+i/2] = arr[i]
		} else {
			ans[i/2] = arr[i]
		}
	}
	return ans
}
func getInit(n int) []int {
	ans := make([]int, n)
	for i := 0; i < n/2; i++ {
		ans[i] = (i + 1) * 2
		ans[i+n/2] = i*2 + 1
	}
	if n%2 == 1 {
		ans[n-1] = n
	}
	return ans
}

// func main() {
// 	// flash([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
// 	// flash([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
// 	fmt.Println(isMagic([]int{2, 4, 3, 1, 5}))
// 	fmt.Println(isMagic([]int{5,4,3,2,1}))
// }
