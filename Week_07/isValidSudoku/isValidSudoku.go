package isValidSudoku

func isValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')                      // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3                                   // 3x3的块索引号
			if (row[i]&cur)|(col[j]&cur)|(block[bi]&cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}
