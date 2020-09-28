package updateBoard

var d = [8][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	a, b := click[0], click[1]
	if board[a][b] == 'M' {
		board[a][b] = 'X'
	} else if board[a][b] == 'E' {
		m, n := len(board), len(board[0])
		var f func(int, int); f = func(i, j int) {
			c := byte('0')
			for _, di := range d {
				x, y := i + di[0], j + di[1]
				if 0 <= x && x < m && 0 <= y && y < n && board[x][y] == 'M' {
					c++
				}
			}
			if c > '0' {
				board[i][j] = c
			} else {
				board[i][j] = 'B'
				for _, di := range d {
					x, y := i + di[0], j + di[1]
					if 0 <= x && x < m && 0 <= y && y < n && board[x][y] == 'E' {
						f(x, y)
					}
				}
			}
		}
		f(a, b)
	}
	return board
}
