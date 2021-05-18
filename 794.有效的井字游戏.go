/*
 * @lc app=leetcode.cn id=794 lang=golang
 *
 * [794] 有效的井字游戏
 */

// @lc code=start
func validTicTacToe(board []string) bool {
	xc, oc := 0, 0
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				xc++
			} else if board[i][j] == 'O' {
				oc++
			}
		}
	}
	if xc-oc < 0 || xc-oc >= 2 {
		return false
	}
	pass := func(win byte) bool {
		for i := 0; i < m; i++ {
			if board[i][0] == win && board[i][1] == win && board[i][2] == win {
				return true
			}
		}
		for i := 0; i < n; i++ {
			if board[0][i] == win && board[1][i] == win && board[2][i] == win {
				return true
			}
		}
		if board[0][0] == win && board[1][1] == win && board[2][2] == win {
			return true
		}
		if board[0][2] == win && board[1][1] == win && board[2][0] == win {
			return true
		}
		return false
	}
	if pass(byte('X')) && xc != oc+1 {
		return false
	}
	if pass(byte('O')) && xc != oc {
		return false
	}
	return true
}

// @lc code=end

// class Solution {
//     public boolean validTicTacToe(String[] board) {
//         int xCount = 0, oCount = 0;
//         for (String row: board)
//             for (char c: row.toCharArray()) {
//                 if (c == 'X') xCount++;
//                 if (c == 'O') oCount++;
//             }

//         if (oCount != xCount && oCount != xCount - 1) return false;
//         if (win(board, 'X') && oCount != xCount - 1) return false;
//         if (win(board, 'O') && oCount != xCount) return false;
//         return true;
//     }

//     public boolean win(String[] B, char P) {
//         // B: board, P: player
//         for (int i = 0; i < 3; ++i) {
//             if (P == B[0].charAt(i) && P == B[1].charAt(i) && P == B[2].charAt(i))
//                 return true;
//             if (P == B[i].charAt(0) && P == B[i].charAt(1) && P == B[i].charAt(2))
//                 return true;
//         }
//         if (P == B[0].charAt(0) && P == B[1].charAt(1) && P == B[2].charAt(2))
//             return true;
//         if (P == B[0].charAt(2) && P == B[1].charAt(1) && P == B[2].charAt(0))
//             return true;
//         return false;
//     }
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/valid-tic-tac-toe-state/solution/you-xiao-de-jing-zi-you-xi-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。