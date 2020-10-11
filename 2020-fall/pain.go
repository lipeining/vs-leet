package fall

import "fmt"

func paintingPlan(n int, k int) int {
	if k == n*n {
		return 1
	}
	if k%n == 0 {
		// 只需要 行，或者 列就可以了
		return getCMN(n, k/n) * 2
	}
	// i, j 分别为 行和列 需要的数量
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			now := (i+j)*n - i*j
			if now == k {
				ans += getCMN(n, i) * getCMN(n, j)
				fmt.Println(i, j, ans)
			}
		}
	}
	return ans
}
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
func getCMN(n, m int) int {
	// if m == 0 {
	// 	return 0
	// }
	ans := 1
	for i := n - m + 1; i <= n; i++ {
		ans *= i
	}
	for i := 1; i <= m; i++ {
		ans /= i
	}
	return ans
}

// func getCmn(n, m int) int {
// 	// cmn = n!/((n-m)!*m!)
// 	if n == m {
// 		return 1
// 	}
// 	return multi(n) / (multi(n-m) * multi(m))
// }
// func multi(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	return n * multi(n-1)
// }
func main() {
	fmt.Println(getCMN(2, 0))
	// fmt.Println(getCMN(2, 1))
	// fmt.Println(getCMN(2, 2))
	// fmt.Println(getCMN(3, 2))
	// fmt.Println(getCMN(3, 1))
	// fmt.Println(getCMN(3, 3))
	// fmt.Println(getCMN(5, 2))
	// fmt.Println(getCMN(5, 4))
}
