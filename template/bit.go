package template

// allZero 表示全部是 0 的二进制数据
// allOne 表示全部是 1 的二进制数据
// x^allZero=x x&allZero=0  x|allZero=x
// x^allOne=~x x&allOne=x   x|allOne=allOne
// x^x=0       x&x=x        x|x=x

// 求 x y 的汉明距离
func hanmmingDistance(x, y int) int {
	diff := x ^ y
	ans := 0
	// 异或，然后统计 1 的数量
	for diff != 0 {
		ans += diff & 1
		diff >>= 1
	}
	return ans
}

// 求 n 二进制反转结果
func reverseBit(n uint) uint {
	var ans uint
	for i := 0; i < 32; i++ {
		ans <<= 1
		ans += n & 1
		n >>= 1
	}
	return ans
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

// countBits  每一个数字的二进制有多少个 1
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i&1 != 0 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i>>1]
		}
	}
	return dp
}
