package searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
	lx := len(matrix)
	if lx == 0 {
		return false // 边界检查
	}
	ly := len(matrix[0])
	if ly == 0 || target < matrix[0][0] || target > matrix[lx-1][ly-1] {
		return false // 特殊情况一步到位
	}
	x0, x1 := 0, lx-1
	for x0 <= x1 {
		mx := x0&x1 + (x0^x1)>>1 // 二分行
		if target == matrix[mx][0] {
			return true
		}
		if target > matrix[mx][0] {
			y0, y1 := 0, ly-1
			for y0 <= y1 {
				my := y0&y1 + (y0^y1)>>1 // 二分列
				if target == matrix[mx][my] {
					return true
				}
				if target > matrix[mx][my] {
					y0 = my + 1
				} else {
					y1 = my - 1
				}
			}
			x0 = mx + 1
		} else {
			x1 = mx - 1
		}
	}
	return false
}
