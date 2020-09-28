package totalNQueens

var count int

func totalNQueens(n int) int {
	count = 0
	dfs(n, 0, 0, 0, 0)
	return count
}

func dfs(n, row, col, pie, na int) {
	if row >= n {
		count++
		return
	}

	bits := (^(col | pie | na)) & ((1 << uint(n)) - 1) //获取所有可填空位
	for bits != 0 {
		bit := bits & -bits                               //取位置排在最后的一个空位
		dfs(n, row+1, col|bit, (pie|bit)<<1, (na|bit)>>1) //递归遍历下一行
		bits = bits & (bits - 1)                          //打掉最后位置上的1， 因为该位置已被占用
	}

}
