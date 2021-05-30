/*
 * @lc app=leetcode.cn id=726 lang=golang
 *
 * [726] 原子的数量
 */
package main

import (
	"sort"
	"strconv"
)

// func main() {
// 	countOfAtoms("(H)")
// 	countOfAtoms("(H2O)")
// 	countOfAtoms("H2O")
// 	countOfAtoms("Mg(OH)2")
// 	countOfAtoms("K4(ON(SO3)2)2")
// }

// @lc code=start
func countOfAtoms(formula string) string {
	type node struct {
		atom string
		num  int
	}
	stack := make([]node, 0)
	n := len(formula)
	ans := ""
	i := 0
	readAtom := func(i int) (string, int) {
		ret := string(formula[i])
		j := i + 1
		for j < n {
			if formula[j] >= 'a' && formula[j] <= 'z' {
				ret += string(formula[j])
				j++
			} else {
				break
			}
		}
		return ret, j
	}
	readNum := func(i int) (int, int) {
		j := i
		num := 0
		for j < n {
			if formula[j] >= '0' && formula[j] <= '9' {
				num = num*10 + int(formula[j]-'0')
				j++
			} else {
				break
			}
		}
		if num == 0 {
			num = 1
		}
		return num, j
	}
	for i < n {
		f := formula[i]
		if f == '(' {
			stack = append(stack, node{atom: "(", num: 1})
			i++
		} else if f == ')' {
			i++
			num, j := readNum(i)
			next := make([]node, 0)
			for {
				size := len(stack)
				if size == 0 {
					break
				}
				top := stack[size-1]
				stack = stack[:size-1]
				if top.atom == "(" {
					break
				}
				top.num *= num
				next = append(next, top)
			}
			for _, node := range next {
				stack = append(stack, node)
			}
			i = j
		} else {
			atom, j := readAtom(i)
			num, j := readNum(j)
			stack = append(stack, node{atom: atom, num: num})
			i = j
		}
	}
	// fmt.Println("cur---------------------", stack)
	cnt := make(map[string]int)
	for i := 0; i < len(stack); i++ {
		atom := stack[i].atom
		num := stack[i].num
		if atom != " " {
			cnt[atom] += num
		}
	}
	keys := make([]string, len(cnt))
	for k := range cnt {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// fmt.Println(keys, cnt)
	for _, k := range keys {
		if cnt[k] > 0 {
			ans += k
			if cnt[k] > 1 {
				ans += strconv.Itoa(cnt[k])
			}
		}
	}
	// fmt.Println(ans)
	return ans
}

// @lc code=end
