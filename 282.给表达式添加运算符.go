/*
 * @lc app=leetcode.cn id=282 lang=golang
 *
 * [282] 给表达式添加运算符
 */
package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	addOperators("123", 6)
// }
// https://leetcode-cn.com/problems/expression-add-operators/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-52/

// @lc code=start
func addOperators(num string, target int) []string {
	// n := len(num)
	ans := make([]string, 0)
	// 左右两部分
	type pair struct {
		path   string
		number int
	}
	// 无法处理 乘法 的优先级问题
	var cacl func(input string) []pair
	cacl = func(input string) []pair {
		list := make([]pair, 0)
		size := len(input)
		number, _ := strconv.Atoi(input)
		list = append(list, pair{
			path:   input,
			number: number,
		})
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				lnum := input[i:j]
				rnum := input[j:size]
				left := cacl(lnum)
				right := cacl(rnum)
				fmt.Println("lnum", lnum, left)
				fmt.Println("rnum", rnum, right)
				for _, l := range left {
					for _, r := range right {
						list = append(list, pair{
							path:   l.path + "+" + r.path,
							number: l.number + r.number,
						})
						list = append(list, pair{
							path:   l.path + "-" + r.path,
							number: l.number - r.number,
						})
						list = append(list, pair{
							path:   l.path + "*" + r.path,
							number: l.number * r.number,
						})
					}
				}
			}
		}
		fmt.Println("for input", input, list)
		return list
	}
	list := cacl(num)
	for _, p := range list {
		if p.number == target {
			ans = append(ans, p.path)
		}
	}
	// fmt.Println(list)
	fmt.Println(ans)
	return ans
}

// @lc code=end
