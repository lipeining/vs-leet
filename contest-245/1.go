package main

func main() {

}
func mergeTriplets(triplets [][]int, target []int) bool {
	n := len(triplets)
	first := make(map[int]bool)
	second := make(map[int]bool)
	third := make(map[int]bool)
	f, s, t := false, false, false
	for i := 0; i < n; i++ {
		triplet := triplets[i]
		if triplet[0] == target[0] && triplet[1] == target[1] && triplet[2] == target[2] {
			return true
		}
		if triplet[0] == target[0] {
			f = true
		}
		if triplet[1] == target[1] {
			s = true
		}
		if triplet[2] == target[2] {
			t = true
		}
		if triplet[0] <= target[0] {
			first[i] = true
		}
		if triplet[1] <= target[1] {
			second[i] = true
		}
		if triplet[2] <= target[2] {
			third[i] = true
		}
	}
	if !f || !s || !t {
		return false
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 看三个的交集是否存在
	exist := func(a, b, c map[int]bool) bool {
		for index := range a {
			if _, ok2 := b[index]; ok2 {
				if _, ok3 := c[index]; ok3 {
					return true
				}
			}
		}
		return false
	}
	start := min(min(len(first), len(second)), len(third))
	ans := false
	if start == len(first) {
		ans = exist(first, second, third)
	} else if start == len(second) {
		ans = exist(second, first, third)
	} else {
		ans = exist(third, first, second)
	}
	return ans
}
func maximumRemovals(s string, p string, removable []int) int {
	ns := len(s)
	np := len(p)
	n := len(removable)
	check := func(k int) bool {
		state := make([]bool, ns)
		for i := 0; i < k; i++ {
			state[removable[i]] = true
		}
		j := 0
		for i := 0; i < ns; i++ {
			if state[i] {
				continue
			}
			if s[i] == p[j] {
				j++
			}
			if j == np {
				return true
			}
		}
		return false
	}
	l := 0
	r := n
	for l < r {
		k := (l + r + 1) >> 1
		if check(k) {
			l = k
		} else {
			r = k - 1
		}
	}
	return l
}
func makeEqual(words []string) bool {
	cnt := make(map[byte]int)
	for _, word := range words {
		for j := 0; j < len(word); j++ {
			cnt[word[j]]++
		}
	}
	n := len(words)
	for _, v := range cnt {
		if v%n != 0 {
			return false
		}
	}
	return true
}
