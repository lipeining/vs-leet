/*
 * @lc app=leetcode.cn id=1106 lang=golang
 *
 * [1106] 解析布尔表达式
 */
package main

import "fmt"

// func main() {
// 	parseBoolExpr("!(f)")
// 	parseBoolExpr("|(f,t)")
// 	parseBoolExpr("&(t,f)")
// 	parseBoolExpr("|(&(t,f,t),!(t))")
// }

// @lc code=start
func parseBoolExpr(expression string) bool {
	n := len(expression)
	exp := make([]byte, 0)
	op := make([]byte, 0)
	// print := func(str []byte) {
	// 	for i := 0; i < len(str); i++ {
	// 		fmt.Print(string(str[i]))
	// 	}
	// }
	cacl := func() {
		// fmt.Println("debug exp op----------start--------")
		// print(exp)
		// fmt.Println("")
		// print(op)
		// fmt.Println("")
		// fmt.Println("debug exp op----------end--------")
		list := make([]bool, 0)
		var operator byte
		if len(op) != 0 {
			operator = op[len(op)-1]
			op = op[:len(op)-1]
		}
		for len(exp) != 0 {
			e := exp[len(exp)-1]
			exp = exp[:len(exp)-1]
			if e == '(' {
				break
			}
			flag := false
			if e == 't' {
				flag = true
			}
			list = append(list, flag)
		}
		want := true
		// fmt.Println("operator", string(operator))
		// fmt.Println("list", list)
		if operator == '!' {
			want = !list[0]
		} else if operator == '&' {
			want = true
			for i := 0; i < len(list); i++ {
				want = want && list[i]
			}
		} else if operator == '|' {
			want = false
			for i := 0; i < len(list); i++ {
				want = want || list[i]
			}
		} else {
			// 未知运算符，忽略
			want = list[0]
		}
		trans := byte('f')
		if want {
			trans = byte('t')
		}
		exp = append(exp, trans)
	}
	for i := 0; i < n; i++ {
		char := expression[i]
		if char == ')' {
			cacl()
		} else {
			if char == '!' || char == '&' || char == '|' {
				op = append(op, char)
			} else {
				if char == ',' {
					continue
				}
				exp = append(exp, char)
			}
		}
	}
	// 计算剩下的表达式
	cacl()
	ans := false
	if exp[0] == 't' {
		ans = true
	}
	fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
