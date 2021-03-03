// 问题等价于: 在一个矩阵中选取一些值, 满足矩阵的每一行和每一列都至少有一个元素被选中, 同时选中元素的总和最小 (此矩阵就是 cost 矩阵).

// 由于矩阵的列数较少, 我们可以用状压 DP 来表示每一行的选取情况, 假设矩阵有 mm 行 nn 列, 那么我们维护一个 DP 矩阵 dp[m][1 << n], dp[i][j]表示当前选取到第 ii 行, 每列的选取状况为 jj 时总的最小开销, 其中 jj 的第 kk 位为 11 即表示第 kk 列已经被选取过了. 那么状态转移方程为

// dp[i][j~|~k] = Math.min(dp[i][j~|~k], dp[i - 1][k] + costMatrix[i][j])
// dp[i][j ∣ k]=Math.min(dp[i][j ∣ k],dp[i−1][k]+costMatrix[i][j])

// 其中 costMatrix[i][j] 表示第 ii 行选取状况为 jj 时该行被选取得元素总和.


// class Solution {
//   public int connectTwoGroups(List<List<Integer>> cost) {
//     int m = cost.size(), n = cost.get(0).size();
//     int[][] costMatrix = new int[m][1 << n];
//     for (int k = 0; k < m; k++) {
//       for (int i = 0; i < (1 << n); i++) {
//         int sum = 0;
//         for (int j = 0; j < n; j++) {
//           if ((i & (1 << j)) > 0)
//             sum += cost.get(k).get(j);
//         }
//         costMatrix[k][i] = sum;
//       }
//     }
//     int[][] dp = new int[m][1 << n];
//     for (int i = 1; i < m; i++)
//       Arrays.fill(dp[i], Integer.MAX_VALUE);
//     dp[0] = costMatrix[0];
//     for (int i = 1; i < m; i++)
//       for (int j = 1; j < (1 << n); j++)
//         for (int k = 1; k < (1 << n); k++)
//           dp[i][j | k] = Math.min(dp[i][j | k], dp[i - 1][k] + costMatrix[i][j]);
//     return dp[m - 1][(1 << n) - 1];
//   }
// }
// 最终结果为 dp[m - 1][(1 << n) - 1], 表示选取到 m - 1 行 (即最后一行), 并且每一列都有元素被选取到条件下得元素最小和. 每行都有元素要被选取的约束是由三重循环中 j 和 k 都由 1 开始满足的.

// 优化
// 感谢 @Freezer 在评论区里提出用 C++ 会超时的问题, 说明上面的解法效率还没有达到最高. 实际上面解法中的三重循环可以进行优化. 考虑到当我们已知截至上一行时的各列选取情况, 其实并没有必要重复地选取上面已经选取过的列, 据此可以对三重循环做如下修改:


// for (int i = 1; i < m; i++) {
//   for (int k = 1; k < (1 << n); k++) {
//     // 首先将第 i 行只选取一个元素的情况都考虑一遍
//     // 这样做的目的是保证第 i 行至少选取了一个元素
//     for (int j = 0; j < n; j++)
//       dp[i][k | (1 << j)] = Math.min(dp[i][k | (1 << j)], dp[i - 1][k] + cost.get(i).get(j));
//     // rest 表示截至第 i 行还没被选过的列
//     int rest = (1 << n) - 1 - k;
//     // 只遍历没选过的列的所有组合
//     for (int j = rest; j  >= 1; j = rest & (j - 1)) {
//       dp[i][j | k] = Math.min(dp[i][j | k], dp[i - 1][k] + costMatrix[i][j]);
//     }
//   }
// }
// 另外注意 dp 中的每一行只依赖上一行的值, 可以将 dp 变为一维数组以优化空间复杂度 (不写了).

// 作者：aerysn
// 链接：https://leetcode-cn.com/problems/minimum-cost-to-connect-two-groups-of-points/solution/java-zhuang-tai-ya-suo-dp-by-aerysn/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。