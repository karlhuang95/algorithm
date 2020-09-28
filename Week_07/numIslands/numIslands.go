package numIslands

func numIslands(grid [][]byte) int {

	len1 := len(grid)
	if len1 == 0 {
		return 0
	}
	len2 := len(grid[0])
	res := 0
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if grid[i][j] == '1' {
				helper(i, j, grid)
				res++
			}
		}
	}
	return res
}

func helper(i, j int, grid [][]byte) {
	grid[i][j] = '0'
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		helper(i+1, j, grid)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' {
		helper(i-1, j, grid)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		helper(i, j+1, grid)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		helper(i, j-1, grid)
	}
}
