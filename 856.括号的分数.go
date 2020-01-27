/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 */

// @lc code=start
func scoreOfParentheses(S string) int {
    stack := make([]int, 0)
    ans := 0
    for i:=0;i<len(S);i++ {
        if S[i] =='(' {
            stack = append(stack, 0)
        } else {
            tmp := 0
            for {
                length := len(stack)
                if length == 0 {
                    // 不可能
                    break
                }
                top := stack[length-1]
                if top == 0 {
                    // 是一个左括号
                    stack = stack[:length-1]
                    if tmp != 0 {
                        tmp *= 2
                    } else {
                        tmp = 1
                    }
                    stack = append(stack, tmp)
                    break
                } else {
                    tmp += top
                    stack = stack[:length-1]
                }
            }
        }
        fmt.Println(stack)
    }
    fmt.Println(stack)
    for i:=0;i<len(stack);i++ {
		ans+=stack[i]
	}
    return ans
}
// @lc code=end

