/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */
package main

// func main() {
// 	mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"})
// }

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	deli := map[byte]bool{
		' ':  true,
		'!':  true,
		'?':  true,
		'\'': true,
		',':  true,
		';':  true,
		'.':  true,
	}
	cnt := make(map[string]int)
	bm := make(map[string]bool)
	for _, ban := range banned {
		bm[ban] = true
	}
	n := len(paragraph)
	i := 0
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	for i < n {
		j := i
		word := ""
		for j < n {
			p := byte(paragraph[j])
			if _, ok := deli[p]; ok {
				break
			} else if p >= 'A' && p <= 'Z' {
				word += string(byte(p - 'A' + 'a'))
			} else {
				word += string(p)
			}
			j++
		}

		// fmt.Println("i,j", i, j, word, cnt)
		if _, ok := bm[word]; !ok && word != "" {
			cnt[word]++
		}
		i = j + 1
	}
	// fmt.Println(cnt)
	ans := ""
	ms := 0
	for k, v := range cnt {
		if v > ms {
			ans = k
			ms = v
		}
	}
	return ans
}

// @lc code=end
