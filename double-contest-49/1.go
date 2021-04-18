package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	countNicePairs([]int{42, 11, 1, 97})
	countNicePairs([]int{13, 10, 35, 24, 76})
	// areSentencesSimilar("My name is Haley", "My Haley")
	// areSentencesSimilar("of", "A lot of words")
	// areSentencesSimilar("Eating right now", "Eating")
	// areSentencesSimilar("Luky", "Luk343y")
	// areSentencesSimilar("Luky", "Luk343y adfd")

	// areSentencesSimilar("Y ggUFOmtf woKuTtO W uwJZ Zan wgm zprl Kgn mAY xLlCH phA UIVKIohfw al g m", "Jfa jfvmGU bKSSX uQ AmTzbBW EF jdc ft Z g VcM oNlI jeX q mNG YnUgGSnejt Y")
}
func countNicePairs(nums []int) int {
	// nlogn 的算法
	// nums[i] - rev[nums[i]] == nums[j] - rev(nums[j])
	rev := func(num int) int {
		s := strconv.Itoa(num)
		rs := ""
		for i := len(s) - 1; i >= 0; i-- {
			rs += string(s[i])
		}
		rn, _ := strconv.Atoi(rs)
		return rn
	}
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sub := nums[i] - rev(nums[i])
		counter[sub]++
	}
	fmt.Println(counter)
	ans := 0
	mod := int(1e9 + 7)
	for _, v := range counter {
		if v >= 2 {
			ans += v * (v - 1) / 2
			ans %= mod
		}
	}
	fmt.Println(ans)
	return ans
}
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	// 右边为比较大的
	l := 0
	r := 0
	fmt.Println(s1, s2)
	for l < len(s1) && r < len(s2) && s1[l] == s2[r] {
		l++
		r++
	}
	if l == len(s1) {
		fmt.Println("s1 enenenenene", "true")
		return true
	}
	for r < len(s2) {
		if s1[l] == s2[r] {
			break
		} else {
			r++
		}
	}
	if r == len(s2) {
		return false
	}
	l++
	r++
	if l == len(s1) && r != len(s2) {
		return false
	}
	if l != len(s1) && r == len(s2) {
		return false
	}
	for l < len(s1) && r < len(s2) {
		if s1[l] != s2[r] {
			return false
		}
		l++
		r++
	}
	fmt.Println("true")
	return true
}
func squareIsWhite(coordinates string) bool {
	char := int(coordinates[0] - 'a')
	num := int(coordinates[1] - '0')
	if char%2 == 0 {
		return num%2 == 0
	}
	return num%2 == 1
}
