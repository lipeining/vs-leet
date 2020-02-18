import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] == arr[j] {
			return false
		}
		return math.Abs(float64(arr[i]-x)) <= math.Abs(float64(arr[j]-x))
	})
	ans := make([]int, k)
	copy(ans, arr[:k])
	sort.Ints(ans)
	return ans
}

// func findClosestElements(arr []int, k int, x int) []int {
// 	index := binary(arr, x)
// 	length := len(arr)
// 	left, right := index-1, index+1
// 	for i := 1; i < k; i++ {
// 		wantLeft := x - 10001
// 		if left >= 0 {
// 			wantLeft = arr[left]
// 		}
// 		wantRight := x + 10001
// 		if right < length {
// 			wantRight = arr[right]
// 		}
// 		if pick(wantLeft, wantRight, x) {
// 			// right
// 			right++
// 		} else {
// 			// left
// 			left--
// 		}
// 		// fmt.Println(left, right)
// 	}
// 	fmt.Println(index, length, left, right)
// 	return arr[left+1:right]
// }
func pick(wantLeft, wantRight, x int) bool {
	l := x - wantLeft
	r := wantRight - x
	if l <= r {
		return false
	}
	return true
}
func binary(arr []int, x int) int {
	length := len(arr)
	left, right := 0, length-1
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] > x {
			right = mid - 1
		} else if arr[mid] < x {
			left = mid + 1
		} else {
			return mid
		}
	}
	// if arr[left] == x {
	// 	return left
	// }
	return left
}

// @lc code=end
