package main

import "fmt"

func main() {
	fmt.Println(string([]byte{'1', '0'}))
	// fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	// fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
	// fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 1, 1, 1, 1}))
	// 	输入：customers = [[1,2],[2,5],[4,3]]
	// 	输出：5.00000
	// 	输入：customers = [[5,2],[5,4],[10,3],[20,1]]
	// 输出：3.25000
	// 解释：
	// fmt.Println(averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}}))
	// fmt.Println(averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}))
	// [[2,3],[6,3],[7,5],[11,3],[15,2],[18,1]]
	//4.16667
	// fmt.Println(averageWaitingTime([][]int{{2, 3}, {6, 3}, {7, 5}, {11, 3}, {15, 2}, {18, 1}}))
}
func maximumBinaryString(binary string) string {
	// 将 10 当做特别的字符串，又从到左，先将 10->01维持多一些 前缀00
	// 但是最开始的1是不会操作的。
	// 然后从左到右，00->10进行一次后，
	// 再从左到右，01->返回来
	n := len(binary)
	str := make([]byte, n)
	for i := 0; i < n; i++ {
		str[i] = binary[i]
	}
	pre1 := 0
	for i := 0; i < n; i++ {
		if binary[i] == '1' {
			pre1++
		} else {
			break
		}
	}
	j := n - 2
	for j > pre1 {
		if binary[j] == '1' && binary[j+1] == '0' {
			str[j] = '0'
			str[j+1] = '1'
			j -= 2
		} else {
			j--
		}
	}
	i := pre1
	for i < n {
		if str[i] == '0' && i < n-1 && str[i+1] == '1' {
			str[i] = '1'
			str[i+1] = '0'
			i += 2
		} else {
			i++
		}
	}
	k:=pre1
	for k<n {
		if str[k] == '0' && k < n-1&& str[k+1] == '1' {
			if binary[k] == '1' {
				
			}
		}
	}
	return ""
}

func countStudents(students []int, sandwiches []int) int {
	n := len(students)
	mark := make([]int, n)
	check := func(sandwich int, j int) int {
		end := j + n
		for j < end {
			if mark[j%n] == 0 {
				if students[j%n] == sandwich {
					return j
				}
			} else if mark[j%n] == 1 {

			} else if mark[j%n] == 2 {
				if students[j%n] == sandwich {
					return j
				}
			}
			j++
		}
		return j
	}
	lastJ := 0
	for i := 0; i < n; i++ {
		wantJ := check(sandwiches[i], lastJ)
		// fmt.Println("lastJ", lastJ, "wantJ", wantJ)
		if wantJ != lastJ+n {
			lastJ = wantJ
			mark[wantJ%n] = 1
		} else {
			// fmt.Println("i", i)
			return n - i
		}
	}
	return 0
}
func countStudents2(students []int, sandwiches []int) int {
	zero := 0
	one := 0
	n := len(students)
	for i := 0; i < len(students); i++ {
		if students[i] == 0 {
			zero++
		} else {
			one++
		}
	}
	if zero == one {
		return 0
	}
	if zero < one {
		// 找到可以消化掉的One的三明治的数量
		for i := 0; i < n; i++ {
			if sandwiches[i] == 0 {
				if zero > 0 {
					zero--
				} else {
					// fmt.Println("i", i)
					return n - i
				}
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if sandwiches[i] == 1 {
				if one > 0 {
					one--
				} else {
					// fmt.Println("i", i)
					return n - i
				}
			}
		}
	}
	return 0
}
func averageWaitingTime(customers [][]int) float64 {
	n := len(customers)
	if n == 0 {
		return 0
	}
	cooker := customers[0][0]
	var total float64
	for i := 0; i < n; i++ {
		curStart := customers[i][0]
		if cooker <= curStart {
			cooker = curStart
		}
		curEnd := cooker + customers[i][1]
		cur := curEnd - curStart
		fmt.Println("start, end, cur, cooker", curStart, curEnd, cur, cooker)
		cooker = curEnd
		total += float64(cur)
	}
	fmt.Println("total", total)
	return total / float64(n)
}
