import "strings"

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dict []string, sentence string) string {
	d := make(map[string]bool)
	for i := 0; i < len(dict); i++ {
		d[dict[i]] = true
	}
	var replace func(word string) string
	replace = func(word string) string {
		for j := 1; j < len(word); j++ {
			if _, ok := d[word[:j]]; ok {
				return word[:j]
			}
		}
		return word
	}
	list := strings.Split(sentence, " ")
	for i := 0; i < len(list); i++ {
		list[i] = replace(list[i])
	}
	return strings.Join(list, " ")
}

// @lc code=end
