package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	i := 0
	for i < len(words) {
		j := i
		width := 0
		space := 0
		for j < len(words) {
			width += len(words[j])
			if j != i {
				// 第二个之后的需要有空格
				space++
			}
			if width+space > maxWidth {
				width -= len(words[j])
				break
			}
			j++
		}
		// if j == i+1 {
		//     ans = append(ans, words[i] + strings.Repeat(" ", maxWidth-len(words[i])))
		//     i = j
		//     continue
		// }
		if j == i+1 || j == len(words) {
			// 最后一行
			cur := ""
			for k := i; k < j; k++ {
				cur += words[k]
				if k != j-1 {
					cur += " "
				}
			}
			// fmt.Println(cur)
			if len(cur) < maxWidth {
				cur = cur + strings.Repeat(" ", maxWidth-len(cur))
			}
			ans = append(ans, cur)
			i = j
			continue
		}
		cur := ""
		fixSpace := maxWidth - width
		eachSpace := fixSpace / (j - i - 1)
		moreSpace := fixSpace % (j - i - 1)
		fmt.Println(eachSpace, moreSpace)
		for k := i; k < j; k++ {
			cur += words[k]
			if k != j-1 {
				mySapce := eachSpace
				if moreSpace > 0 {
					mySapce++
					moreSpace--
				}
				if fixSpace < eachSpace {
					mySapce = fixSpace
				}
				// fmt.Println(mySapce)
				cur += strings.Repeat(" ", mySapce)
				fixSpace -= mySapce
			}
		}
		fmt.Println(cur)
		ans = append(ans, cur)
		// i 从 j 开始
		i = j
	}
	return ans
}

// func main() {
// 	tests := []struct {
// 		arr    []string
// 		target int
// 	}{
// 		{arr: []string{"This", "is", "an", "example", "of", "text", "justification."}, target: 16},
// 		{arr: []string{"What", "must", "be", "acknowledgment", "shall", "be"}, target: 16},
// 		{arr: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, target: 20},
// 	}
// 	for _, t := range tests {
// 		fmt.Println(fullJustify(t.arr, t.target))
// 	}
// }

// ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]
// 20
// 输出：
// ["Science   is what we","understand      well","enough to explain to","a   computer. Art is","everything  else  we","do                  "]
// 预期结果：
// ["Science  is  what we","understand      well","enough to explain to","a  computer.  Art is","everything  else  we","do                  "]
