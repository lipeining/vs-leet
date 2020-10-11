package template

// for (int j = 0; j < s.size(); j++) {
// 	窗口右端扩展，加进s[j], 更新条件
// 	while(不满足条件) {
// 		窗口左端移除s[i]，更新条件，然后i++
// 	}
// 	此时重新满足条件，和最优比较并记录
// }

// func solve(nums []int) {
// 	i := 0
// 	for j := 0; j < n; j++ {
// 		// add nums[j]
// 		for {
// 			// remove nums[i]
// 			i++
// 			// break
// 		}
// 		// compare ans
// 	}
// }

// 1 最短摘要：求s种包含t中所有字符的最短子串
func SIncludesT(s, t string) string {
	i := 0
	tMap := make(map[byte]int)
	// for _, char := range t {
	// 	tMap[char]++
	// }
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	sMap := make(map[byte]int)
	num := 0
	minStart := -1
	minLen := len(s) + 1
	for j := 0; j < len(s); j++ {
		if _, ok := tMap[s[j]]; !ok {
			continue
		}
		if sMap[s[j]] < tMap[s[j]] {
			num++
		}
		sMap[s[j]]++
		for num == len(t) {
			if j-i+1 < minLen {
				minLen = j - i + 1
				minStart = i
			}
			sMap[s[i]]--
			if tMap[s[i]] > 0 && sMap[s[i]] < tMap[s[i]] {
				num--
			}
			i++
		}
	}
	if minStart == -1 {
		return ""
	}
	return s[minStart : minStart+minLen+1]
}

// 2 和大于给定值的最小子数组，（均为正数）
func minSubArrayLen(nums []int, target int) int {
	i := 0
	sum := 0
	n := len(nums)
	minLen := n + 1
	for j := 0; j < n; j++ {
		sum += nums[j]
		for sum >= target {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}
	if minLen == n+1 {
		return -1
	}
	return minLen
}

// 3最长无重复字符子串：
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := make(map[byte]int)
	maxLen := 0
	start := 0
	for j := 0; j < n; j++ {
		if last, ok := m[s[j]]; ok {
			start = last + 1
		}
		m[s[j]] = j
		if j-start+1 > maxLen {
			maxLen = j - start + 1
		}
	}
	return maxLen
}

// 4最长的只有k个不同字符的子串：
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	n := len(s)
	maxLen, start := 0, 0
	number := 0
	m := make(map[byte]int)
	for j := 0; j < n; j++ {
		last, ok := m[s[j]]
		if !ok || last < start {
			number++
		}
		m[s[j]] = j
		for number > k {
			if m[s[start]] == start {
				number--
			}
			start++
		}
		if j-start+1 > maxLen {
			maxLen = j - start + 1
		}
	}
	return maxLen
}
