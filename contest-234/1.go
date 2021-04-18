package main

import (
	"fmt"
)

func main() {
	// func1()
	// func2()
	func3()
}
func maxNiceDivisors(primeFactors int) int {
	// prime 应该都取值 2,3,5,7,11,13 如果取到了 13 那么，有用吗？
	// 区间 dp ，假定现在有 i 个质数，里面不同的个数为 j
	// 那么，每一种都要至少 1 个。
	mod := int(1e9 + 7)
	return mod
}
func func3() {
	evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}})
	evaluate("hi(name)", [][]string{{"a", "b"}})
	evaluate("(a)(a)(a)aaa", [][]string{{"a", "yes"}})
	evaluate("(a)(b)", [][]string{{"a", "b"}, {"b", "a"}})
}
func evaluate(s string, knowledge [][]string) string {
	KM := make(map[string]string)
	for _, know := range knowledge {
		KM[know[0]] = know[1]
	}
	fmt.Println(KM)
	i := 0
	n := len(s)
	ans := ""
	for i < n {
		if s[i] == '(' {
			j := i + 1
			for j < n {
				if s[j] == ')' {
					break
				} else {
					j++
				}
			}
			key := s[i+1 : j]
			fmt.Println("the key is ", key)
			now := "?"
			if value, ok := KM[key]; ok {
				now = value
			}
			ans += now
			i = j + 1
		} else {
			ans += string(s[i])
			i++
		}
	}
	fmt.Println("ansss", ans)
	return ans
}
func func2() {
	reinitializePermutation(2)
	reinitializePermutation(4)
	reinitializePermutation(6)
}
func reinitializePermutation(n int) int {
	perm := make([]int, n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	ans := 1
	check := func() bool {
		for i := 0; i < n; i++ {
			if i != perm[i] {
				return false
			}
		}
		return true
	}
	action := func() {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		copy(perm, arr)
	}
	action()
	for !check() {
		ans++
		action()
	}
	fmt.Println("anssss", ans)
	return ans
}
func func1() {
	numDifferentIntegers("a123bc34d8ef34")
	numDifferentIntegers("leet1234code234")
	numDifferentIntegers("a1b01c001")
	numDifferentIntegers("a1b0c001")
	numDifferentIntegers("abc")
	numDifferentIntegers("abc0")
	numDifferentIntegers("0abc")
	numDifferentIntegers("uu717761265362523668772526716127260222200144937319826j717761265362523668772526716127260222200144937319826k717761265362523668772526716127260222200144937319826b7177612653625236687725267161272602222001449373198262hgb9utohfvfrxprqva3y5cdfdudf7zuh64mobfr6mhy17")
}
func numDifferentIntegers(word string) int {
	m := make(map[string]bool)
	n := len(word)
	i := 0
	for i < n {
		if word[i] >= '0' && word[i] <= '9' {
			j := i + 1
			for j < n {
				if word[j] >= '0' && word[j] <= '9' {
					j++
				} else {
					break
				}
			}
			if j != i+1 {
				for i < j {
					if word[i] == '0' {
						i++
					} else {
						break
					}
				}
			}
			num := word[i:j]
			m[num] = true
			i = j
		} else {
			i++
		}
	}
	fmt.Println("anssss", len(m), m)
	return len(m)
}
