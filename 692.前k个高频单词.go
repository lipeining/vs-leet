/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int)
	for i:=0;i<len(words);i++ {
		counter[words[i]]++
	}
	list := make([]Node, len(counter))
	i:=0
	for k,v := range counter {
		list[i] = Node{k,v}
		i++
	}
	sort.Slice(list, func(i int, j int)bool {
		t := list[i].count - list[j].count
		if t!= 0 {
			return t > 0
		}
		return list[i].val < list[j].val
	})
	fmt.Println(counter)
	fmt.Println(list)
	ans := make([]string, k)
	for i:=0;i<k;i++ {
		ans[i] = list[i].val
	}
	return ans
}
type Node struct {
	val string
	count int
}
// @lc code=end

