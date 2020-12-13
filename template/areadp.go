package template

import (
	"fmt"
	"math"
)

// area

// //mst(dp,0) 初始化DP数组
// for(int i=1;i<=n;i++)
// {
//     dp[i][i]=初始值
// }
// for(int len=2;len<=n;len++)  //区间长度
// for(int i=1;i<=n;i++)        //枚举起点
// {
//     int j=i+len-1;           //区间终点
//     if(j>n) break;           //越界结束
//     for(int k=i;k<j;k++)     //枚举分割点，构造状态转移方程
//     {
//         dp[i][j]=max(dp[i][j],dp[i][k]+dp[k+1][j]+w[i][j]);
//     }
// }

func areaTemplate(w [][]int) {
	n := len(w)
	dp := make([][]int, n+1)
	initValue := 0
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		dp[i][i] = initValue
	}
	for length := 2; length <= n; length++ {
		for i := 1; i <= n; i++ {
			j := i + length - 1
			if j > n {
				break
			}
			for k := i; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k+1][j]+w[i][j])
			}
		}
	}
	fmt.Println(dp)
}

// 给出一个的只有'(',')','[',']'四种括号组成的字符串，求最多有多少个括号满足题目里所描述的完全匹配。
// 【思路】
// 用dp[i][j]表示区间[i,j]里最大完全匹配数。
// 只要得到了dp[i][j]，那么就可以得到dp[i-1][j+1]
// dp[i-1][j+1]=dp[i][j]+(s[i-1]与[j+1]匹配 ? 2 : 0)
// 然后利用状态转移方程更新一下区间最优解即可。
// dp[i][j]=max(dp[i][j],dp[i][k]+dp[k+1][j])

func leftrightpairs(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for length := 2; length <= n; length++ {
		for i := 1; i <= n; i++ {
			j := i + length - 1
			if j > n {
				break
			}
			if s[i] == '(' && s[j] == ')' || s[i] == '[' && s[j] == ']' {
				dp[i][j] = dp[i+1][j-1] + 2
			}
			for k := i; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j])
			}
		}
	}
	return dp[1][n]
	// 求需要使得字符串匹配的最少数量
	// return n-dp[1][n]
}

// 问题是我们经常见到的整数划分，给出两个整数 n , m ,
// 要求在 n 中加入m - 1 个乘号，将n分成m段，求出这m段的最大乘积

// 用dp[i][j]表示从第一位到第i位共插入j个乘号后乘积的最大值。
// 根据区间DP的思想我们可以从插入较少乘号的结果算出插入较多乘号的结果。
// dp[i][j]=max(dp[i][j],dp[k][j-1]*num[k+1][i])
// num[i][j]表示从s[i]到s[j]这段连续区间代表的数值。
func partNM(s string, m int) int {
	n := len(s)
	dp := make([][]int, n+1)
	byteMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	num := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m)
		num[i] = make([]int, m)
		num[i][i] = byteMap[s[i]]
		for j := i + 1; j <= n; j++ {
			num[i][j] = num[i][j-1]*10 + byteMap[s[j]]
		}
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = num[1][i]
	}
	for j := 1; j < m; j++ {
		for i := j + 1; i <= n; i++ {
			for k := j; k < i; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j-1]*num[k+1][i])
			}
		}
	}
	return dp[n][m-1]
}

// 给定一个具有N(N<=50)个顶点(从1到N编号)的凸多边形，
// 每个顶点的权值已知。问如何把这个凸多边形划分成N-2个互不相交的三角形，
// 使得这些三角形顶点的权值的乘积之和最小。
// dp[i][j]=min(dp[i][j],dp[i][k]+dp[k][j]+a[i]*a[k]*a[j]);
func triangle(a []int) int {
	n := len(a)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for length := 3; length <= n; length++ {
		for i := 1; i <= n; i++ {
			j := i + length - 1
			if j > n {
				break
			}
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]+a[i]*a[k]*a[j])
			}
		}
	}
	return dp[1][n]
}
