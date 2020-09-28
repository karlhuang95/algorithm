package solve

func solve(board [][]byte) {
	// 思路，从四条边的"O"出发，寻找所有可达的"O"，将剩余的覆盖为"X"即可
	m := len(board)
	if m < 2 {
		return
	}
	n := len(board[0])
	if n < 2 {
		return
	}

	arr := make([][]bool, m, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]bool, n, n)
	}

	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			dfs(0, i, board, arr)
		}
		if board[m-1][i] == 'O' {
			dfs(m-1, i, board, arr)
		}
	}

	for i := 1; i < m-1; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0, board, arr)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1, board, arr)
		}
	}

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if board[i][j] == 'O' && !arr[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

var nx = []int{1, 0, -1, 0}
var ny = []int{0, 1, 0, -1}

func dfs(x, y int, board [][]byte, arr [][]bool) {
	if x < 0 || y < 0 || x == len(board) || y == len(board[0]) {
		return
	}
	if board[x][y] != 'O' {
		return
	}
	if board[x][y] == 'O' && arr[x][y] {
		return
	}
	arr[x][y] = true
	for i := 0; i < 4; i++ {
		dfs(x+nx[i], y+ny[i], board, arr)
	}
}
