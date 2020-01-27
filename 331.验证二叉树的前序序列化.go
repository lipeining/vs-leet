/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	// list := strings.Split(preorder, ",")
	// for len(list) != 0 {
	// 	length := len(list)
	// 	if length < 3 {
	// 		return false
	// 	}
	// 	next := list[length - 1]
	// 	prev := list[length - 2]
	// 	if next == "#" && prev == "#" {
	// 		list = list[:length-2]
	// 	} else if next == "#" && prev != "#" {

	// 	} else if next != "#" && prev == "#" {

	// 	} else {

	// 	}
	// 	top := list[length - 3]
	// 	if top == "#" {
			
	// 	} else {
	// 		list = list[:length-3]
	// 	}		
	// 	list = append(list, "#")
	// 	fmt.Println(list)
	// }
	// return true
	list := strings.Split(preorder, ",")
	stack := make([]string, 0)
	for i:=0;i<len(list);i++ {
		stack = append(stack, list[i])
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" {
			stack = stack[:len(stack)-2]
			if stack[len(stack)-1] == "#" {
				return false
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, "#")
		}
	}
	return len(stack) == 1 && stack[0] == "#"
}
// @lc code=end

