/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 */

// @lc code=start
func reorganizeString(S string) string {
	counter := make(map[string]int)
	for i:=0;i<len(S);i++ {
		counter[string(S[i])]++
	}
	list := make([]Node, len(counter))
	i:=0
	for k,v := range counter {
		list[i] = Node{k,v}
		i++
	}
	sort.Slice(list, func(i int, j int)bool {
		return list[i].count > list[j].count
	})
	fmt.Println(counter)
	fmt.Println(list)
	ans := make([]string, len(S))
	start := 0
	for _,v := range list {
		val,count := v.val, v.count
		max := len(S) / 2
		if len(S) % 2 == 1 {
			max++
		}
 		if count > max {
			return ""
		}
		for count > 0 {
			ans[start] = val
			start += 2
			count--
			if start >= len(S) {
				start = 1
			}
		}
	}
	return strings.Join(ans, "")
}
type Node struct {
	val string
	count int
}
// @lc code=end

