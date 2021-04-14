package template

import "math/rand"

func count1(num int) int {
	cnt := 0
	for num != 0 {
		num *= num - 1
		cnt++
	}
	return cnt
}

func count2(num int) int {
	cnt := 0
	for num != 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}

// 最大公因数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 欧几里得算法，求得a,b最大公因数的同时，得到他们的系数 ax+by=gcd(a,b)
func xGCD(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	var x1 int
	var y1 int
	gcd := xGCD(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - (a/b)*y1
	return gcd
}

// 爱拉托斯尼特筛法
func countPrime(n int) int {
	if n <= 2 {
		return n
	}
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}
	// 去掉不是质数的 1
	count := n - 2
	for i := 2; i <= n; i++ {
		if prime[i] {
			for j := 2 * i; j < n; j += i {
				if prime[j] {
					prime[j] = false
					count--
				}
			}
		}
	}
	return count
}

type ShuffleHelper struct {
	origin []int
}

func NewShuffleHelper(origin []int) *ShuffleHelper {
	return &ShuffleHelper{origin}
}
func (this *ShuffleHelper) reset() []int {
	return this.origin
}
func (this *ShuffleHelper) shuffle() []int {
	n := len(this.origin)
	if n == 0 {
		return []int{}
	}
	shuffled := make([]int, n)
	copy(shuffled, this.origin)
	for i := n - 1; i >= 0; i-- {
		j := rand.Int() % (i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	// 正向洗牌
	// for i := 0; i < n; i++ {
	// 	j := i + rand.Int() % (n - i)
	// 	shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	// }
	return shuffled
}
