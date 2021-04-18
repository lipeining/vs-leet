package main

import (
	"fmt"
)

// func main() {
// 	//a := 0
// 	//fmt.Scan(&a)
// 	//fmt.Printf("%d\n", a)
// 	fmt.Printf("Hello World!\n")
// 	reverse("hello world Hi!")
// }
func reverse(s string) string {
	n := len(s)
	i := 0
	ans := ""
	for i < n {
		if s[i] >= 'a' && s[i] <= 'z' {
			j := i + 1
			for j < n {
				if s[j] >= 'a' && s[j] <= 'z' {
					j++
				} else {
					break
				}
			}
			for k := j - 1; k >= i; k-- {
				ans += string(s[k])
			}
			i = j
		} else {
			ans += string(s[i])
			i++
		}
	}
	return ans
}
func findKthNumber(n, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	getCount := func(prefix int) int {
		count := 0
		cur := prefix
		next := prefix + 1
		for cur <= n {
			count += min(n+1, next) - cur
			cur *= 10
			next *= 10
		}
		return count
	}
	p := 1
	prefix := 1
	for p < k {
		count := getCount(prefix)
		if p+count > k {
			prefix *= 10
			p++
		} else {
			prefix++
			p += count
		}
	}
	return prefix
}

// 设计一个千万的排行榜
// 查看一些资料字后，觉得这样做比较好。
// 需要考虑时间性
// 使用 redis.zset
// 1.通过 kafka 的 consumer 离线计算对应的时间点，比如以，小时，15分钟划分的间隔的
// 视频的排行榜，可以汇总到 topk * n 的数据。
// 不管如何，最后，会得到一个 zset
// 需要通过 YYYY-MM-DD HH:MM 的一个间隔的排行榜
// 那么，可以得到 A-B 这样的一个时间戳为  key 的 zset
// 这个 zset 会记录视频和对应的数据 value，并且可以排序
// 那么，如果得到小时的，需要 zuion 这个小时的数据的 zset 得到一个大的排行榜。
// 如果是计算一天，一周，一个月，可以分步进行，避免大集合的操作导致时间超长。
// https://www.cnblogs.com/thisiswhy/p/14470861.html
