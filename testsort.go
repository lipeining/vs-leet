package main

import "fmt"

func qsort(arr []int) {
	partition := func(low, high int) int {
		privot := arr[low]
		for low < high {
			for low < high && arr[high] >= privot {
				high--
			}
			arr[low], arr[high] = arr[high], arr[low]
			for low < high && arr[low] <= privot {
				low++
			}
			arr[low], arr[high] = arr[high], arr[low]
		}
		return low
	}
	var dfs func(left int, right int)
	dfs = func(left int, right int) {
		if left >= right {
			return
		}
		p := partition(left, right)
		fmt.Println("now p:", p, arr)
		dfs(left, p-1)
		dfs(p+1, right)
	}
	fmt.Println("before ", arr)
	dfs(0, len(arr)-1)
	fmt.Println("after ", arr)
}

func main() {
	// qsort([]int{1, 2, 3, 4, 5, 6, 7})
	// qsort([]int{1, 2, 3, 4, 5, 5, 6, 7})
	qsort([]int{7, 6, 6, 5, 5, 4, 1, 2, 3})
	// qsort([]int{7, 6, 8, 5, 5, 4, 1, 2, 3})
}
