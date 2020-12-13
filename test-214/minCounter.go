package main

import "fmt"

// func main() {
// 	// "abcabc"
// 	//     "aaabbbcc"
// 	minDeletions("aab")
// 	minDeletions("abcabc")
// 	minDeletions("aaabbbcc")
// 	minDeletions("ceabaacb")
// }
func minDeletions(s string) int {
	counter := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	// 得到相同频次的列表
	same := make(map[int]int)
	start := 0
	for _, v := range counter {
		same[v]++
		start = max(start, v)
	}
	fmt.Println(counter, same, start)
	ans := 0
	for num := start; num > 0; num-- {
		seque := same[num]
		if seque > 1 {
			ans += seque - 1
			same[num-1] += seque - 1
		}
	}
	fmt.Println("ans", ans)
	fmt.Println("-----------")
	return ans
}
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
