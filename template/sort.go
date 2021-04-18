package template

import "fmt"

func qsort(arr []int) {
	partition := func(left, right int) int {
		privot := arr[left]
		l := left + 1
		r := right
		for l < r {
			for l < r && arr[l] <= privot {
				l++
			}
			for l < r && arr[r] >= privot {
				r--
			}
			if l >= r {
				break
			}
			arr[l], arr[r] = arr[r], arr[l]
		}
		arr[left], arr[l] = arr[l], arr[left]
		return l
	}
	var dfs func(left int, right int)
	dfs = func(left int, right int) {
		if left >= right {
			return
		}
		p := partition(left, right)
		fmt.Println("now p:", p)
		dfs(left, p-1)
		dfs(p+1, right)
	}
	fmt.Println("before ", arr)
	dfs(0, len(arr)-1)
	fmt.Println("after ", arr)
}
