package main

func main() {
	maxRepeating("ababc", "ab")
	maxRepeating("ababc", "ac")
	maxRepeating("ababc", "ba")
}

func maxRepeating(sequence string, word string) int {
	n := len(sequence)
	m := len(word)
	if n < m {
		return 0
	}
	i := 0
	cnt := 0
	ans := 0
	for i < n {
		j := i + m
		if j > n {
			break
		}
		// fmt.Println(sequence[i:j], word)
		if sequence[i:j] == word {
			cnt++
			if cnt > ans {
				ans = cnt
			}
			i = j
		} else {
			cnt = 0
			i++
		}
	}
	// fmt.Println("ans", ans)
	return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	slow, fast := list1, list1
	i := 0
	for i != a-1 {
		slow = slow.Next
		fast = fast.Next
		i++
	}
	for i != b {
		fast = fast.Next
		i++
	}
	fmt.Println(slow, fast)
	slow.Next = list2
	ptr := list2 
	for ptr.Next !=nil {
		ptr = ptr.Next
	}
	fmt.Println(ptr)
	ptr.Next = fast.Next
	return list1
}
